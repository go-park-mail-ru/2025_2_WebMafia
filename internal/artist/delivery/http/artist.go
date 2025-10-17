package http

import (
	"log"
	"net/http"
	"spotify/pkg/response"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

const (
	DefaultLimit  = 100
	DefaultOffset = 0
	MaxLimit      = 1000
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
		h.handleError(w, err, "delivery.GetArtistByID: service error")
		return
	}

	response.JSON(w, http.StatusOK, artist)
}

func (h *Handler) GetAllArtists(w http.ResponseWriter, r *http.Request) {
	limit, offset := parsePagination(r)

	artists, err := h.service.GetAllArtists(r.Context(), limit, offset)
	if err != nil {
		h.handleError(w, err, "delivery.GetAllArtists: service error")
		return
	}

	response.JSON(w, http.StatusOK, artists)
}

func parsePagination(r *http.Request) (uint64, uint64) {
	query := r.URL.Query()
	limitStr := query.Get("limit")
	offsetStr := query.Get("offset")

	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil || limit == 0 {
		limit = DefaultLimit
	}
	if limit > MaxLimit {
		limit = MaxLimit
	}

	offset, err := strconv.ParseUint(offsetStr, 10, 64)
	if err != nil {
		offset = DefaultOffset
	}
	return limit, offset
}
