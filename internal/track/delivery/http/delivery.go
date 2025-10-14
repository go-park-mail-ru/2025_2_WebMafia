package http

import (
	"context"
	"spotify/internal/track/dto"

	"github.com/google/uuid"
)

type IService interface {
	GetTrackByID(ctx context.Context, id uuid.UUID) (*dto.Track, error)
	GetAllTracks(ctx context.Context) ([]dto.Track, error)
	GetTracksByArtistID(ctx context.Context, artistID uuid.UUID) ([]dto.Track, error)
	GetTracksByAlbumID(ctx context.Context, albumID uuid.UUID) ([]dto.Track, error)
	GetTracksByGenreID(ctx context.Context, genreID uuid.UUID) ([]dto.Track, error)
}

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{
		service: service,
	}
}
