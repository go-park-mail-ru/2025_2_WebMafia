package http

import (
	"context"
	"spotify/microservices/catalog/dto"
	"spotify/pkg/ws"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

//go:generate mockgen -destination=../../../../mocks/catalog/service/http/service_mock.go -package=mock_catalog_service spotify/microservices/catalog/delivery/http IService
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

	PostComment(ctx context.Context, userID uuid.UUID, req dto.PostCommentRequest) (*dto.Comment, error)
	GetCommentsByTrackID(ctx context.Context, trackID uuid.UUID, limit, offset uint64) ([]dto.Comment, error)
}

type Handler struct {
	service    IService
	hub        *ws.Hub
	wsUpgrader websocket.Upgrader
	wsConfig   ws.Config
}

func NewHandler(service IService, hub *ws.Hub, wsConfig ws.Config, allowedOrigins []string) *Handler {
	upgrader := ws.NewUpgrader(allowedOrigins, wsConfig)

	return &Handler{
		service:    service,
		hub:        hub,
		wsUpgrader: upgrader,
		wsConfig:   wsConfig,
	}
}
