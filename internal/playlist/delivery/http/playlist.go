package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"spotify/internal/middleware"
	"spotify/internal/playlist/dto"
	"spotify/internal/playlist/service"
	"spotify/pkg/response"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	defaultLimit     = 100
	defaultOffset    = 0
	maxLimit         = 1000
	queryParamLimit  = "limit"
	queryParamOffset = "offset"
)

func (h *Handler) CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	const op = "handler.CreatePlaylist"
	log := middleware.LoggerFromContext(r.Context())
	defer r.Body.Close()

	userID, ok := middleware.GetUserID(r.Context())
	if !ok || userID == "" {
		response.UnauthorizedJSON(w)
		return
	}

	var req dto.CreatePlaylistRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Errorf("[%s]: invalid request body: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	req.UserID = userID

	playlist, err := h.service.CreatePlaylist(r.Context(), req)
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusCreated, playlist)
}

func (h *Handler) GetPlaylistByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.GetPlaylistByID"
	log := middleware.LoggerFromContext(r.Context())

	id := mux.Vars(r)["id"]
	if id == "" {
		log.Errorf("[%s]: missing playlist id", op)
		response.BadRequestJSON(w)
		return
	}

	req := dto.GetPlaylistRequest{ID: id}

	playlist, err := h.service.GetPlaylist(r.Context(), req)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Infof("[%s]: playlist not found: %v", op, err)
		} else {
			log.Errorf("[%s]: service error: %v", op, err)
		}
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, playlist)
}

func (h *Handler) GetAllPlaylistsByUserID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.GetAllPlaylistsByUserID"
	log := middleware.LoggerFromContext(r.Context())

	userID := mux.Vars(r)["userId"]
	if userID == "" {
		log.Errorf("[%s]: missing userId", op)
		response.BadRequestJSON(w)
		return
	}

	limit, offset := parsePagination(r)

	req := dto.GetPlaylistsByUserRequest{
		UserID: userID,
		Limit:  limit,
		Offset: offset,
	}

	playlists, err := h.service.GetPlaylistsByUser(r.Context(), req)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Infof("[%s]: playlists not found for user %s", op, userID)
		} else {
			log.Errorf("[%s]: service error: %v", op, err)
		}
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, playlists)
}

func (h *Handler) UpdatePlaylist(w http.ResponseWriter, r *http.Request) {
	const op = "handler.UpdatePlaylist"
	log := middleware.LoggerFromContext(r.Context())
	defer r.Body.Close()

	id := mux.Vars(r)["id"]
	if id == "" {
		log.Errorf("[%s]: missing playlist id", op)
		response.BadRequestJSON(w)
		return
	}

	var req dto.UpdatePlaylistRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Errorf("[%s]: invalid request body: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	req.ID = id

	updated, err := h.service.UpdatePlaylist(r.Context(), req)
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, updated)
}

func (h *Handler) DeletePlaylist(w http.ResponseWriter, r *http.Request) {
	const op = "handler.DeletePlaylist"
	log := middleware.LoggerFromContext(r.Context())

	id := mux.Vars(r)["id"]
	if id == "" {
		log.Errorf("[%s]: missing playlist id", op)
		response.BadRequestJSON(w)
		return
	}

	req := dto.DeletePlaylistRequest{ID: id}

	if err := h.service.DeletePlaylist(r.Context(), req); err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func (h *Handler) AddTrackToFavorite(w http.ResponseWriter, r *http.Request) {
	const op = "handler.AddTrackToFavorite"

	log := middleware.LoggerFromContext(r.Context())

	userID, ok := middleware.GetUserID(r.Context())
	if !ok {
		response.UnauthorizedJSON(w)
		return
	}

	var req dto.AddTrackToFavoriteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Errorf("[%s]: invalid request body: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	req.UserID = userID
	if err := h.service.AddTrackToFavorite(r.Context(), req); err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		h.handleError(w, err)
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) GetFavoritePlaylist(w http.ResponseWriter, r *http.Request) {
	const op = "handler.GetFavoritePlaylist"

	log := middleware.LoggerFromContext(r.Context())

	userID, ok := middleware.GetUserID(r.Context())
	if !ok {
		response.UnauthorizedJSON(w)
		return
	}

	req := dto.GetFavoritePlaylistRequest{
		UserID: userID,
	}

	playlist, err := h.service.GetFavoritePlaylist(r.Context(), req)
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		h.handleError(w, err)
		return
	}
	response.JSON(w, http.StatusOK, playlist)
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
