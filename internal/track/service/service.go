package service

import (
	"context"
	albumService "spotify/internal/album/service"
	artistService "spotify/internal/artist/service"
	"spotify/internal/model"

	"github.com/google/uuid"
)

//go:generate mockgen -destination=../../mocks/track/repository_mock.go -package=track spotify/internal/track/service IRepository
type IRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*model.Track, error)
	GetAll(ctx context.Context, limit, offset uint64) ([]model.Track, error)
	GetByArtistID(ctx context.Context, artistID uuid.UUID, limit, offset uint64) ([]model.Track, error)
	GetByAlbumID(ctx context.Context, albumID uuid.UUID, limit, offset uint64) ([]model.Track, error)
	GetByGenreID(ctx context.Context, genreID uuid.UUID, limit, offset uint64) ([]model.Track, error)

	GetAlbumIDsForTracks(ctx context.Context, trackIDs []uuid.UUID) (map[uuid.UUID]uuid.UUID, error)
	GetArtistIDsForTracks(ctx context.Context, trackIDs []uuid.UUID) (map[uuid.UUID][]uuid.UUID, error)
	GetGenresForTracks(ctx context.Context, trackIDs []uuid.UUID) (map[uuid.UUID][]model.Genre, error)
}

type Service struct {
	trackRepo     IRepository
	albumService  *albumService.Service
	artistService *artistService.Service
}

func New(repo IRepository, albumService *albumService.Service, artistService *artistService.Service) *Service {
	return &Service{
		trackRepo:     repo,
		albumService:  albumService,
		artistService: artistService,
	}
}
