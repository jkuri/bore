package server

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
)

const (
	minPort = 55000
	maxPort = 65000
)

// SSHServer defines SSH server instance.
type SSHServer struct {
	mu        sync.Mutex
	opts      *Options
	listener  net.Listener
	config    *ssh.ServerConfig
	running   chan error
	isRunning bool
	clients   map[string]*client
	addr      string
	domain    string
	password  string
	logger    *zap.SugaredLogger
}

type client struct {
	mu        sync.Mutex
	id        string
	tcpConn   net.Conn
	sshConn   *ssh.ServerConn
	ch        ssh.Channel
	listeners map[string]net.Listener
	addr      string
	port      uint32
	password  string
}

func (c *client) write(data string) {
	if c.ch != nil {
		io.WriteString(c.ch, data)
	}
}

// NewSSHServer returns new instance of SSHServer.
func NewSSHServer(opts *Options, logger *zap.SugaredLogger) *SSHServer {
	return &SSHServer{
		opts: opts,
		config: &ssh.ServerConfig{
			NoClientAuth: true,
		},
		running:   make(chan error, 1),
		clients:   make(map[string]*client),
		logger:    logger,
		isRunning: true,
	}
}

// Run starts the SSH server.
func (s *SSHServer) Run() error {
	privateKeyContent, err := os.ReadFile(s.opts.PrivateKey)
	if err != nil {
		return err
	}
	private, err := ssh.ParsePrivateKey(privateKeyContent)
	if err != nil {
		return err
	}
	s.config.AddHostKey(private)
	s.addr = s.opts.SSHAddr
	s.domain = s.opts.Domain
	s.password = s.opts.Password

	go s.closeWith(s.listen())
	return nil
}

// Close closes and stops the SSH server.
func (s *SSHServer) Close() error {
	s.closeWith(nil)
	return s.listener.Close()
}

// Wait waits for server to be stopped
func (s *SSHServer) Wait() error {
	if !s.isRunning {
		return fmt.Errorf("already closed")
	}
	return <-s.running
}

func (s *SSHServer) closeWith(err error) {
	if !s.isRunning {
		return
	}
	s.isRunning = false
	s.running <- err
}

func (s *SSHServer) listen() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	s.listener = listener

	s.logger.Infof("starting SSH server on %s", s.addr)

	for {
		tcpConn, err := s.listener.Accept()
		if err != nil {
			s.logger.Errorf("failed to accept incoming connection: %v", err)
			continue
		}

		sshConn, chans, reqs, err := ssh.NewServerConn(tcpConn, s.config)
		if err != nil {
			s.logger.Errorf("failed to handshake: %v", err)
			continue
		}

	genid:
		id := randID()
		if _, ok := s.clients[id]; ok {
			goto genid
		}

		c := &client{
			id:        id,
			tcpConn:   tcpConn,
			sshConn:   sshConn,
			listeners: make(map[string]net.Listener),
			addr:      "",
			port:      0,
			password:  s.password,
		}
		s.logger.Infof("new SSH connection from %s (%s)", sshConn.RemoteAddr().String(), sshConn.ClientVersion())

		go func(c *client) {
			err := c.sshConn.Wait()
			s.logger.Infof("[%s] SSH connection closed: %v", c.id, err)

			c.mu.Lock()
			for bind, listener := range c.listeners {
				s.logger.Debugf("[%s] closing listener bound to %s", c.id, bind)
				listener.Close()
			}
			c.mu.Unlock()

			s.mu.Lock()
			delete(s.clients, c.id)
			s.mu.Unlock()
		}(c)

		go s.handleRequests(c, reqs)
		go s.handleChannels(c, chans)
	}
}

func (s *SSHServer) handleChannels(client *client, chans <-chan ssh.NewChannel) {
	for nch := range chans {
		chconn, _, err := nch.Accept()
		if err != nil {
			s.logger.Errorf("[%s] could not accept channel: %v", client.id, err)
			return
		}
		client.ch = chconn
	}
}

