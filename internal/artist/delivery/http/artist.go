package http

import (
	"errors"
	"log"
	"net/http"
	"spotify/internal/artist/dto"
	"spotify/internal/artist/service"
	"spotify/pkg/response"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type ArtistsResponse struct {
	Artists []dto.Artist `json:"artists"`
}

type ArtistResponse struct {
	Artist *dto.Artist `json:"artist"`
}

func (h *Handler) GetArtistByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		log.Println("ERROR: delivery.GetArtistByID: id is missing")
		response.BadRequestJSON(w)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("ERROR: delivery.GetArtistByID: failed to parse id: %v", err)
		response.BadRequestJSON(w)
		return
	}

	artist, err := h.service.GetArtistByID(r.Context(), id)
	if err != nil {
		log.Printf("ERROR: delivery.GetArtistByID: service error: %v", err)
		if errors.Is(err, service.ErrNotFound) {
			response.NotFoundJSON(w)
			return
		}
		response.InternalErrorJSON(w)
		return
	}

	response.JSON(w, http.StatusOK, ArtistResponse{Artist: artist})
}

func (h *Handler) GetAllArtists(w http.ResponseWriter, r *http.Request) {
	artists, err := h.service.GetAllArtists(r.Context())
	if err != nil {
		log.Printf("ERROR: delivery.GetAllArtists: service error: %v", err)
		response.InternalErrorJSON(w)
		return
	}

	response.JSON(w, http.StatusOK, ArtistsResponse{Artists: artists})
}
