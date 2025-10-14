package service

import (
	"context"
	"spotify/internal/album/model"

	"github.com/google/uuid"
)

type IRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*model.Album, *model.Artist, error)
	GetAll(ctx context.Context) ([]model.Album, []model.Artist, error)
}

type Service struct {
	repo IRepository
}

func New(repo IRepository) *Service {
	return &Service{
		repo: repo,
	}
}
