package service

import (
	"context"
	"io"

	"spotify/internal/model"
	"spotify/microservices/playlist/repository/postgres"

	pbCatalog "spotify/proto/catalog"

	"github.com/google/uuid"
)

//go:generate mockgen -destination=../mocks/repository/repository_mock.go -package=repository_mock spotify/microservices/playlist/service IRepository
type IRepository interface {
	CreatePlaylist(ctx context.Context, playlist model.Playlist, userID uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Playlist, error)
	GetAllByUser(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]model.Playlist, error)
	UpdatePlaylist(ctx context.Context, id uuid.UUID, upd postgres.PlaylistUpdate) error
	DeletePlaylist(ctx context.Context, id uuid.UUID) error
	GetFavoritePlaylist(ctx context.Context, userID uuid.UUID) (*model.Playlist, error)
	AddTrackToPlaylist(ctx context.Context, playlistID uuid.UUID, trackID string) error
	UpdatePlaylistAvatar(ctx context.Context, id uuid.UUID, avatar string) error
	GetTracksByPlaylist(ctx context.Context, playlistID uuid.UUID) ([]string, error)
	RemoveTrackFromPlaylist(ctx context.Context, playlistID uuid.UUID, trackID string) error
	// любимые альбомы
	AddAlbumToFavorite(ctx context.Context, userID uuid.UUID, albumID string) error
	RemoveAlbumFromFavorite(ctx context.Context, userID uuid.UUID, albumID string) error
	GetFavoriteAlbumIDs(ctx context.Context, userID uuid.UUID) ([]string, error)
	// любимые артисты
	AddArtistToFavorite(ctx context.Context, userID uuid.UUID, artistID string) error
	RemoveArtistFromFavorite(ctx context.Context, userID uuid.UUID, artistID string) error
	GetFavoriteArtistIDs(ctx context.Context, userID uuid.UUID) ([]string, error)
}

//go:generate mockgen -destination=../mocks/storage/storage_mock.go -package=storage_mock spotify/microservices/playlist/service IStorage
type IStorage interface {
	UploadAvatar(ctx context.Context, file io.Reader, size int64, contentType string) (string, error)
	DeleteAvatar(ctx context.Context, objectName string) error
}

//go:generate mockgen -destination=../mocks/catalog/catalog_mock.go -package=catalog_mock spotify/proto/catalog CatalogServiceClient
type Service struct {
	repo    IRepository
	storage IStorage
	catalog pbCatalog.CatalogServiceClient
}

func New(repo IRepository, storage IStorage, catalog pbCatalog.CatalogServiceClient) *Service {
	return &Service{repo: repo, storage: storage, catalog: catalog}
}
