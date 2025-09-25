package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"spotify/internal/service"
	"spotify/pkg/response"
)

func (h *Handlers) GetAllTracksHandler(w http.ResponseWriter, r *http.Request) {
	tracks, _ := h.store.GetAllTracks()
	artistID := r.URL.Query().Get("artist_id")
	tracks = service.FilterTracksByArtist(tracks, artistID)
	response.JSON(w, http.StatusOK, map[string]interface{}{
		"status": 200,
		"body":   map[string]interface{}{"tracks": tracks},
	})
}

func (h *Handlers) GetTrackByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trackID := vars["id"]
	tracks, _ := h.store.GetAllTracks()
	for _, t := range tracks {
		if t.TrackID == trackID {
			response.JSON(w, http.StatusOK, map[string]interface{}{
				"status": 200,
				"body":   map[string]interface{}{"track": t},
			})
			return
		}
	}
	response.JSON(w, http.StatusOK, map[string]interface{}{
		"status": 404,
		"body":   map[string]interface{}{},
		"error":  "track not found",
	})
}
