package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"spotify/internal/user/dto"
	"spotify/internal/user/model"
	"spotify/internal/user/tools"
	"time"
)

func (s *Service) Register(ctx context.Context, req dto.RegisterRequest) (*dto.RegisterResponse, error) {

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

	return &dto.RegisterResponse{
		ID:    user.ID.String(),
		Login: user.Login,
		Email: user.Email,
	}, nil
}
