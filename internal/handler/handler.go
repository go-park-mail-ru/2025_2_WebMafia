package handler

import (
	"spotify/internal/service"
)

type Handlers struct {
	authService service.AuthService
}

func NewHandler(authService service.AuthService) *Handlers {
	return &Handlers{
		authService: authService,
	}
}
