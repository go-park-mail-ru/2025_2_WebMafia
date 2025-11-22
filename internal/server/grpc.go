package server

import (
	"fmt"
	"net"
	"spotify/internal/app"
	"spotify/internal/interceptors"
	"spotify/pkg/logger"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	server *grpc.Server
	cfg    *app.GRPCConfig
	logger logger.Logger
}

func NewGRPCServer(cfg *app.GRPCConfig, log logger.Logger, registerServices func(server *grpc.Server)) *GRPCServer {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptors.RequestLogger(log),
			interceptors.PanicRecovery,
		),
	)

	if registerServices != nil {
		registerServices(grpcServer)
	}

	return &GRPCServer{
		server: grpcServer,
		cfg:    cfg,
		logger: log,
	}
}

func (s *GRPCServer) Run() error {
	lis, err := net.Listen("tcp", ":"+s.cfg.Port)
	if err != nil {
		return fmt.Errorf("failed to listen gRPC port: %w", err)
	}

	s.logger.Infof("gRPC server is listening on :%s", s.cfg.Port)
	if err := s.server.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve gRPC: %w", err)
	}
	s.logger.Infof("gRPC server stopped.")
	return nil
}

func (s *GRPCServer) Stop() {
	s.logger.Infof("stopping gRPC server...")
	s.server.GracefulStop()
}
