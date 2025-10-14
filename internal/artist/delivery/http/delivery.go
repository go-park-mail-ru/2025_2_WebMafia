package http

import (
	"context"
	"spotify/internal/artist/dto"

	"github.com/google/uuid"
)

type IService interface {
	GetArtistByID(ctx context.Context, id uuid.UUID) (*dto.Artist, error)
	GetAllArtists(ctx context.Context) ([]dto.Artist, error)
}

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{
		service: service,
	}
}
