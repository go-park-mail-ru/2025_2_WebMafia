package http

import (
	"errors"
	"net/http"
	"spotify/internal/middleware"
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
	log := middleware.LoggerFromContext(r.Context()).With("op", "handler.GetTrackByID")

	vars := mux.Vars(r)

	idStr, ok := vars["id"]
	if !ok {
		log.Errorw("id is missing in URL vars")
		response.BadRequestJSON(w)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Warnw("Failed to parse track ID from URL", "error", err)
		response.BadRequestJSON(w)
		return
	}

	track, err := h.service.GetTrackByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Infow("Resource not found", "error", err)
		} else {
			log.Errorw("Service error", "error", err)
		}
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, track)
}

func (h *Handler) GetAllTracks(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With("op", "handler.GetAllTracks")

	limit, offset := parsePagination(r)
	tracks, err := h.service.GetAllTracks(r.Context(), limit, offset)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Infow("Resource not found", "error", err)
		} else {
			log.Errorw("Service error", "error", err)
		}
		h.handleError(w, err)
		return
	}
	response.JSON(w, http.StatusOK, tracks)
}

func (h *Handler) GetTracksByArtist(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With("op", "handler.GetTracksByArtist")

	vars := mux.Vars(r)
	idStr, ok := vars["artistId"]
	if !ok {
		log.Errorw("id is missing in URL vars")
		response.BadRequestJSON(w)
		return
	}

	artistID, err := uuid.Parse(idStr)
	if err != nil {
		log.Warnw("Failed to parse track ID from URL", "error", err)
		response.BadRequestJSON(w)
		return
	}

	limit, offset := parsePagination(r)
	tracks, err := h.service.GetTracksByArtistID(r.Context(), artistID, limit, offset)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Infow("Resource not found", "error", err)
		} else {
			log.Errorw("Service error", "error", err)
		}
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, tracks)
}

func (h *Handler) GetTracksByAlbum(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With("op", "handler.GetTracksByAlbum")

	vars := mux.Vars(r)
	idStr, ok := vars["albumId"]
	if !ok {
		log.Errorw("id is missing in URL vars")
		response.BadRequestJSON(w)
		return
	}

	albumID, err := uuid.Parse(idStr)
	if err != nil {
		log.Warnw("Failed to parse albumId from URL", "error", err)
		response.BadRequestJSON(w)
		return
	}

	limit, offset := parsePagination(r)
	tracks, err := h.service.GetTracksByAlbumID(r.Context(), albumID, limit, offset)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Infow("Resource not found", "error", err)
		} else {
			log.Errorw("Service error", "error", err)
		}
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, tracks)
}

func (h *Handler) GetTracksByGenre(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With("op", "handler.GetTracksByGenre")

	vars := mux.Vars(r)
	idStr, ok := vars["genreId"]
	if !ok {
		log.Errorw("id is missing in URL vars")
		response.BadRequestJSON(w)
		return
	}
	genreID, err := uuid.Parse(idStr)
	if err != nil {
		log.Warnw("Failed to parse genreId from URL", "error", err)
		response.BadRequestJSON(w)
		return
	}

	limit, offset := parsePagination(r)
	tracks, err := h.service.GetTracksByGenreID(r.Context(), genreID, limit, offset)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Infow("Resource not found", "error", err)
		} else {
			log.Errorw("Service error", "error", err)
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
