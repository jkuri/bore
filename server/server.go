package server

import (
	"net/http"

	"github.com/google/wire"
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
	httpServer *HTTPServer
}

// NewBoreServer returns new instance of BoreServer.
func NewBoreServer(opts *Options, logger *zap.Logger) *BoreServer {
	log := logger.Sugar()
	return &BoreServer{
		opts:       opts,
		httpServer: NewHTTPServer(log),
	}
}

// Run starts the bore server.
func (s *BoreServer) Run() error {
	if err := s.httpServer.Run(s.opts.HTTPAddr, s.handleHTTP()); err != nil {
		return err
	}

	return s.httpServer.Wait()
}

func (s *BoreServer) handleHTTP() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// host := r.Host
		// pos := strings.IndexByte(host, '.')
		// if pos > 0 {
		// 	userID := host[:pos]

		// }

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
	})
}
