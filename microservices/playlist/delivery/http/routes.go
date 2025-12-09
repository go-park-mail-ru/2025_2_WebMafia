package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) RegisterRoutes(public, protected, csrfProtected *mux.Router) {
	protected.HandleFunc("/playlists/my", h.GetMyPlaylists).Methods(http.MethodGet, http.MethodOptions)
	protected.HandleFunc("/playlists/favorite", h.GetFavoritePlaylist).Methods(http.MethodGet, http.MethodOptions)
	protected.HandleFunc("/favorite/albums", h.GetFavoriteAlbums).Methods(http.MethodGet, http.MethodOptions)
	protected.HandleFunc("/favorite/artists", h.GetFavoriteArtists).Methods(http.MethodGet, http.MethodOptions)

	public.HandleFunc("/playlists/{id}", h.GetPlaylistByID).Methods(http.MethodGet, http.MethodOptions)

	csrfProtected.HandleFunc("/playlists/favorite/add-track", h.AddTrackToFavorite).Methods(http.MethodPost, http.MethodOptions)
	csrfProtected.HandleFunc("/playlists", h.CreatePlaylist).Methods(http.MethodPost, http.MethodOptions)
	csrfProtected.HandleFunc("/playlists/{id}", h.UpdatePlaylist).Methods(http.MethodPut, http.MethodOptions)
	csrfProtected.HandleFunc("/playlists/{id}", h.DeletePlaylist).Methods(http.MethodDelete, http.MethodOptions)
	csrfProtected.HandleFunc("/playlists/{id}/avatar", h.UploadPlaylistAvatar).Methods(http.MethodPost, http.MethodOptions)
	csrfProtected.HandleFunc("/playlists/{id}/avatar", h.DeletePlaylistAvatar).Methods(http.MethodDelete, http.MethodOptions)
	csrfProtected.HandleFunc("/playlists/{id}/tracks", h.AddTrackToPlaylist).Methods(http.MethodPost, http.MethodOptions)
	csrfProtected.HandleFunc("/playlists/{id}/tracks", h.RemoveTrackFromPlaylist).Methods(http.MethodDelete, http.MethodOptions)
	csrfProtected.HandleFunc("/favorite/albums/{id}", h.AddAlbumToFavorite).Methods(http.MethodPost, http.MethodOptions)
	csrfProtected.HandleFunc("/favorite/albums/{id}", h.RemoveAlbumFromFavorite).Methods(http.MethodDelete, http.MethodOptions)
	csrfProtected.HandleFunc("/favorite/artists/{id}", h.AddArtistToFavorite).Methods(http.MethodPost, http.MethodOptions)
	csrfProtected.HandleFunc("/favorite/artists/{id}", h.RemoveArtistFromFavorite).Methods(http.MethodDelete, http.MethodOptions)
}
