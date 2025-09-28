package handler

import (
	"fmt"
	"net/http"
	"spotify/internal/model"
	"spotify/pkg/response"

	"github.com/gorilla/mux"
)

type TracksResponse struct {
	Tracks []model.Track `json:"tracks"`
}

type TrackResponse struct {
	Track model.Track `json:"track"`
}

func (h *Handlers) GetAllTracksHandler(w http.ResponseWriter, r *http.Request) {
	tracks, _ := h.store.GetAllTracks()
	response.JSON(w, http.StatusOK, TracksResponse{Tracks: tracks})
}

func (h *Handlers) GetTrackByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trackID := vars["id"]

	tracks, _ := h.store.GetAllTracks()
	for _, t := range tracks {
		if fmt.Sprint(t.TrackID) == trackID {
			response.JSON(w, http.StatusOK, TrackResponse{Track: t})
			return
		}
	}
	response.JSON(w, http.StatusNotFound, fmt.Errorf("track not found"))

}
