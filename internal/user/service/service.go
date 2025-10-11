package service

import (
	"context"
	"fmt"
	"spotify/internal/user/model"
	"spotify/internal/user/tools"
	"time"

	"github.com/google/uuid"
)

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
		return nil, fmt.Errorf("failed to hash password: %w", ErrInternal)
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
		return nil, mapRepositoryError(err)
	}

	return &RegisterResponse{
		ID:    user.ID.String(),
		Login: user.Login,
		Email: user.Email,
	}, nil
}
