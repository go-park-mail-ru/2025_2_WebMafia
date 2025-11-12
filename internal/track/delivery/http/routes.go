package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) RegisterRoutes(public, protected, csrfProtected *mux.Router) {
	public.HandleFunc("/tracks/search", h.Search).Methods(http.MethodGet, http.MethodOptions)
	public.HandleFunc("/tracks", h.GetAllTracks).Methods(http.MethodGet, http.MethodOptions)
	public.HandleFunc("/tracks/{id}", h.GetTrackByID).Methods(http.MethodGet, http.MethodOptions)
	csrfProtected.HandleFunc("/tracks/{id}/listen", h.RegisterPlay).Methods(http.MethodPost, http.MethodOptions)
	public.HandleFunc("/artists/{artistId}/tracks", h.GetTracksByArtist).Methods(http.MethodGet, http.MethodOptions)
	public.HandleFunc("/albums/{albumId}/tracks", h.GetTracksByAlbum).Methods(http.MethodGet, http.MethodOptions)
	public.HandleFunc("/genres/{genreId}/tracks", h.GetTracksByGenre).Methods(http.MethodGet, http.MethodOptions)
}
