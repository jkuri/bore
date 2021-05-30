package server

import (
	"fmt"
	"net"
	"net/http"

	"go.uber.org/zap"
)

// HTTPServer extends net/http Server with graceful
// shutdowns.
type HTTPServer struct {
	*http.Server
	listener  net.Listener
	isRunning bool
	running   chan error
	logger    *zap.SugaredLogger
}

// NewHTTPServer creates a new HTTPServer instance.
func NewHTTPServer(logger *zap.SugaredLogger) *HTTPServer {
	return &HTTPServer{
		Server:    &http.Server{},
		running:   make(chan error),
		logger:    logger,
		isRunning: true,
	}
}

// Run starts HTTPServer instance and listens on specified addr.
func (h *HTTPServer) Run(addr string, handler http.Handler) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	h.Handler = handler
	h.listener = listener

	h.logger.Infof("starting HTTP server on %s", addr)

	go h.closeWith(h.Serve(listener))
	return nil
}

// Close closes the HTTPServer instance
func (h *HTTPServer) Close() error {
	h.closeWith(nil)
	return h.listener.Close()
}

// Wait waits for server to be stopped
func (h *HTTPServer) Wait() error {
	if !h.isRunning {
		return fmt.Errorf("already closed")
	}
	return <-h.running
}

func (h *HTTPServer) closeWith(err error) {
	if !h.isRunning {
		return
	}
	h.isRunning = false
	h.running <- err
}
