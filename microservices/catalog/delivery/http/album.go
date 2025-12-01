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

func (h *Handler) GetAlbumByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.GetAlbumByID"
	log := middleware.LoggerFromContext(r.Context())

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		log.Errorf("[%s]: id is missing in URL vars", op)
		response.BadRequestJSON(w)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Warnf("[%s]: failed to parse album ID from URL: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	album, err := h.service.GetAlbumByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Infof("[%s]: resource not found: %v", op, err)
		} else {
			log.Errorf("[%s]: service error: %v", op, err)
		}
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, album)
}

func (h *Handler) GetAllAlbums(w http.ResponseWriter, r *http.Request) {
	const op = "handler.GetAllAlbums"
	log := middleware.LoggerFromContext(r.Context())
	limit, offset := parsePagination(r)

	albums, err := h.service.GetAllAlbums(r.Context(), limit, offset)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Infof("[%s]: resource not found: %v", op, err)
		} else {
			log.Errorf("[%s]: service error: %v", op, err)
		}
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, albums)
}

func (h *Handler) GetAlbumsByArtistID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.GetAlbumsByArtistID"
	log := middleware.LoggerFromContext(r.Context())

	vars := mux.Vars(r)
	artistIDStr, ok := vars["artistId"]
	if !ok {
		log.Errorf("[%s]: artistId is missing in URL vars", op)
		response.BadRequestJSON(w)
		return
	}

	artistID, err := uuid.Parse(artistIDStr)
	if err != nil {
		log.Warnf("[%s]: failed to parse artistId from URL: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	limit, offset := parsePagination(r)

	albums, err := h.service.GetAlbumsByArtistID(r.Context(), artistID, limit, offset)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Infof("[%s]: resource not found: %v", op, err)
		} else {
			log.Errorf("[%s]: service error: %v", op, err)
		}
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, albums)
}

func (h *Handler) SearchAlbums(w http.ResponseWriter, r *http.Request) {
	const op = "handler.SearchAlbums"
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

	results, err := h.service.SearchAlbums(r.Context(), query, limit)
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, results)
}
