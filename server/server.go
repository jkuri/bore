package server

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/felixge/httpsnoop"
	"github.com/google/wire"
	_ "github.com/jkuri/bore/internal/ui/landing" // landing UI
	"github.com/jkuri/statik/fs"
	"github.com/yhat/wsutil"
	"go.uber.org/zap"
)

// ProviderSet exports for wire DI.
var ProviderSet = wire.NewSet(
	NewConfig,
	NewOptions,
	NewBoreServer,
)

// BoreServer defines main struct for bore server and
// includes HTTP and SSH server instances.
type BoreServer struct {
	opts       *Options
	sshServer  *SSHServer
	httpServer *HTTPServer
	UI         http.Handler
}

// NewBoreServer returns new instance of BoreServer.
func NewBoreServer(opts *Options, logger *zap.Logger) *BoreServer {
	log := logger.Sugar()
	landingFS, _ := fs.New()

	return &BoreServer{
		opts:       opts,
		sshServer:  NewSSHServer(opts, log),
		httpServer: NewHTTPServer(log),
		UI:         http.FileServer(&statikWrapper{landingFS}),
	}
}

// Run starts the bore server.
func (s *BoreServer) Run() error {
	errch := make(chan error)

	go func() {
		if err := s.httpServer.Run(s.opts.HTTPAddr, s.getHandler(s.handleHTTP())); err != nil {
			errch <- err
		}
	}()

	go func() {
		if err := s.sshServer.Run(); err != nil {
			errch <- err
		}
	}()

	go func() {
		if err := s.httpServer.Wait(); err != nil {
			errch <- err
		}
	}()

	go func() {
		if err := s.sshServer.Wait(); err != nil {
			errch <- err
		}
	}()

	return <-errch
}

func (s *BoreServer) getHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(handler, w, r)
		remote := r.Header.Get("X-Forwarded-For")
		if remote == "" {
			remote = r.RemoteAddr
		}
		log := fmt.Sprintf(
			"%s %s (code=%d dt=%s written=%s remote=%s)",
			r.Method,
			r.URL,
			m.Code,
			m.Duration,
			humanize.Bytes(uint64(m.Written)),
			remote,
		)
		s.httpServer.logger.Debug(log)

		userID := strings.Split(r.Host, ".")[0]
		if client, ok := s.sshServer.clients[userID]; ok {
			client.write(log)
		}
	})
}

func (s *BoreServer) handleHTTP() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host, _, err := net.SplitHostPort(r.Host)
		if err != nil {
			host = r.Host
		}

		if host != s.opts.Domain {
			splitted := strings.Split(host, ".")
			userID := splitted[0]

			if client, ok := s.sshServer.clients[userID]; ok {
				w.Header().Set("X-Proxy", "bore")

				if strings.ToLower(r.Header.Get("Upgrade")) == "websocket" {
					url := &url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%d", client.addr, client.port)}
					proxy := wsutil.NewSingleHostReverseProxy(url)
					proxy.ServeHTTP(w, r)
					return
				}

				url := &url.URL{Scheme: "http", Host: fmt.Sprintf("%s:%d", client.addr, client.port)}
				proxy := httputil.NewSingleHostReverseProxy(url)
				proxy.ServeHTTP(w, r)
				return
			}

			url := &url.URL{Scheme: r.URL.Scheme, Host: s.opts.Domain, Path: "not-found", RawQuery: fmt.Sprintf("tunnelID=%s", userID)}
			http.Redirect(w, r, url.String(), http.StatusMovedPermanently)
			return
		}

		s.UI.ServeHTTP(w, r)
	})
}
