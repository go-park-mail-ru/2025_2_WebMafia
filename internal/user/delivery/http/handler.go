package http

import (
	"context"
	"spotify/internal/user/dto"
	"spotify/pkg/jwtmanager"
)

type IService interface {
	Register(ctx context.Context, req dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error)
	UploadAvatar(ctx context.Context, req dto.UploadAvatarRequest) (*dto.UploadAvatarResponse, error)
	DeleteAvatar(ctx context.Context, req dto.DeleteAvatarRequest) error
	UpdateProfile(ctx context.Context, req dto.UpdateProfileRequest) (*dto.UpdateProfileResponse, error)
}

type CSRFManager interface {
	Generate(userID, sessionID string) (string, error)
}

type Handler struct {
	svc         IService
	jwtManager  *jwtmanager.Manager
	csrfManager CSRFManager
}

func NewHandler(svc IService, jwtManager *jwtmanager.Manager, csrfManager CSRFManager) *Handler {
	return &Handler{
		svc:         svc,
		jwtManager:  jwtManager,
		csrfManager: csrfManager,
	}
}
