package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) RegisterRoutes(protected *mux.Router) {
	protected.HandleFunc("/csrf-token", h.GetCSRFToken).Methods(http.MethodGet)
}
