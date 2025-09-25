package handler

import (
	"net/http"
	"spotify/pkg/response"
)

func (h *Handlers) HomeHandler(w http.ResponseWriter, r *http.Request) {
	tracks, err := h.store.GetAllTracks()
	if err != nil {
		response.JSON(w, http.StatusOK, map[string]interface{}{
			"status": 500,
			"body":   map[string]interface{}{},
			"error":  "failed to fetch tracks",
		})
		return
	}
	artists, err := h.store.GetAllArtists()
	if err != nil {
		response.JSON(w, http.StatusOK, map[string]interface{}{
			"status": 500,
			"body":   map[string]interface{}{},
			"error":  "failed to fetch artists",
		})
		return
	}
	albums, err := h.store.GetAllAlbums()
	if err != nil {
		response.JSON(w, http.StatusOK, map[string]interface{}{
			"status": 500,
			"body":   map[string]interface{}{},
			"error":  "failed to fetch albums",
		})
		return
	}
	response.JSON(w, http.StatusOK, map[string]interface{}{
		"status": 200,
		"body": map[string]interface{}{
			"tracks":  tracks,
			"artists": artists,
			"albums":  albums,
		},
	})
}
