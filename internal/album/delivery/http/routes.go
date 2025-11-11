package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/albums", h.GetAllAlbums).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/albums/{id}", h.GetAlbumByID).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/artists/{artistId}/albums", h.GetAlbumsByArtistID).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/albums/search", h.Search).Methods(http.MethodGet, http.MethodOptions)
}
