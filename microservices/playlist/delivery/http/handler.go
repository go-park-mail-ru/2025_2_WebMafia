package http

import (
	"context"
	"spotify/microservices/playlist/dto"

	"github.com/google/uuid"
)

//go:generate mockgen -destination=../../mocks/service/service_mock.go -package=service_mock spotify/microservices/playlist/delivery/http IService
type IService interface {
	CreatePlaylist(ctx context.Context, req dto.CreatePlaylistRequest) (*dto.Playlist, error)
	GetPlaylistsByUser(ctx context.Context, req dto.GetPlaylistsByUserRequest) ([]dto.Playlist, error)
	UpdatePlaylist(ctx context.Context, req dto.UpdatePlaylistRequest) (*dto.Playlist, error)
	DeletePlaylist(ctx context.Context, req dto.DeletePlaylistRequest) error
	GetFavoritePlaylist(ctx context.Context, req dto.GetFavoritePlaylistRequest) (*dto.Playlist, error)
	AddTrackToFavorite(ctx context.Context, req dto.AddTrackToFavoriteRequest) error
	UploadPlaylistAvatar(ctx context.Context, req dto.UploadPlaylistAvatarRequest) (*dto.UploadPlaylistAvatarResponse, error)
	DeletePlaylistAvatar(ctx context.Context, req dto.DeletePlaylistAvatarRequest) error
	GetPlaylistWithTracks(ctx context.Context, id uuid.UUID) (*dto.Playlist, error)
	AddTrackToPlaylist(ctx context.Context, req dto.AddTrackToPlaylistRequest) error
	RemoveTrackFromPlaylist(ctx context.Context, req dto.RemoveTrackFromPlaylistRequest) error
}

type Handler struct {
	service            IService
	allowedAvatarTypes []string
}

func NewHandler(service IService, allowedAvatarTypes []string) *Handler {
	return &Handler{
		service:            service,
		allowedAvatarTypes: allowedAvatarTypes,
	}
}
