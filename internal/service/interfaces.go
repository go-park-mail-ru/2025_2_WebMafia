package service

import (
	"context"
	"spotify/internal/model"
)

type AuthService interface {
	Register(ctx context.Context, email, login, password string) (model.User, error)
	Login(ctx context.Context, login, password string) (string, error)
}
