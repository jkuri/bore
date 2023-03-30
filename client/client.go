package client

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"time"

	"golang.org/x/crypto/ssh"
)

// BoreClient defines bore client.
type BoreClient struct {
	config         Config
	sshConfig      *ssh.ClientConfig
	sshClient      *ssh.Client
	LocalEndpoint  endpoint // local service to be forwarded
	ServerEndpoint endpoint // remote SSH server
	RemoteEndpoint endpoint // remote forwarding port (on remote SSH server network)
	id             string
}

type idRequestPayload struct {
	ID string
}

// NewBoreClient returns new instance of BoreClient.
func NewBoreClient(config Config) BoreClient {
	return BoreClient{
		config:         config,
		LocalEndpoint:  endpoint{config.LocalServer, config.LocalPort},
		ServerEndpoint: endpoint{config.RemoteServer, config.RemotePort},
		RemoteEndpoint: endpoint{"0.0.0.0", config.BindPort},
		sshConfig:      &ssh.ClientConfig{HostKeyCallback: ssh.InsecureIgnoreHostKey()},
		id:             config.ID,
	}
}

// Run starts the client.
func (c *BoreClient) Run() error {
	// Healthcheck
	local, err := net.Dial("tcp", c.LocalEndpoint.String())
	if err != nil {
		return err
	}
	_ = local.Close()

	ch := make(chan os.Signal, 1)
	errch := make(chan error)
	signal.Notify(ch, os.Interrupt)

	client, err := ssh.Dial("tcp", c.ServerEndpoint.String(), c.sshConfig)
	if err != nil {
		return err
	}
	c.sshClient = client

	done := make(chan struct{})
	if c.config.KeepAlive {
		go keepAliveTicker(c.sshClient, done)
	}

	if c.id != "" {
		_, _, err = c.sshClient.SendRequest("set-id", true, ssh.Marshal(&idRequestPayload{c.id}))
		if err != nil {
			return err
		}
	}

	if err := c.writeStdout(); err != nil {
		return err
	}

	listener, err := c.sshClient.Listen("tcp", c.RemoteEndpoint.String())
	if err != nil {
		return err
	}
	defer listener.Close()

	go func() {
		for {
			local, err := net.Dial("tcp", c.LocalEndpoint.String())
			if err != nil {
				errch <- err
				return
			}

			client, err := listener.Accept()
			if err != nil {
				errch <- err
				return
			}

			go handleClient(client, local)
		}
	}()

	select {
	case <-ch:
		return nil
	case err := <-errch:
		return err
	}
}

func (c *BoreClient) writeStdout() error {
	session, err := c.sshClient.NewSession()
	if err != nil {
		return err
	}
	stdout, err := session.StdoutPipe()
	if err != nil {
		return err
	}

	go func() {
		defer session.Close()
		io.Copy(os.Stdout, stdout)
	}()

	return nil
}

type endpoint struct {
	host string
	port int
}

func (e *endpoint) String() string {
	return fmt.Sprintf("%s:%d", e.host, e.port)
}

func handleClient(client net.Conn, remote net.Conn) {
	defer client.Close()
	defer remote.Close()
	done := make(chan struct{})

	go func() {
		io.Copy(client, remote)
		done <- struct{}{}
	}()

	go func() {
		io.Copy(remote, client)
		done <- struct{}{}
	}()

	<-done
}

func keepAliveTicker(client *ssh.Client, done <-chan struct{}) error {
	t := time.NewTicker(time.Minute)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			_, _, err := client.SendRequest("keepalive", true, nil)
			if err != nil {
				return err
			}
		case <-done:
			return nil
		}
	}
}
