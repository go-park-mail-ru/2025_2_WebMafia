package handler

import (
	"spotify/internal/model"
	"spotify/internal/store"
	"sync"
)

type Handlers struct {
	users []model.User
	mu    *sync.RWMutex
	store *memory_store.MockStore
}

func NewHandler() *Handlers {
	return &Handlers{
		users: make([]model.User, 0),
		mu:    &sync.RWMutex{},
		store: memory_store.NewMockStore(),
	}
}
