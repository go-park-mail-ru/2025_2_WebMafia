package http

import (
	"context"
	"spotify/internal/user/dto"
	"spotify/pkg/jwtmanager"
)

type IService interface {
	Register(ctx context.Context, req dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error)
}
type Handler struct {
	svc        IService
	jwtManager *jwtmanager.Manager
}

func NewHandler(svc IService, jwtManager *jwtmanager.Manager) *Handler {
	return &Handler{
		svc:        svc,
		jwtManager: jwtManager,
	}
}
