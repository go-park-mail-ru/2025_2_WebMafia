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

func (h *Handler) GetArtistByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.GetArtistByID"
	log := middleware.LoggerFromContext(r.Context())

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		log.Errorf("[%s]: id is missing in URL vars")
		response.BadRequestJSON(w)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Warnf("[%s]: failed to parse artist ID from URL: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	artist, err := h.service.GetArtistByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Infof("[%s]: resource not found: %v", op, err)
		} else {
			log.Errorf("[%s]: service error: %v", op, err)
		}
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, artist)
}

func (h *Handler) GetAllArtists(w http.ResponseWriter, r *http.Request) {
	const op = "handler.GetAllArtists"
	log := middleware.LoggerFromContext(r.Context())

	limit, offset := parsePagination(r)

	artists, err := h.service.GetAllArtists(r.Context(), limit, offset)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Infof("[%s]: resource not found: %v", op, err)
		} else {
			log.Errorf("[%s]: service error: %v", op, err)
		}
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, artists)
}

func (h *Handler) SearchArtists(w http.ResponseWriter, r *http.Request) {
	const op = "handler.SearchArtists"
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

	results, err := h.service.SearchArtists(r.Context(), query, limit)
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, results)
}
