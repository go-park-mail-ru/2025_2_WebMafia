package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/tracks", h.GetAllTracks).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/tracks/{id}", h.GetTrackByID).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/artists/{artistId}/tracks", h.GetTracksByArtist).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/albums/{albumId}/tracks", h.GetTracksByAlbum).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/genres/{genreId}/tracks", h.GetTracksByGenre).Methods(http.MethodGet, http.MethodOptions)
}
