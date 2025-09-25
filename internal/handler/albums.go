package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"spotify/internal/service"
	"spotify/pkg/response"
)

func (h *Handlers) GetAllAlbumsHandler(w http.ResponseWriter, r *http.Request) {
	albums, _ := h.store.GetAllAlbums()
	artistID := r.URL.Query().Get("artist_id")
	albums = service.FilterAlbumsByArtist(albums, artistID)
	response.JSON(w, http.StatusOK, map[string]interface{}{
		"status": 200,
		"body":   map[string]interface{}{"albums": albums},
	})
}

func (h *Handlers) GetAlbumByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	albumID := vars["id"]
	albums, _ := h.store.GetAllAlbums()
	for _, album := range albums {
		if album.AlbumID == albumID {
			response.JSON(w, http.StatusOK, map[string]interface{}{
				"status": 200,
				"body":   map[string]interface{}{"album": album},
			})
			return
		}
	}
	response.JSON(w, http.StatusOK, map[string]interface{}{
		"status": 404,
		"body":   map[string]interface{}{},
		"error":  "album not found",
	})
}
