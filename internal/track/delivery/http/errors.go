package http

import (
	"errors"
	"log"
	"net/http"
	"spotify/internal/track/service"
	"spotify/pkg/response"
)

func (h *Handler) handleError(w http.ResponseWriter, err error, location string) {
	switch {
	case errors.Is(err, service.ErrNotFound):
		log.Printf("INFO: %s: resource not found: %v", location, err)
		response.NotFoundJSON(w)

	default:
		log.Printf("ERROR: %s: internal server error: %v", location, err)
		response.InternalErrorJSON(w)
	}
}
