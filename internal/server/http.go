package server

import (
	"context"
	"net/http"
	"spotify/internal/app"
	"spotify/pkg/logger"
)

type Server struct {
	server *http.Server
	logger logger.Logger
	cfg    *app.HTTPConfig
}

func NewHTTPServer(cfg *app.HTTPConfig, handler http.Handler, log logger.Logger) *Server {
	return &Server{
		server: &http.Server{
			Addr:         ":" + cfg.Port,
			Handler:      handler,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
			IdleTimeout:  cfg.IdleTimeout,
		},
		logger: log,
		cfg:    cfg,
	}
}

func (s *Server) Run() error {
	s.logger.Infof("HTTP server is starting on port %s", s.server.Addr)
	err := s.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		s.logger.Errorf("http server error: %v", err)
		return err
	}
	s.logger.Infof("HTTP server stopped.")
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Infof("shutting down HTTP server...")
	return s.server.Shutdown(ctx)
}
