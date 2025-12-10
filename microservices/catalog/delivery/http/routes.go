package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) RegisterRoutes(public *mux.Router, protected *mux.Router) {
	// Artist
	public.HandleFunc("/artists/search", h.SearchArtists).Methods(http.MethodGet, http.MethodOptions)
	public.HandleFunc("/artists", h.GetAllArtists).Methods(http.MethodGet, http.MethodOptions)
	public.HandleFunc("/artists/{id}", h.GetArtistByID).Methods(http.MethodGet, http.MethodOptions)

	// Album
	public.HandleFunc("/albums/search", h.SearchAlbums).Methods(http.MethodGet, http.MethodOptions)
	public.HandleFunc("/albums", h.GetAllAlbums).Methods(http.MethodGet, http.MethodOptions)
	public.HandleFunc("/albums/{id}", h.GetAlbumByID).Methods(http.MethodGet, http.MethodOptions)
	public.HandleFunc("/artists/{artistId}/albums", h.GetAlbumsByArtistID).Methods(http.MethodGet, http.MethodOptions)

	// Track
	public.HandleFunc("/tracks/search", h.SearchTracks).Methods(http.MethodGet, http.MethodOptions)
	public.HandleFunc("/tracks", h.GetAllTracks).Methods(http.MethodGet, http.MethodOptions)
	public.HandleFunc("/tracks/{id}", h.GetTrackByID).Methods(http.MethodGet, http.MethodOptions)
	protected.HandleFunc("/tracks/{id}/listen", h.RegisterPlay).Methods(http.MethodPost, http.MethodOptions)
	public.HandleFunc("/artists/{artistId}/tracks", h.GetTracksByArtist).Methods(http.MethodGet, http.MethodOptions)
	public.HandleFunc("/albums/{albumId}/tracks", h.GetTracksByAlbum).Methods(http.MethodGet, http.MethodOptions)
	public.HandleFunc("/genres/{genreId}/tracks", h.GetTracksByGenre).Methods(http.MethodGet, http.MethodOptions)

	// Comments
	public.HandleFunc("/tracks/{id}/comments", h.GetTrackComments).Methods(http.MethodGet, http.MethodOptions)
	protected.HandleFunc("/tracks/{id}/comments/ws", h.ServeWS).Methods(http.MethodGet, http.MethodOptions)
}
