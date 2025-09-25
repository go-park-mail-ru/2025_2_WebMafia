package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"spotify/internal/service"
	"spotify/pkg/response"
)

func (h *Handlers) GetAllArtistsHandler(w http.ResponseWriter, r *http.Request) {
	artists, _ := h.store.GetAllArtists()
	name := r.URL.Query().Get("name")
	artists = service.FilterArtistsByName(artists, name)
	response.JSON(w, http.StatusOK, map[string]interface{}{
		"status": 200,
		"body":   map[string]interface{}{"artists": artists},
	})
}

func (h *Handlers) GetArtistByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	artistID := vars["id"]
	artists, _ := h.store.GetAllArtists()
	for _, a := range artists {
		if a.ArtistID == artistID {
			response.JSON(w, http.StatusOK, map[string]interface{}{
				"status": 200,
				"body":   map[string]interface{}{"artist": a},
			})
			return
		}
	}
	response.JSON(w, http.StatusOK, map[string]interface{}{
		"status": 404,
		"body":   map[string]interface{}{},
		"error":  "artist not found",
	})
}
