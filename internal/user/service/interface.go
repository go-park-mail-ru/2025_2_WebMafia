package service

import (
	"context"
	"spotify/internal/user/model"
)

type IRepository interface {
	CreateUser(ctx context.Context, user model.User) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByLogin(ctx context.Context, login string) (*model.User, error)
}

type Service struct {
	repo IRepository
}

func NewUserService(repo IRepository) *Service {
	return &Service{repo: repo}
}
