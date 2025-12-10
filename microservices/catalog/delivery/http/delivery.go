package http

import (
	"context"
	"net/http"
	"spotify/microservices/catalog/dto"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	wsReadBufferSize  = 1024
	wsWriteBufferSize = 1024
)

//go:generate mockgen -destination=../../mocks/service/service_mock.go -package=service_mock spotify/microservices/catalog/delivery/http IService
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
	hub        *Hub
	wsUpgrader websocket.Upgrader
}

func NewHandler(service IService, hub *Hub, allowedOrigins []string) *Handler {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  wsReadBufferSize,
		WriteBufferSize: wsWriteBufferSize,
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			for _, allowed := range allowedOrigins {
				if allowed == origin {
					return true
				}
			}
			return false
		},
	}

	return &Handler{
		service:    service,
		hub:        hub,
		wsUpgrader: upgrader,
	}
}
