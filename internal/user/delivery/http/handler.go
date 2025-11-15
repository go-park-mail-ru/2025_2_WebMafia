package http

import (
	"context"
	"spotify/internal/user/dto"
	"spotify/pkg/jwtmanager"
)

//go:generate mockgen -destination=../../../mocks/user/service_mock.go -package=user spotify/internal/user/delivery/http IService,CSRFManager
type IService interface {
	Register(ctx context.Context, req dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error)
	UploadAvatar(ctx context.Context, req dto.UploadAvatarRequest) (*dto.UploadAvatarResponse, error)
	DeleteAvatar(ctx context.Context, req dto.DeleteAvatarRequest) error
	UpdateProfile(ctx context.Context, req dto.UpdateProfileRequest) (*dto.UpdateProfileResponse, error)
	GetProfile(ctx context.Context, req dto.GetProfileRequest) (*dto.GetProfileResponse, error)
	UpdateRole(ctx context.Context, req dto.UpdateRoleRequest) (*dto.UpdateRoleResponse, error)
}

type CSRFManager interface {
	Generate(userID, sessionID string) (string, error)
}

type Handler struct {
	svc                IService
	jwtManager         *jwtmanager.Manager
	csrfManager        CSRFManager
	allowedAvatarTypes []string
}

func NewHandler(svc IService, jwtManager *jwtmanager.Manager, csrfManager CSRFManager, allowedAvatarTypes []string) *Handler {
	return &Handler{
		svc:                svc,
		jwtManager:         jwtManager,
		csrfManager:        csrfManager,
		allowedAvatarTypes: allowedAvatarTypes,
	}
}
