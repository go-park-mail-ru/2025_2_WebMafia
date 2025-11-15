package http

import (
	"context"
	"spotify/internal/artist/dto"

	"github.com/google/uuid"
)

//go:generate mockgen -destination=../../../mocks/artist/service_mock.go -package=artist spotify/internal/artist/delivery/http IService
type IService interface {
	GetArtistByID(ctx context.Context, id uuid.UUID) (*dto.Artist, error)
	GetAllArtists(ctx context.Context, limit, offset uint64) ([]dto.Artist, error)
	Search(ctx context.Context, query string, limit uint64) ([]dto.ArtistSearch, error)
}

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{
		service: service,
	}
}
