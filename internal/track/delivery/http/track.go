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

func (h *Handler) GetTrackByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("ERROR: delivery.GetTrackByID: failed to parse id: %v", err)
		response.BadRequestJSON(w)
		return
	}

	track, err := h.service.GetTrackByID(r.Context(), id)
	if err != nil {
		h.handleError(w, err, "delivery.GetTrackByID")
		return
	}

	response.JSON(w, http.StatusOK, track)
}

func (h *Handler) GetAllTracks(w http.ResponseWriter, r *http.Request) {
	limit, offset := parsePagination(r)
	tracks, err := h.service.GetAllTracks(r.Context(), limit, offset)
	if err != nil {
		h.handleError(w, err, "delivery.GetAllTracks")
		return
	}
	response.JSON(w, http.StatusOK, tracks)
}

func (h *Handler) GetTracksByArtist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	artistID, err := uuid.Parse(vars["artistId"])
	if err != nil {
		log.Printf("ERROR: delivery.GetTracksByArtist: failed to parse artistId: %v", err)
		response.BadRequestJSON(w)
		return
	}

	limit, offset := parsePagination(r)
	tracks, err := h.service.GetTracksByArtistID(r.Context(), artistID, limit, offset)
	if err != nil {
		h.handleError(w, err, "delivery.GetTracksByArtist")
		return
	}

	response.JSON(w, http.StatusOK, tracks)
}

func (h *Handler) GetTracksByAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	albumID, err := uuid.Parse(vars["albumId"])
	if err != nil {
		log.Printf("ERROR: delivery.GetTracksByAlbum: failed to parse albumId: %v", err)
		response.BadRequestJSON(w)
		return
	}

	limit, offset := parsePagination(r)
	tracks, err := h.service.GetTracksByAlbumID(r.Context(), albumID, limit, offset)
	if err != nil {
		h.handleError(w, err, "delivery.GetTracksByAlbum")
		return
	}

	response.JSON(w, http.StatusOK, tracks)
}

func (h *Handler) GetTracksByGenre(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	genreID, err := uuid.Parse(vars["genreId"])
	if err != nil {
		log.Printf("ERROR: delivery.GetTracksByGenre: failed to parse genreId: %v", err)
		response.BadRequestJSON(w)
		return
	}

	limit, offset := parsePagination(r)
	tracks, err := h.service.GetTracksByGenreID(r.Context(), genreID, limit, offset)
	if err != nil {
		h.handleError(w, err, "delivery.GetTracksByGenre")
		return
	}

	response.JSON(w, http.StatusOK, tracks)
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
