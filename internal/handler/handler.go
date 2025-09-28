package handler

import (
	"context"
	"spotify/internal/model"
	"spotify/pkg/jwtmanager"

	"github.com/google/uuid"
)

const sessionTokenCookie = "session_token"

type storege interface {
	CreateUser(ctx context.Context, user model.User) (*model.User, error)
	GetUserByLogin(ctx context.Context, login string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error)

	GetAllTracks() ([]model.Track, error)
	GetAllArtists() ([]model.Artist, error)
	GetAllAlbums() ([]model.Album, error)
}

type Handlers struct {
	store      storege
	jwtManager *jwtmanager.Manager
}

func NewHandler(store storege, jwtManager *jwtmanager.Manager) *Handlers {
	return &Handlers{
		store:      store,
		jwtManager: jwtManager,
	}
}
