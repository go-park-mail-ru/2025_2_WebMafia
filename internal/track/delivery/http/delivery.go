package http

import (
	"context"
	"spotify/internal/track/dto"

	"github.com/google/uuid"
)

type IService interface {
	GetTrackByID(ctx context.Context, id uuid.UUID) (*dto.Track, error)
	GetAllTracks(ctx context.Context, limit, offset uint64) ([]dto.Track, error)
	GetTracksByArtistID(ctx context.Context, artistID uuid.UUID, limit, offset uint64) ([]dto.Track, error)
	GetTracksByAlbumID(ctx context.Context, albumID uuid.UUID, limit, offset uint64) ([]dto.Track, error)
	GetTracksByGenreID(ctx context.Context, genreID uuid.UUID, limit, offset uint64) ([]dto.Track, error)
}

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{
		service: service,
	}
}