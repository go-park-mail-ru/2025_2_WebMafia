package http

import (
	"errors"
	"log"
	"net/http"
	"spotify/internal/track/service"
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
		if errors.Is(err, service.ErrNotFound) {
			log.Printf("INFO: delivery.GetTrackByID: resource not found: %v", err)
		} else {
			log.Printf("ERROR: delivery.GetTrackByID: internal server error: %v", err)
		}
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, track)
}

func (h *Handler) GetAllTracks(w http.ResponseWriter, r *http.Request) {
	limit, offset := parsePagination(r)
	tracks, err := h.service.GetAllTracks(r.Context(), limit, offset)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Printf("INFO: delivery.GetAllTracks: resource not found: %v", err)
		} else {
			log.Printf("ERROR: delivery.GetAllTracks: internal server error: %v", err)
		}
		h.handleError(w, err)
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
		if errors.Is(err, service.ErrNotFound) {
			log.Printf("INFO: delivery.GetTracksByArtist: resource not found: %v", err)
		} else {
			log.Printf("ERROR: delivery.GetTracksByArtist: internal server error: %v", err)
		}
		h.handleError(w, err)
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
		if errors.Is(err, service.ErrNotFound) {
			log.Printf("INFO: delivery.GetTracksByAlbum: resource not found: %v", err)
		} else {
			log.Printf("ERROR: delivery.GetTracksByAlbum: internal server error: %v", err)
		}
		h.handleError(w, err)
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
		if errors.Is(err, service.ErrNotFound) {
			log.Printf("INFO: delivery.GetTracksByGenre: resource not found: %v", err)
		} else {
			log.Printf("ERROR: delivery.GetTracksByGenre: internal server error: %v", err)
		}
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, tracks)
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
