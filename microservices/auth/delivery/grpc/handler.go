package grpc

import (
	"context"
	"spotify/microservices/auth/dto"
	"spotify/pkg/csrfmanager"
	"spotify/pkg/jwtmanager"
	pb "spotify/proto/auth"
)

//go:generate mockgen -destination=../../mocks/grpc_service/user_service_mock.go -package=auth_mocks spotify/microservices/auth/delivery/grpc IUserService
type IUserService interface {
	GetUsersBatch(ctx context.Context, ids []string) ([]dto.GetProfileResponse, error)
}

type Handler struct {
	pb.UnimplementedAuthServiceServer
	jwtManager  *jwtmanager.Manager
	csrfManager *csrfmanager.Manager
	userService IUserService
}

func NewHandler(jwtManager *jwtmanager.Manager, csrfManager *csrfmanager.Manager, us IUserService) *Handler {
	return &Handler{
		jwtManager:  jwtManager,
		csrfManager: csrfManager,
		userService: us,
	}
}
