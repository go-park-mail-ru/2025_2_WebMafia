package handler

import (
	"spotify/internal/store"
	"spotify/pkg/jwtmanager"
)

type Handlers struct {
	store      store.Store
	jwtManager *jwtmanager.Manager
}

func NewHandler(store store.Store, jwtManager *jwtmanager.Manager) *Handlers {
	return &Handlers{
		store:      store,
		jwtManager: jwtManager,
	}
}
