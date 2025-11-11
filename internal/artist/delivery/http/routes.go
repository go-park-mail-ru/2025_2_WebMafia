package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/artists", h.GetAllArtists).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/artists/{id}", h.GetArtistByID).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/artists/search", h.Search).Methods(http.MethodGet, http.MethodOptions)
}
