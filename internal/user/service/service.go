package service

import (
	"context"
	"spotify/internal/user/model"
	"spotify/internal/user/tools"
	"time"

	"github.com/google/uuid"
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

type RegisterRequest struct {
	Login    string
	Email    string
	Password string
}
type RegisterResponse struct {
	ID    string
	Login string
	Email string
}

func (s *Service) Register(ctx context.Context, req RegisterRequest) (*RegisterResponse, error) {

	hash, err := tools.Hash(req.Password)
	if err != nil {
		return nil, NewServiceError(ErrInternal, "failed to hash password")
	}

	user := model.User{
		ID:           uuid.New(),
		Login:        req.Login,
		Email:        req.Email,
		PasswordHash: hash,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return &RegisterResponse{
		ID:    user.ID.String(),
		Login: user.Login,
		Email: user.Email,
	}, nil
}
