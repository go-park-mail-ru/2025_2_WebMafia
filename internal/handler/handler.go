package handler

import (
	"spotify/internal/store"
	"time"
)

type Handlers struct {
	store          *store.MemoryStore
	jwtSecretKey   string
	accessTokenTTL time.Duration
}

func NewHandler(store *store.MemoryStore, jwtSecretKey string, accessTokenTTL time.Duration) *Handlers {
	return &Handlers{
		store:          store,
		jwtSecretKey:   jwtSecretKey,
		accessTokenTTL: accessTokenTTL,
	}
}
