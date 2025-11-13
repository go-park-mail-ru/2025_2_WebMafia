package http

import (
	"context"
	"spotify/internal/playlist/dto"
)

type IService interface {
	CreatePlaylist(ctx context.Context, req dto.CreatePlaylistRequest) (*dto.Playlist, error)
	GetPlaylist(ctx context.Context, req dto.GetPlaylistRequest) (*dto.Playlist, error)
	GetPlaylistsByUser(ctx context.Context, req dto.GetPlaylistsByUserRequest) ([]dto.Playlist, error)
	UpdatePlaylist(ctx context.Context, req dto.UpdatePlaylistRequest) (*dto.Playlist, error)
	DeletePlaylist(ctx context.Context, req dto.DeletePlaylistRequest) error
	GetFavoritePlaylist(ctx context.Context, req dto.GetFavoritePlaylistRequest) (*dto.Playlist, error)
	AddTrackToFavorite(ctx context.Context, req dto.AddTrackToFavoriteRequest) error
}

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{
		service: service,
	}
}
