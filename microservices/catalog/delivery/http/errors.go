package http

import (
	"errors"
	"net/http"
	"spotify/microservices/catalog/service"
	"spotify/pkg/response"
)

func (h *Handler) handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, service.ErrNotFound):
		response.NotFoundJSON(w)
	default:
		response.InternalErrorJSON(w)
	}
}
