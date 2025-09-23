package handler

import (
	"spotify/internal/model"
	"sync"
)

type Handlers struct {
	users []model.User
	mu    *sync.RWMutex
}

func NewHandler() *Handlers {
	return &Handlers{
		users: make([]model.User, 0),
		mu:    &sync.RWMutex{},
	}
}
