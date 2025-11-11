package service

import (
	"context"
	albumService "spotify/internal/album/service"
	artistService "spotify/internal/artist/service"
	"spotify/internal/model"
	"spotify/internal/track/dto"

	"github.com/google/uuid"
)

//go:generate mockgen -destination=../../mocks/track/repository_mock.go -package=track spotify/internal/track/service IRepository
type IRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*model.Track, error)
	GetAll(ctx context.Context, limit, offset uint64) ([]model.Track, error)
	GetByArtistID(ctx context.Context, artistID uuid.UUID, limit, offset uint64) ([]model.Track, error)
	GetByAlbumID(ctx context.Context, albumID uuid.UUID, limit, offset uint64) ([]model.Track, error)
	GetByGenreID(ctx context.Context, genreID uuid.UUID, limit, offset uint64) ([]model.Track, error)
	Search(ctx context.Context, query string, limit uint64) ([]dto.SearchResult, error)

	IncrementPlayCount(ctx context.Context, trackID uuid.UUID) error

	GetTotalPlaysByArtistID(ctx context.Context, artistID uuid.UUID) (int64, error)
	GetTotalPlaysByArtistIDs(ctx context.Context, artistIDs []uuid.UUID) (map[uuid.UUID]int64, error)
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

func (s *Service) SetAlbumService(albumService *albumService.Service) {
	s.albumService = albumService
}

func (s *Service) SetArtistService(artistService *artistService.Service) {
	s.artistService = artistService
}
