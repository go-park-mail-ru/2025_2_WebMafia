package handler

import (
	"net/http"
	"spotify/pkg/response"
)

func (h *Handlers) HomeHandler(w http.ResponseWriter, r *http.Request) {
	tracks, _ := h.store.GetAllTracks()
	artists, _ := h.store.GetAllArtists()
	albums, _ := h.store.GetAllAlbums()

	response.JSON(w, http.StatusOK, map[string]interface{}{
		"status": 200,
		"body": map[string]interface{}{
			"tracks":  tracks,
			"artists": artists,
			"albums":  albums,
		},
	})
}
