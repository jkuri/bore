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
	metricsHub *MetricsHub
	UI         http.Handler
}

// NewBoreServer returns new instance of BoreServer.
func NewBoreServer(opts *Options, logger *zap.Logger) *BoreServer {
	log := logger.Sugar()
	landingFS, _ := fs.New()

	sshServer := NewSSHServer(opts, log)
	metricsHub := NewMetricsHub(sshServer, log)
	sshServer.metricsHub = metricsHub

	return &BoreServer{
		opts:       opts,
		sshServer:  sshServer,
		httpServer: NewHTTPServer(log),
		metricsHub: metricsHub,
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

		host, _, err := net.SplitHostPort(r.Host)
		if err != nil {
			host = r.Host
		}

		var log string
		if host != s.opts.Domain {
			userID := strings.Split(host, ".")[0]
			log = fmt.Sprintf(
				"[%s] %s %s (code=%d dt=%s written=%s remote=%s)",
				userID,
				r.Method,
				r.URL,
				m.Code,
				m.Duration,
				humanize.Bytes(uint64(m.Written)),
				remote,
			)
			s.httpServer.logger.Debug(log)

			s.sshServer.mu.Lock()
			client, ok := s.sshServer.clients[userID]
			s.sshServer.mu.Unlock()
			if ok {
				client.write(fmt.Sprintf("%s\n", log))
				s.metricsHub.RecordTraffic(userID, uint64(r.ContentLength), uint64(m.Written))
			}
		} else {
			log = fmt.Sprintf(
				"%s %s (code=%d dt=%s written=%s remote=%s)",
				r.Method,
				r.URL,
				m.Code,
				m.Duration,
				humanize.Bytes(uint64(m.Written)),
				remote,
			)
			s.httpServer.logger.Debug(log)
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

			s.sshServer.mu.Lock()
			client, ok := s.sshServer.clients[userID]
			s.sshServer.mu.Unlock()

			if ok {
				w.Header().Set("X-Proxy", "bore")

				if strings.ToLower(r.Header.Get("Upgrade")) == "websocket" {
					url := &url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%d", client.addr, client.port)}
					proxy := wsutil.NewSingleHostReverseProxy(url)
					proxy.ServeHTTP(w, r)
					return
				}

				url := &url.URL{Scheme: "http", Host: fmt.Sprintf("%s:%d", client.addr, client.port)}
				proxy := httputil.NewSingleHostReverseProxy(url)
				proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
					if strings.Contains(err.Error(), "connection refused") || strings.Contains(err.Error(), "EOF") {
						s.httpServer.logger.Debugf("[%s] tunnel closed during request: %v", userID, err)
					} else {
						s.httpServer.logger.Errorf("[%s] proxy error: %v", userID, err)
					}
					w.WriteHeader(http.StatusBadGateway)
				}
				proxy.ServeHTTP(w, r)
				return
			}

			url := &url.URL{Scheme: r.URL.Scheme, Host: s.opts.Domain, Path: "not-found", RawQuery: fmt.Sprintf("tunnelID=%s", userID)}
			http.Redirect(w, r, url.String(), http.StatusMovedPermanently)
			return
		}

		if r.URL.Path == "/api/ws/dashboard" {
			s.metricsHub.HandleWebSocket(w, r)
			return
		}

		s.UI.ServeHTTP(w, r)
	})
}
