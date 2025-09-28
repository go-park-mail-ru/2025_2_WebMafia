package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"spotify/internal/model"
	"spotify/pkg/response"
)

type ArtistsResponse struct {
	Artists []model.Artist `json:"artists"`
}

type ArtistResponse struct {
	Artist model.Artist `json:"artist"`
}

func (h *Handlers) GetAllArtistsHandler(w http.ResponseWriter, r *http.Request) {
	artists, _ := h.store.GetAllArtists()
	for i := range artists {
		enrichArtistURLs(h.cfg, r, &artists[i])
	}
	response.JSON(w, http.StatusOK, ArtistsResponse{Artists: artists})
}

func (h *Handlers) GetArtistByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	artistID := vars["id"]

	artists, _ := h.store.GetAllArtists()
	for _, a := range artists {
		if fmt.Sprint(a.ArtistID) == artistID {
			enrichArtistURLs(h.cfg, r, &a)
			response.JSON(w, http.StatusOK, ArtistResponse{Artist: a})
			return
		}
	}
	response.JSON(w, http.StatusNotFound, response.ErrorResponse{Error: "artist not found"})
}
