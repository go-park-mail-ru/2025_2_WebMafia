package service

import (
	"context"
	"spotify/internal/model"

	"github.com/google/uuid"
)

type IRepository interface {
	CreatePlaylist(ctx context.Context, playlist model.Playlist, userID uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Playlist, error)
	GetAllByUser(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]model.Playlist, error)
	UpdatePlaylist(ctx context.Context, id uuid.UUID, title *string, description *string, isFavorite *bool) error
	DeletePlaylist(ctx context.Context, id uuid.UUID) error
	GetFavoritePlaylist(ctx context.Context, userID uuid.UUID) (*model.Playlist, error)
	AddTrackToPlaylist(ctx context.Context, playlistID uuid.UUID, trackID string) error
}

type Service struct {
	repo IRepository
}

func New(repo IRepository) *Service {
	return &Service{repo: repo}
}
