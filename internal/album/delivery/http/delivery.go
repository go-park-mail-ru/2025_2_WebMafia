package http

import (
	"context"
	"spotify/internal/album/dto"

	"github.com/google/uuid"
)

type IService interface {
	GetAlbumByID(ctx context.Context, id uuid.UUID) (*dto.Album, error)
	GetAllAlbums(ctx context.Context) ([]dto.Album, error)
}

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{
		service: service,
	}
}
