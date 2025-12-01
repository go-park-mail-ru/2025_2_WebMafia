package http

import (
	"errors"
	"net/http"

	"spotify/microservices/auth/service"
	"spotify/pkg/response"
)

func handleServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, service.ErrValidation):
		response.BadRequestJSON(w)
	case errors.Is(err, service.ErrConflict):
		response.ConflictJSON(w)
	default:
		response.InternalErrorJSON(w)
	}
}
