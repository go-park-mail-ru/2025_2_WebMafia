package http

import (
	"context"
	"spotify/internal/album/dto"

	"github.com/google/uuid"
)

type IService interface {
	GetAlbumByID(ctx context.Context, id uuid.UUID) (*dto.Album, error)
	GetAllAlbums(ctx context.Context, limit, offset uint64) ([]dto.Album, error)
	GetAlbumsByArtistID(ctx context.Context, artistID uuid.UUID, limit, offset uint64) ([]dto.Album, error)
}

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{
		service: service,
	}
}
