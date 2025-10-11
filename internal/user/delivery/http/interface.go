package http

import (
	"context"
	"spotify/internal/user/model"
	"spotify/pkg/jwtmanager"
)

type IService interface {
	Register(ctx context.Context, login, email, password string) (*model.User, error)
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
