package handler

import (
	"spotify/internal/mock"
	"spotify/internal/model"
	"sync"
)

type Handlers struct {
	users []model.User
	mu    *sync.RWMutex
	store *mock_store.MockStore
}

func NewHandler() *Handlers {
	return &Handlers{
		users: make([]model.User, 0),
		mu:    &sync.RWMutex{},
		store: mock_store.NewMockStore(),
	}
}
