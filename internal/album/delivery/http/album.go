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

func (h *Handler) GetAlbumByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		log.Println("ERROR: delivery.GetAlbumByID: id is missing in URL vars")
		response.BadRequestJSON(w)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("ERROR: delivery.GetAlbumByID: failed to parse id '%s': %v", idStr, err)
		response.BadRequestJSON(w)
		return
	}

	artist, err := h.service.GetAlbumByID(r.Context(), id)
	if err != nil {
		h.handleError(w, err, "delivery.GetAlbumByID: service error")
		return
	}

	response.JSON(w, http.StatusOK, artist)
}

func (h *Handler) GetAllAlbums(w http.ResponseWriter, r *http.Request) {
	limit, offset := parsePagination(r)

	artists, err := h.service.GetAllAlbums(r.Context(), limit, offset)
	if err != nil {
		h.handleError(w, err, "delivery.GetAllAlbums: service error")
		return
	}

	response.JSON(w, http.StatusOK, artists)
}

func (h *Handler) GetAlbumsByArtistID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	artistIDStr, ok := vars["artistId"]
	if !ok {
		log.Println("ERROR: delivery.GetAlbumsByArtistID: artistId is missing in URL vars")
		response.BadRequestJSON(w)
		return
	}

	artistID, err := uuid.Parse(artistIDStr)
	if err != nil {
		log.Printf("ERROR: delivery.GetAlbumsByArtistID: failed to parse artistId '%s': %v", artistIDStr, err)
		response.BadRequestJSON(w)
		return
	}

	limit, offset := parsePagination(r)

	albums, err := h.service.GetAlbumsByArtistID(r.Context(), artistID, limit, offset)
	if err != nil {
		h.handleError(w, err, "delivery.GetAlbumsByArtistID")
		return
	}

	response.JSON(w, http.StatusOK, albums)
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
