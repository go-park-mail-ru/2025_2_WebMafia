package handler

import (
	"net/http"
	"spotify/internal/model"
	"spotify/pkg/response"
)

type HomeResponse struct {
	Tracks  []model.Track  `json:"tracks"`
	Artists []model.Artist `json:"artists"`
	Albums  []model.Album  `json:"albums"`
}

func (h *Handlers) HomeHandler(w http.ResponseWriter, r *http.Request) {
	tracks, _ := h.store.GetAllTracks()
	artists, _ := h.store.GetAllArtists()
	albums, _ := h.store.GetAllAlbums()

	response.JSON(w, http.StatusOK, HomeResponse{
		Tracks:  tracks,
		Artists: artists,
		Albums:  albums,
	})

}
