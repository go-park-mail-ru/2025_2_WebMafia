package handler

import (
	"context"
	"spotify/internal/model"
	model2 "spotify/internal/user/model"
	"spotify/pkg/jwtmanager"

	"github.com/google/uuid"
)

const sessionTokenCookie = "session_token"

type storage interface {
	CreateUser(ctx context.Context, user model2.User) (*model2.User, error)
	GetUserByLogin(ctx context.Context, login string) (*model2.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model2.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*model2.User, error)

	GetAllTracks() ([]model.Track, error)
	GetAllArtists() ([]model.Artist, error)
	GetAllAlbums() ([]model.Album, error)
}

type Handlers struct {
	store      storage
	jwtManager *jwtmanager.Manager
}

func NewHandler(store storage, jwtManager *jwtmanager.Manager) *Handlers {
	return &Handlers{
		store:      store,
		jwtManager: jwtManager,
	}
}