func (s *SSHServer) handleRequests(client *client, reqs <-chan *ssh.Request) {
	authenticated := s.password == ""

	for req := range reqs {
		client.tcpConn.SetDeadline(time.Now().Add(2 * time.Minute))

		if !authenticated {
			if req.Type != "password" {
				req.Reply(false, nil)
				s.logger.Errorf("[%s] error: password authentication required", client.id)
				client.tcpConn.Close()
				return
			}

			var payload passwordRequestPayload
			if err := ssh.Unmarshal(req.Payload, &payload); err != nil {
				s.logger.Errorf("[%s] Unable to unmarshal payload: %v", client.id, err)
				req.Reply(false, nil)
				client.tcpConn.Close()
				return
			}

			if payload.Password != s.password {
				req.Reply(false, nil)
				s.logger.Errorf("[%s] error, disconnecting: wrong password", client.id)
				client.tcpConn.Close()
				return
			}

			authenticated = true
			req.Reply(true, nil)
			continue
		}

		if req.Type == "set-id" {
			var payload idRequestPayload
			if err := ssh.Unmarshal(req.Payload, &payload); err != nil {
				s.logger.Errorf("[%s] Unable to unmarshal payload: %v", client.id, err)
			}
			if payload.ID != "" {
				if _, ok := s.clients[payload.ID]; !ok {
					s.mu.Lock()
					delete(s.clients, client.id)
					client.id = payload.ID
					s.clients[client.id] = client
					s.mu.Unlock()
				}
			}
			req.Reply(true, []byte{})
			continue
		}

		if req.Type == "tcpip-forward" {
			listener, bindInfo, err := s.handleForward(client, req)
			if err != nil {
				s.logger.Errorf("[%s] error, disconnecting: %v", client.id, err)
				client.tcpConn.Close()
				continue
			}

			client.addr = bindInfo.Addr
			client.port = bindInfo.Port

			client.mu.Lock()
			client.listeners[bindInfo.Bound] = listener
			client.mu.Unlock()

			s.mu.Lock()
			s.clients[client.id] = client
			s.mu.Unlock()

			go s.handleListener(client, bindInfo, listener)

			if client.ch != nil {
				data := clientResponse{
					id:     client.id,
					domain: s.domain,
					port:   client.port,
				}

				renderMessage(data, client.ch)
				renderTable(data, client.ch)
			}
		} else {
			req.Reply(false, []byte{})
		}
	}
}

func (s *SSHServer) handleListener(client *client, bindInfo *bindInfo, listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			neterr := err.(net.Error)
			if neterr.Timeout() {
				s.logger.Errorf("[%s] accept failed with timeout: %v", client.id, err)
				continue
			}
			if neterr.Temporary() {
				s.logger.Errorf("[%s] accept failed with temporary: %v", client.id, err)
				continue
			}

			break
		}

		go s.handleForwardTCPIP(client, bindInfo, conn)
	}
}

func (s *SSHServer) handleForwardTCPIP(client *client, bindInfo *bindInfo, conn net.Conn) {
	remoteAddr := conn.RemoteAddr().(*net.TCPAddr)
	raddr := remoteAddr.IP.String()
	rport := uint32(remoteAddr.Port)

	payload := forwardedTCPPayload{bindInfo.Addr, bindInfo.Port, raddr, rport}
	mpayload := ssh.Marshal(&payload)

	// open channel with client
	c, requests, err := client.sshConn.OpenChannel("forwarded-tcpip", mpayload)
	if err != nil {
		s.logger.Errorf("[%s] unable to get channel: %v. Hanging up requesting party!", client.id, err)
		conn.Close()
		return
	}
	s.logger.Debugf("[%s] channel opened for client %s:%d <-> %s", client.id, bindInfo.Addr, bindInfo.Port, remoteAddr.String())
	go ssh.DiscardRequests(requests)
	go s.handleForwardTCPIPTransfer(c, conn)
}

func (s *SSHServer) handleForward(client *client, req *ssh.Request) (net.Listener, *bindInfo, error) {
	var payload tcpIPForwardPayload
	if err := ssh.Unmarshal(req.Payload, &payload); err != nil {
		s.logger.Errorf("[%s] unable to unmarshal payload: %v", client.id, err)
		req.Reply(false, []byte{})
		return nil, nil, fmt.Errorf("unable to parse payload")
	}

	s.logger.Debugf("[%s] request: %s %v %v", client.id, req.Type, req.WantReply, payload)

listen:
	bind := fmt.Sprintf("%s:%d", payload.Addr, payload.Port)
	if payload.Port == 0 {
		bind = fmt.Sprintf("%s:%d", payload.Addr, randomPort(minPort, maxPort))
	}

	ln, err := net.Listen("tcp", bind)
	if err != nil {
		if payload.Port == 0 {
			s.logger.Errorf("[%s] listen failed for: %s %v, retrying on another port", client.id, bind, err)
			goto listen
		}
		s.logger.Errorf("[%s] listen failed for: %s %v", client.id, bind, err)
		req.Reply(false, []byte{})
		return nil, nil, fmt.Errorf("unable to listen")
	}
	port := ln.Addr().(*net.TCPAddr).Port
	bind = fmt.Sprintf("%s:%d", payload.Addr, port)

	s.logger.Debugf("[%s] listening on %s", client.id, bind)
	reply := tcpIPForwardPayloadReply{uint32(port)}
	req.Reply(true, ssh.Marshal(&reply))

	return ln, &bindInfo{bind, uint32(port), payload.Addr}, nil
}

func (s *SSHServer) handleForwardTCPIPTransfer(c ssh.Channel, conn net.Conn) {
	defer conn.Close()
	defer c.Close()
	done := make(chan struct{})

	go func() {
		io.Copy(c, conn)
		done <- struct{}{}
	}()

	go func() {
		io.Copy(conn, c)
		done <- struct{}{}
	}()

	<-done
}
