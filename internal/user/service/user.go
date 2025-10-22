package service

import (
	"context"
	"fmt"
	"spotify/internal/model"
	"spotify/internal/user/dto"
	"spotify/internal/user/tools"
	"time"

	"github.com/google/uuid"
)

func (s *Service) Register(ctx context.Context, req dto.RegisterRequest) (*dto.RegisterResponse, error) {
	const op = "service.Register"
	hash, err := tools.Hash(req.Password)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to hash password: %w", op, ErrInternal)
	}

	user := model.User{
		ID:           uuid.New(),
		Login:        req.Login,
		Email:        req.Email,
		PasswordHash: hash,
		AvatarURL:    "",
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

func (s *Service) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	const op = "service.Login"
	user, err := s.repo.GetUserByLogin(ctx, req.Login)
	if err != nil {
		return nil, mapRepositoryError(err)
	}

	if err := tools.Compare(user.PasswordHash, req.Password); err != nil {
		return nil, fmt.Errorf("[%s]: invalid credentials: %w", op, ErrValidation)
	}

	return &dto.LoginResponse{
		ID: user.ID.String(),
	}, nil
}
