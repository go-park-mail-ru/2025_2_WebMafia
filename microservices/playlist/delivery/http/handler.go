package http

import (
	"context"
	"spotify/microservices/playlist/dto"

	"github.com/google/uuid"
)

//go:generate mockgen -destination=../../../../mocks/playlist/service/service_mock.go -package=mock_playlist_service spotify/microservices/playlist/delivery/http IService
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
	AddAlbumToFavorite(ctx context.Context, req dto.AddAlbumToFavoriteRequest) error
	RemoveAlbumFromFavorite(ctx context.Context, req dto.RemoveAlbumFromFavoriteRequest) error
	GetFavoriteAlbums(ctx context.Context, userID uuid.UUID) ([]dto.FavoriteAlbum, error)
	AddArtistToFavorite(ctx context.Context, req dto.AddArtistToFavoriteRequest) error
	RemoveArtistFromFavorite(ctx context.Context, req dto.RemoveArtistFromFavoriteRequest) error
	GetFavoriteArtists(ctx context.Context, userID uuid.UUID) ([]dto.FavoriteArtist, error)
	GeneratePlaylistMeta(ctx context.Context, playlistID uuid.UUID) (*dto.GeneratedMeta, error)
	ConfirmPlaylistMeta(ctx context.Context, playlistID uuid.UUID, title, description string) error
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
