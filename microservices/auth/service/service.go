package service

import (
	"context"
	"io"
	"spotify/internal/model"
)

//go:generate mockgen -destination=../../../mocks/auth/repository/repository_mock.go -package=mock_auth_repo spotify/microservices/auth/service IRepository
type IRepository interface {
	CreateUser(ctx context.Context, user model.User) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByLogin(ctx context.Context, login string) (*model.User, error)
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
	UpdateUserAvatar(ctx context.Context, userID string, avatarPath string) error
	UpdateUserProfile(ctx context.Context, user model.User) error
	GetUsersByIDs(ctx context.Context, ids []string) ([]model.User, error)
}

//go:generate mockgen -destination=../../../mocks/pkg/storage/storage_mock.go -package=mock_storage spotify/microservices/auth/service IStorage
type IStorage interface {
	UploadAvatar(ctx context.Context, file io.Reader, size int64, contentType string) (string, error)
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
