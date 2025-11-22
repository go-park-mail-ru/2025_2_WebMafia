package http

import (
	"context"
	"spotify/microservices/catalog/dto"

	"github.com/google/uuid"
)

type IService interface {
	GetArtistByID(ctx context.Context, id uuid.UUID) (*dto.Artist, error)
	GetAllArtists(ctx context.Context, limit, offset uint64) ([]dto.Artist, error)
	SearchArtists(ctx context.Context, query string, limit uint64) ([]dto.ArtistSearch, error)

	GetAlbumByID(ctx context.Context, id uuid.UUID) (*dto.Album, error)
	GetAllAlbums(ctx context.Context, limit, offset uint64) ([]dto.Album, error)
	GetAlbumsByArtistID(ctx context.Context, artistID uuid.UUID, limit, offset uint64) ([]dto.Album, error)
	SearchAlbums(ctx context.Context, query string, limit uint64) ([]dto.AlbumSearch, error)

	GetTrackByID(ctx context.Context, id uuid.UUID) (*dto.Track, error)
	GetAllTracks(ctx context.Context, limit, offset uint64) ([]dto.Track, error)
	GetTracksByArtistID(ctx context.Context, artistID uuid.UUID, limit, offset uint64) ([]dto.Track, error)
	GetTracksByAlbumID(ctx context.Context, albumID uuid.UUID, limit, offset uint64) ([]dto.Track, error)
	GetTracksByGenreID(ctx context.Context, genreID uuid.UUID, limit, offset uint64) ([]dto.Track, error)
	RegisterPlay(ctx context.Context, trackID uuid.UUID) error
	SearchTracks(ctx context.Context, query string, limit uint64) ([]dto.TrackSearch, error)
}

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{
		service: service,
	}
}
