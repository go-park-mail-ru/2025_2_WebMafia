package service

import (
	"context"
	"fmt"
	"spotify/internal/model"
	"spotify/microservices/auth/dto"
	"spotify/microservices/auth/tools"
	"time"

	"github.com/google/uuid"
)

// Register&Auth
func (s *Service) Register(ctx context.Context, req dto.RegisterRequest) (*dto.RegisterResponse, error) {
	const op = "service.Register"
	hash, err := tools.Hash(req.Password)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to hash password: %w", op, err)
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

func (s *Service) GetUsersByIDs(ctx context.Context, ids []string) ([]dto.GetProfileResponse, error) {
	const op = "service.GetUsersByIDs"

	if len(ids) == 0 {
		return []dto.GetProfileResponse{}, nil
	}
	users, err := s.repo.GetUsersByIDs(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("%s: repo error: %w", op, err)
	}

	out := make([]dto.GetProfileResponse, 0, len(users))
	for _, u := range users {
		out = append(out, dto.GetProfileResponse{
			ID:        u.ID.String(),
			Login:     u.Login,
			Email:     u.Email,
			AvatarURL: u.AvatarURL,
		})
	}
	return out, nil
}

// Avatar Upload
func (s *Service) UploadAvatar(ctx context.Context, req dto.UploadAvatarRequest) (*dto.UploadAvatarResponse, error) {
	const op = "service.UpdateAvatar"

	objectName, err := s.storage.UploadAvatar(ctx, req.File, req.Size, req.ContentType)
	if err != nil {
		return nil, err
	}

	if err := s.repo.UpdateUserAvatar(ctx, req.UserID, objectName); err != nil {
		if delErr := s.storage.DeleteAvatar(ctx, objectName); delErr != nil {
			return nil, fmt.Errorf("[%s]: failed to delete uploaded avatar %q): %w", op, objectName, delErr)
		}
		return nil, mapRepositoryError(err)
	}

	return &dto.UploadAvatarResponse{URL: objectName}, nil
}

func (s *Service) DeleteAvatar(ctx context.Context, req dto.DeleteAvatarRequest) error {
	const op = "service.DeleteAvatar"

	user, err := s.repo.GetUserByID(ctx, req.UserID)
	if err != nil {
		return mapRepositoryError(err)
	}

	if user.AvatarURL != "" {
		if err := s.storage.DeleteAvatar(ctx, user.AvatarURL); err != nil {
			return fmt.Errorf("[%s]: delete avatar from storage: %w", op, err)
		}
		if err := s.repo.UpdateUserAvatar(ctx, req.UserID, ""); err != nil {
			return mapRepositoryError(err)
		}

	}

	return nil
}

func (s *Service) UpdateProfile(ctx context.Context, req dto.UpdateProfileRequest) (*dto.UpdateProfileResponse, error) {
	const op = "service.UpdateProfile"

	id, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, fmt.Errorf("[%s]: invalid user ID: %w", op, err)
	}

	user := model.User{
		ID:    id,
		Login: req.Login,
		Email: req.Email,
	}

	if req.Password != "" {
		hash, err := tools.Hash(req.Password)
		if err != nil {
			return nil, fmt.Errorf("[%s]: failed to hash password: %w", op, err)
		}
		user.PasswordHash = hash
	}

	if err := s.repo.UpdateUserProfile(ctx, user); err != nil {
		return nil, mapRepositoryError(err)
	}

	return &dto.UpdateProfileResponse{
		ID:    user.ID.String(),
		Login: user.Login,
		Email: user.Email,
	}, nil
}

func (s *Service) GetProfile(ctx context.Context, req dto.GetProfileRequest) (*dto.GetProfileResponse, error) {
	user, err := s.repo.GetUserByID(ctx, req.UserID)
	if err != nil {
		return nil, mapRepositoryError(err)
	}

	return &dto.GetProfileResponse{
		ID:        user.ID.String(),
		Login:     user.Login,
		Email:     user.Email,
		AvatarURL: user.AvatarURL,
	}, nil
}
