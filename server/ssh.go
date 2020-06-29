package server

import (
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"sync"
	"time"

	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
)

// SSHServer defines SSH server instance.
type SSHServer struct {
	mu        sync.Mutex
	listener  net.Listener
	config    *ssh.ServerConfig
	running   chan error
	isRunning bool
	clients   map[string]*client
	addr      string
	domain    string
	logger    *zap.SugaredLogger
}

type client struct {
	mu        sync.Mutex
	id        string
	tcpConn   net.Conn
	sshConn   *ssh.ServerConn
	listeners map[string]net.Listener
	addr      string
	port      uint32
}

// NewSSHServer returns new instance of SSHServer.
func NewSSHServer(logger *zap.SugaredLogger) *SSHServer {
	return &SSHServer{
		config: &ssh.ServerConfig{
			NoClientAuth: true,
		},
		running: make(chan error, 1),
		clients: make(map[string]*client),
		logger:  logger,
	}
}

// Run starts the SSH server.
func (s *SSHServer) Run(opts *Options) error {
	privateKeyContent, err := ioutil.ReadFile(opts.PrivateKey)
	if err != nil {
		return err
	}
	private, err := ssh.ParsePrivateKey(privateKeyContent)
	if err != nil {
		return err
	}
	s.config.AddHostKey(private)
	s.addr = opts.SSHAddr
	s.domain = opts.Domain

	go s.closeWith(s.listen())

	return nil
}

// Close closes and stops the SSH server.
func (s *SSHServer) Close() error {
	s.closeWith(nil)
	return s.listener.Close()
}

// Wait waits for SSH server to close.
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
		panic(err)
	}
	s.listener = listener

	s.logger.Infof("Starting SSH server on %s", s.addr)

	for {
		tcpConn, err := s.listener.Accept()
		if err != nil {
			s.logger.Errorf("Failed to accept incoming connection: %v", err)
			continue
		}

		sshConn, chans, reqs, err := ssh.NewServerConn(tcpConn, s.config)
		if err != nil {
			s.logger.Errorf("Failed to handshake: %v", err)
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
		}
		s.logger.Infof("New SSH connection from %s (%s)", sshConn.RemoteAddr().String(), sshConn.ClientVersion())

		go func(c *client) {
			err := c.sshConn.Wait()
			s.logger.Infof("[%s] SSH connection closed: %v", c.id, err)

			c.mu.Lock()
			for bind, listener := range c.listeners {
				s.logger.Debugf("[%s] Closing listener bound to %s", c.id, bind)
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
			s.logger.Errorf("[%s] Could not accept channel: %v", client.id, err)
			return
		}

		url := fmt.Sprintf("http://%s.%s", client.id, s.domain)
		io.WriteString(chconn, url)
	}
}

func (s *SSHServer) handleRequests(client *client, reqs <-chan *ssh.Request) {
	for req := range reqs {
		client.tcpConn.SetDeadline(time.Now().Add(2 * time.Minute))

		if req.Type == "tcpip-forward" {
			client.mu.Lock()
			listener, bindInfo, err := s.handleForward(client, req)
			if err != nil {
				s.logger.Errorf("[%s] Error, disconnecting: %v", client.id, err)
				client.mu.Unlock()
				client.tcpConn.Close()
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
				s.logger.Errorf("[%s] Accept failed with timeout: %v", client.id, err)
				continue
			}
			if neterr.Temporary() {
				s.logger.Errorf("[%s] Accept failed with temporary: %v", client.id, err)
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
		s.logger.Errorf("[%s] Unable to get channel: %v. Hanging up requesting party!", client.id, err)
		conn.Close()
		return
	}
	s.logger.Debugf("[%s] Channel opened for client %s:%d <-> %s", client.id, bindInfo.Addr, bindInfo.Port, remoteAddr.String())
	go ssh.DiscardRequests(requests)
	go s.handleForwardTCPIPTransfer(c, conn)
}

func (s *SSHServer) handleForward(client *client, req *ssh.Request) (net.Listener, *bindInfo, error) {
	var payload tcpIPForwardPayload
	if err := ssh.Unmarshal(req.Payload, &payload); err != nil {
		s.logger.Errorf("[%s] Unable to unmarshal payload: %v", client.id, err)
		req.Reply(false, []byte{})
		return nil, nil, fmt.Errorf("Unable to parse payload")
	}

	s.logger.Debugf("[%s] Request: %s %v %v", client.id, req.Type, req.WantReply, payload)

	bind := fmt.Sprintf("%s:%d", payload.Addr, payload.Port)
	ln, err := net.Listen("tcp", bind)
	if err != nil {
		s.logger.Errorf("[%s] Listen failed for: %s %v", client.id, bind, err)
		req.Reply(false, []byte{})
		return nil, nil, err
	}

	s.logger.Debugf("[%s] Listening on %s", client.id, bind)
	reply := tcpIPForwardPayloadReply{payload.Port}
	req.Reply(true, ssh.Marshal(&reply))

	return ln, &bindInfo{bind, payload.Port, payload.Addr}, nil
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

func randID() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
