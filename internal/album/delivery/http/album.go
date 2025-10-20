package http

import (
	"errors"
	"net/http"
	"spotify/internal/album/service"
	"spotify/internal/middleware"
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

func (h *Handler) GetAlbumByID(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With("op", "handler.GetArtistByID")

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		log.Errorw("id is missing in URL vars")
		response.BadRequestJSON(w)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Warnw("Failed to parse album ID from URL", "error", err)
		response.BadRequestJSON(w)
		return
	}

	artist, err := h.service.GetAlbumByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Infow("Resource not found", "error", err)
		} else {
			log.Errorw("Service error", "error", err)
		}
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, artist)
}

func (h *Handler) GetAllAlbums(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With("op", "handler.GetAllAlbums")
	limit, offset := parsePagination(r)

	artists, err := h.service.GetAllAlbums(r.Context(), limit, offset)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Infow("Resource not found", "error", err)
		} else {
			log.Errorw("Service error", "error", err)
		}
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, artists)
}

func (h *Handler) GetAlbumsByArtistID(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With("op", "handler.GetAlbumsByArtistID")

	vars := mux.Vars(r)
	artistIDStr, ok := vars["artistId"]
	if !ok {
		log.Errorw("artistId is missing in URL vars")
		response.BadRequestJSON(w)
		return
	}

	artistID, err := uuid.Parse(artistIDStr)
	if err != nil {
		log.Warnw("Failed to parse artistId from URL", "error", err)
		response.BadRequestJSON(w)
		return
	}

	limit, offset := parsePagination(r)

	albums, err := h.service.GetAlbumsByArtistID(r.Context(), artistID, limit, offset)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Infow("Resource not found", "error", err)
		} else {
			log.Errorw("Service error", "error", err)
		}
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, albums)
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
