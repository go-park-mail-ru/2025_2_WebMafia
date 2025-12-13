package http

import (
	"errors"
	"net/http"
	"spotify/internal/middleware"
	"spotify/microservices/catalog/service"
	"spotify/pkg/response"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (h *Handler) GetTrackByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.GetTrackByID"
	log := middleware.LoggerFromContext(r.Context())

	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		log.Errorf("[%s]: id is missing in URL vars", op)
		response.BadRequestJSON(w)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Warnf("[%s]: Failed to parse track ID from URL: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	track, err := h.service.GetTrackByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Infof("[%s]: Resource not found: %v", op, err)
		} else {
			log.Errorf("[%s]: Service error: %v", op, err)
		}
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, track)
}

func (h *Handler) GetAllTracks(w http.ResponseWriter, r *http.Request) {
	const op = "handler.GetAllTracks"
	log := middleware.LoggerFromContext(r.Context())

	limit, offset := parsePagination(r)
	tracks, err := h.service.GetAllTracks(r.Context(), limit, offset)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Infof("[%s]: Resource not found: %v", op, err)
		} else {
			log.Errorf("[%s]: Service error: %v", op, err)
		}
		h.handleError(w, err)
		return
	}
	response.JSON(w, http.StatusOK, tracks)
}

func (h *Handler) GetTracksByArtist(w http.ResponseWriter, r *http.Request) {
	const op = "handler.GetTracksByArtist"
	log := middleware.LoggerFromContext(r.Context())

	vars := mux.Vars(r)
	idStr, ok := vars["artistId"]
	if !ok {
		log.Errorf("[%s]: id is missing in URL vars", op)
		response.BadRequestJSON(w)
		return
	}

	artistID, err := uuid.Parse(idStr)
	if err != nil {
		log.Warnf("[%s]: Failed to parse track ID from URL: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	limit, offset := parsePagination(r)
	tracks, err := h.service.GetTracksByArtistID(r.Context(), artistID, limit, offset)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Infof("[%s]: Resource not found: %v", op, err)
		} else {
			log.Errorf("[%s]: Service error: %v", op, err)
		}
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, tracks)
}

func (h *Handler) GetTracksByAlbum(w http.ResponseWriter, r *http.Request) {
	const op = "handler.GetTracksByAlbum"
	log := middleware.LoggerFromContext(r.Context())

	vars := mux.Vars(r)
	idStr, ok := vars["albumId"]
	if !ok {
		log.Errorf("[%s]: id is missing in URL vars", op)
		response.BadRequestJSON(w)
		return
	}

	albumID, err := uuid.Parse(idStr)
	if err != nil {
		log.Warnf("[%s]: Failed to parse albumId from URL: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	limit, offset := parsePagination(r)
	tracks, err := h.service.GetTracksByAlbumID(r.Context(), albumID, limit, offset)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Infof("[%s]: Resource not found: %v", op, err)
		} else {
			log.Errorf("[%s]: Service error: %v", op, err)
		}
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, tracks)
}

func (h *Handler) GetTracksByGenre(w http.ResponseWriter, r *http.Request) {
	const op = "handler.GetTracksByGenre"
	log := middleware.LoggerFromContext(r.Context())

	vars := mux.Vars(r)
	idStr, ok := vars["genreId"]
	if !ok {
		log.Errorf("[%s]: id is missing in URL vars", op)
		response.BadRequestJSON(w)
		return
	}
	genreID, err := uuid.Parse(idStr)
	if err != nil {
		log.Warnf("[%s]: Failed to parse genreId from URL: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	limit, offset := parsePagination(r)
	tracks, err := h.service.GetTracksByGenreID(r.Context(), genreID, limit, offset)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Infof("[%s]: Resource not found: %v", op, err)
		} else {
			log.Errorf("[%s]: Service error: %v", op, err)
		}
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, tracks)
}

func (h *Handler) RegisterPlay(w http.ResponseWriter, r *http.Request) {
	const op = "handler.RegisterPlay"
	log := middleware.LoggerFromContext(r.Context())

	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		log.Errorf("[%s]: id is missing in URL vars", op)
		response.BadRequestJSON(w)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Warnf("[%s]: Failed to parse track ID from URL: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	err = h.service.RegisterPlay(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Infof("[%s]: Resource not found: %v", op, err)
		} else {
			log.Errorf("[%s]: Service error: %v", op, err)
		}
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusAccepted, nil)
}

func (h *Handler) SearchTracks(w http.ResponseWriter, r *http.Request) {
	const op = "handler.SearchTracks"
	log := middleware.LoggerFromContext(r.Context())

	query := r.URL.Query().Get(queryParamSearch)
	if query == "" {
		log.Warnf("[%s]: search query is empty", op)
		response.BadRequestJSON(w)
		return
	}

	limit, _ := parsePagination(r)
	if limit == defaultLimit {
		limit = defaultSearchLimit
	}

	results, err := h.service.SearchTracks(r.Context(), query, limit)
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, results)
}
