package service

import (
	"context"
	"spotify/internal/model"
)

type IRepository interface {
	CreateUser(ctx context.Context, user model.User) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByLogin(ctx context.Context, login string) (*model.User, error)
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
	UpdateUserAvatar(ctx context.Context, userID string, avatarPath string) error
}

type IStorage interface {
	UploadAvatar(ctx context.Context, objectName string, data []byte, contentType string) error
	DeleteAvatar(ctx context.Context, objectName string) error
	GetAvatarURL(ctx context.Context, objectName string) (string, error)
}

type Service struct {
	repo    IRepository
	storage IStorage
}

func NewUserService(repo IRepository, storage IStorage) *Service {
	return &Service{repo: repo, storage: storage}
}
