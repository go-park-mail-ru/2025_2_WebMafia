package http

import (
	"errors"
	"log"
	"net/http"
	"spotify/internal/artist/service"
	"spotify/pkg/response"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

const (
	defaultLimit     = 100
	defaultOffset    = 0
	maxLimit         = 1000
	queryParamLimit  = "limit"
	queryParamOffset = "offset"
)

func (h *Handler) GetArtistByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		log.Println("ERROR: delivery.GetArtistByID: id is missing in URL vars")
		response.BadRequestJSON(w)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("ERROR: delivery.GetArtistByID: failed to parse id '%s': %v", idStr, err)
		response.BadRequestJSON(w)
		return
	}

	artist, err := h.service.GetArtistByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Printf("INFO: delivery.GetArtistByID: resource not found: %v", err)
		} else {
			log.Printf("ERROR: delivery.GetArtistByID: internal server error: %v", err)
		}
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, artist)
}

func (h *Handler) GetAllArtists(w http.ResponseWriter, r *http.Request) {
	limit, offset := parsePagination(r)

	artists, err := h.service.GetAllArtists(r.Context(), limit, offset)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Printf("INFO: delivery.GetAllArtists: resource not found: %v", err)
		} else {
			log.Printf("ERROR: delivery.GetAllArtists: internal server error: %v", err)
		}
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, artists)
}

func parsePagination(r *http.Request) (uint64, uint64) {
	query := r.URL.Query()
	limitStr := query.Get(queryParamLimit)
	offsetStr := query.Get(queryParamOffset)

	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil || limit == 0 {
		limit = defaultLimit
	}
	if limit > maxLimit {
		limit = maxLimit
	}

	offset, err := strconv.ParseUint(offsetStr, 10, 64)
	if err != nil {
		offset = defaultOffset
	}
	return limit, offset
}
