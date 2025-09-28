package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"spotify/internal/model"
	"spotify/pkg/response"
)

type TracksResponse struct {
	Tracks []model.Track `json:"tracks"`
}

type TrackResponse struct {
	Track model.Track `json:"track"`
}

func (h *Handlers) GetAllTracksHandler(w http.ResponseWriter, r *http.Request) {
	tracks, _ := h.store.GetAllTracks()
	for i := range tracks {
		enrichTrackURLs(h.cfg, r, &tracks[i])
	}
	response.JSON(w, http.StatusOK, TracksResponse{Tracks: tracks})
}

func (h *Handlers) GetTrackByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trackID := vars["id"]

	tracks, _ := h.store.GetAllTracks()
	for _, t := range tracks {
		if fmt.Sprint(t.TrackID) == trackID {
			enrichTrackURLs(h.cfg, r, &t)
			response.JSON(w, http.StatusOK, TrackResponse{Track: t})
			return
		}
	}
	response.JSON(w, http.StatusNotFound, response.ErrorResponse{Error: "track not found"})

}
