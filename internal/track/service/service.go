package service

import (
	"context"
	"spotify/internal/track/model"

	"github.com/google/uuid"
)

type IRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*model.Track, *model.Album, []model.Artist, []model.Genre, error)
	GetAll(ctx context.Context) ([]model.Track, []model.Album, [][]model.Artist, [][]model.Genre, error)
	GetByArtistID(ctx context.Context, artistID uuid.UUID) ([]model.Track, []model.Album, [][]model.Artist, [][]model.Genre, error)
	GetByAlbumID(ctx context.Context, albumID uuid.UUID) ([]model.Track, []model.Album, [][]model.Artist, [][]model.Genre, error)
	GetByGenreID(ctx context.Context, genreID uuid.UUID) ([]model.Track, []model.Album, [][]model.Artist, [][]model.Genre, error)
}

type Service struct {
	repo IRepository
}

func New(repo IRepository) *Service {
	return &Service{
		repo: repo,
	}
}
