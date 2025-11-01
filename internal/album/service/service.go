package service

import (
	"context"
	artistService "spotify/internal/artist/service"
	"spotify/internal/model"

	"github.com/google/uuid"
)

//go:generate mockgen -destination=../../mocks/album/repository_mock.go -package=album spotify/internal/album/service IRepository
type IRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*model.Album, error)
	GetAll(ctx context.Context, limit, offset uint64) ([]model.Album, error)
	GetByIDs(ctx context.Context, ids []uuid.UUID) ([]model.Album, error)
	GetByArtistID(ctx context.Context, artistID uuid.UUID, limit, offset uint64) ([]model.Album, error)
}

type Service struct {
	albumRepo     IRepository
	artistService *artistService.Service
}

func New(repo IRepository, artistService *artistService.Service) *Service {
	return &Service{
		albumRepo:     repo,
		artistService: artistService,
	}
}
