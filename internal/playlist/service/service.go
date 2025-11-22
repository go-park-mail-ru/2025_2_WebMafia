package service

import (
	"context"
	"io"
	"spotify/internal/model"
	"spotify/internal/playlist/repository/postgres"

	"github.com/google/uuid"
)

type IRepository interface {
	CreatePlaylist(ctx context.Context, playlist model.Playlist, userID uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Playlist, error)
	GetAllByUser(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]model.Playlist, error)
	UpdatePlaylist(ctx context.Context, id uuid.UUID, upd postgres.PlaylistUpdate) error
	DeletePlaylist(ctx context.Context, id uuid.UUID) error
	GetFavoritePlaylist(ctx context.Context, userID uuid.UUID) (*model.Playlist, error)
	AddTrackToPlaylist(ctx context.Context, playlistID uuid.UUID, trackID string) error
	UpdatePlaylistAvatar(ctx context.Context, id uuid.UUID, avatar string) error
}

type IStorage interface {
	UploadAvatar(ctx context.Context, file io.Reader, size int64, contentType string) (string, error)
	DeleteAvatar(ctx context.Context, objectName string) error
}

type Service struct {
	repo    IRepository
	storage IStorage
}

func New(repo IRepository, storage IStorage) *Service {
	return &Service{repo: repo, storage: storage}
}
