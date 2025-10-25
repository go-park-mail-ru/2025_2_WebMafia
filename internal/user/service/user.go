package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"spotify/internal/model"
	"spotify/internal/user/dto"
	"spotify/internal/user/tools"
	"strings"
	"time"
)

// Register&Auth
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
	user, err := s.repo.GetUserByLogin(ctx, req.Login)
	if err != nil {
		return nil, mapRepositoryError(err)
	}

	if err := tools.Compare(user.PasswordHash, req.Password); err != nil {
		return nil, fmt.Errorf("invalid credentials: %w", ErrValidation)
	}

	return &dto.LoginResponse{
		ID: user.ID.String(),
	}, nil
}

// Avatar Upload
func (s *Service) UploadAvatar(ctx context.Context, req dto.UploadAvatarRequest) (*dto.UploadAvatarResponse, error) {
	ext := strings.TrimPrefix(req.ContentType, "image/")
	objectName := fmt.Sprintf("%s.%s", uuid.New().String(), ext)

	if err := s.storage.UploadAvatar(ctx, objectName, req.File, req.Size, req.ContentType); err != nil {
		return nil, ErrInternal
	}

	if err := s.repo.UpdateUserAvatar(ctx, req.UserID, objectName); err != nil {
		if delErr := s.storage.DeleteAvatar(ctx, objectName); delErr != nil {
			return nil, fmt.Errorf("failed to delete uploaded avatar %q): %w", objectName, delErr)
		}
		return nil, mapRepositoryError(err)
	}

	url, err := s.storage.GetAvatarURL(ctx, objectName)
	if err != nil {
		return nil, ErrInternal
	}

	return &dto.UploadAvatarResponse{URL: url}, nil
}

func (s *Service) DeleteAvatar(ctx context.Context, req dto.DeleteAvatarRequest) error {

	user, err := s.repo.GetUserByID(ctx, req.UserID)
	if err != nil {
		return mapRepositoryError(err)
	}

	if user.AvatarURL != "" {
		if err := s.storage.DeleteAvatar(ctx, user.AvatarURL); err != nil {
			return fmt.Errorf("delete avatar from storage: %w", ErrInternal)
		}
		if err := s.repo.UpdateUserAvatar(ctx, req.UserID, ""); err != nil {
			return mapRepositoryError(err)
		}

	}

	return nil
}
