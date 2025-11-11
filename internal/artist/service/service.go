package service

import (
	"context"
	"spotify/internal/artist/dto"
	"spotify/internal/model"

	"github.com/google/uuid"
)

//go:generate mockgen -destination=../../mocks/artist/repository_mock.go -package=artist spotify/internal/artist/service IRepository
type IRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*model.Artist, error)
	GetAll(ctx context.Context, limit, offset uint64) ([]model.Artist, error)
	GetByIDs(ctx context.Context, ids []uuid.UUID) ([]model.Artist, error)
	Search(ctx context.Context, query string, limit uint64) ([]dto.SearchResult, error)
}

//go:generate mockgen -destination=../../mocks/artist/track_service_mock.go -package=artist spotify/internal/artist/service ITrackService
type ITrackService interface {
	GetTotalPlaysByArtistID(ctx context.Context, artistID uuid.UUID) (int64, error)
	GetTotalPlaysByArtistIDs(ctx context.Context, artistIDs []uuid.UUID) (map[uuid.UUID]int64, error)
}

type Service struct {
	repo         IRepository
	trackService ITrackService
}

func New(repo IRepository, trackService ITrackService) *Service {
	return &Service{
		repo:         repo,
		trackService: trackService,
	}
}

func (s *Service) SetTrackService(ts ITrackService) {
	s.trackService = ts
}
