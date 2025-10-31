package service

import (
	"context"
	"spotify/internal/model"

	"github.com/google/uuid"
)

//go:generate mockgen -destination=../../mocks/artist/repository_mock.go -package=artist spotify/internal/artist/service IRepository
type IRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*model.Artist, error)
	GetAll(ctx context.Context, limit, offset uint64) ([]model.Artist, error)
	GetByIDs(ctx context.Context, ids []uuid.UUID) ([]model.Artist, error)
}

type Service struct {
	repo IRepository
}

func New(repo IRepository) *Service {
	return &Service{
		repo: repo,
	}
}
