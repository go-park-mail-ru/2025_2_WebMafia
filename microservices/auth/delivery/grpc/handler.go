package grpc

import (
	"spotify/pkg/csrfmanager"
	"spotify/pkg/jwtmanager"
	pb "spotify/proto/auth"
)

type Handler struct {
	pb.UnimplementedAuthServiceServer
	jwtManager  *jwtmanager.Manager
	csrfManager *csrfmanager.Manager
}

func NewHandler(jwtManager *jwtmanager.Manager, csrfManager *csrfmanager.Manager) *Handler {
	return &Handler{
		jwtManager:  jwtManager,
		csrfManager: csrfManager,
	}
}
