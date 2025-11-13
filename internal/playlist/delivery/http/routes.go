package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) RegisterRoutes(public, protected *mux.Router) {
	protected.HandleFunc("/playlists/favorite", h.GetFavoritePlaylist).Methods(http.MethodGet, http.MethodOptions)
	protected.HandleFunc("/playlists/favorite/add-track", h.AddTrackToFavorite).Methods(http.MethodPost)

	public.HandleFunc("/playlists/{id}", h.GetPlaylistByID).Methods(http.MethodGet, http.MethodOptions)
	public.HandleFunc("/users/{userId}/playlists", h.GetAllPlaylistsByUserID).Methods(http.MethodGet, http.MethodOptions)

	protected.HandleFunc("/playlists", h.CreatePlaylist).Methods(http.MethodPost, http.MethodOptions)
	protected.HandleFunc("/playlists/{id}", h.UpdatePlaylist).Methods(http.MethodPut, http.MethodOptions)
	protected.HandleFunc("/playlists/{id}", h.DeletePlaylist).Methods(http.MethodDelete, http.MethodOptions)

}
