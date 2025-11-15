package http

import (
	"errors"
	"net/http"
	"spotify/internal/ticket/service"
	"spotify/pkg/response"
)

func handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, service.ErrNotFound):
		response.NotFoundJSON(w)
	case errors.Is(err, service.ErrForbidden):
		response.ForbiddenJSON(w)
	case errors.Is(err, service.ErrInvalidStateForAction),
		errors.Is(err, service.ErrInvalidRating):
		response.BadRequestJSON(w)
	default:
		response.InternalErrorJSON(w)
	}
}
