package service

import (
	"context"
	"spotify/internal/model"
	"spotify/microservices/catalog/dto"

	"github.com/google/uuid"
)

//go:generate mockgen -destination=../mocks/repository/repository_mock.go -package=repository_mock spotify/microservices/catalog/service IRepository
type IRepository interface {
	GetArtistByID(ctx context.Context, id uuid.UUID) (*model.Artist, error)
	GetAllArtists(ctx context.Context, limit, offset uint64) ([]model.Artist, error)
	GetArtistsByIDs(ctx context.Context, ids []uuid.UUID) ([]model.Artist, error)
	SearchArtists(ctx context.Context, query string, limit uint64) ([]dto.ArtistSearchResult, error)

	GetAlbumByID(ctx context.Context, id uuid.UUID) (*model.Album, error)
	GetAllAlbums(ctx context.Context, limit, offset uint64) ([]model.Album, error)
	GetAlbumsByIDs(ctx context.Context, ids []uuid.UUID) ([]model.Album, error)
	GetAlbumsByArtistID(ctx context.Context, artistID uuid.UUID, limit, offset uint64) ([]model.Album, error)
	SearchAlbums(ctx context.Context, query string, limit uint64) ([]dto.AlbumSearchResult, error)

	GetTrackByID(ctx context.Context, id uuid.UUID) (*model.Track, error)
	GetAllTracks(ctx context.Context, limit, offset uint64) ([]model.Track, error)
	GetTracksByArtistID(ctx context.Context, artistID uuid.UUID, limit, offset uint64) ([]model.Track, error)
	GetTracksByAlbumID(ctx context.Context, albumID uuid.UUID, limit, offset uint64) ([]model.Track, error)
	GetTracksByGenreID(ctx context.Context, genreID uuid.UUID, limit, offset uint64) ([]model.Track, error)
	SearchTracks(ctx context.Context, query string, limit uint64) ([]dto.TrackSearchResult, error)
	IncrementPlayCount(ctx context.Context, trackID uuid.UUID) error
	GetTotalPlaysByArtistID(ctx context.Context, artistID uuid.UUID) (int64, error)
	GetTotalPlaysByArtistIDs(ctx context.Context, artistIDs []uuid.UUID) (map[uuid.UUID]int64, error)
	GetAlbumIDsForTracks(ctx context.Context, trackIDs []uuid.UUID) (map[uuid.UUID]uuid.UUID, error)
	GetArtistIDsForTracks(ctx context.Context, trackIDs []uuid.UUID) (map[uuid.UUID][]uuid.UUID, error)
	GetGenresForTracks(ctx context.Context, trackIDs []uuid.UUID) (map[uuid.UUID][]model.Genre, error)
}

type Service struct {
	repo IRepository
}

func New(repo IRepository) *Service {
	return &Service{
		repo: repo,
	}
}
