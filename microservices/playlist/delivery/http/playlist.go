package http

import (
	"fmt"
	"net/http"
	"net/url"
	"spotify/internal/middleware"
	"spotify/microservices/playlist/dto"
	"spotify/pkg/response"
	"strconv"

	"github.com/google/uuid"
	"github.com/mailru/easyjson"

	"github.com/gorilla/mux"
)

const (
	defaultLimit          = 100
	defaultOffset         = 0
	maxLimit              = 1000
	queryParamLimit       = "limit"
	queryParamOffset      = "offset"
	paramPlaylistID       = "id"
	paramUserID           = "userId"
	maxPlaylistAvatarSize = 5 << 20
)

func (h *Handler) CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	const op = "handler.CreatePlaylist"
	log := middleware.LoggerFromContext(r.Context())
	defer r.Body.Close()

	rawUserID, ok := middleware.GetUserID(r.Context())
	if !ok || rawUserID == "" {
		log.Errorf("[%s]: missing userId", op)
		response.InternalErrorJSON(w)
		return
	}

	userID, err := uuid.Parse(rawUserID)
	if err != nil {
		log.Errorf("[%s]: invalid userId from context: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	var req dto.CreatePlaylistRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &req); err != nil {
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

	rawID := mux.Vars(r)[paramPlaylistID]
	if rawID == "" {
		log.Errorf("[%s]: missing playlist id", op)
		response.BadRequestJSON(w)
		return
	}

	id, err := uuid.Parse(rawID)
	if err != nil {
		log.Errorf("[%s]: invalid playlist id: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	playlist, err := h.service.GetPlaylistWithTracks(r.Context(), id)
	if err != nil {
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, playlist)
}

func (h *Handler) UpdatePlaylist(w http.ResponseWriter, r *http.Request) {
	const op = "handler.UpdatePlaylist"
	log := middleware.LoggerFromContext(r.Context())
	defer r.Body.Close()

	rawID := mux.Vars(r)[paramPlaylistID]
	if rawID == "" {
		log.Errorf("[%s]: missing playlist id", op)
		response.BadRequestJSON(w)
		return
	}

	id, err := uuid.Parse(rawID)
	if err != nil {
		log.Errorf("[%s]: invalid playlist id: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	var req dto.UpdatePlaylistRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &req); err != nil {
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

	rawID := mux.Vars(r)[paramPlaylistID]
	if rawID == "" {
		log.Errorf("[%s]: missing playlist id", op)
		response.BadRequestJSON(w)
		return
	}

	id, err := uuid.Parse(rawID)
	if err != nil {
		log.Errorf("[%s]: invalid playlist id: %v", op, err)
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

	rawUserID, ok := middleware.GetUserID(r.Context())
	if !ok {
		log.Errorf("[%s]: missing userId", op)
		response.InternalErrorJSON(w)
		return
	}

	userID, err := uuid.Parse(rawUserID)
	if err != nil {
		log.Errorf("[%s]: invalid userId: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	var req dto.AddTrackToFavoriteRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &req); err != nil {
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

	rawUserID, ok := middleware.GetUserID(r.Context())
	if !ok {
		log.Errorf("[%s] missing userId", op)
		response.UnauthorizedJSON(w)
		return
	}

	userID, err := uuid.Parse(rawUserID)
	if err != nil {
		log.Errorf("[%s] invalid userId: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	req := dto.GetFavoritePlaylistRequest{UserID: userID}

	pl, err := h.service.GetFavoritePlaylist(r.Context(), req)
	if err != nil {
		log.Errorf("[%s] service error: %v", op, err)
		h.handleError(w, err)
		return
	}

	id, err := uuid.Parse(pl.ID)
	if err != nil {
		log.Errorf("[%s] failed to parse playlist ID '%s': %v", op, pl.ID, err)
		response.InternalErrorJSON(w)
		return
	}

	full, err := h.service.GetPlaylistWithTracks(r.Context(), id)
	if err != nil {
		log.Errorf("[%s] failed to load tracks: %v", op, err)
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, full)
}

func (h *Handler) validatePlaylistAvatar(contentType string, size int64) error {
	if size == 0 {
		return fmt.Errorf("empty file")
	}
	if size > maxPlaylistAvatarSize {
		return fmt.Errorf("file too large (max 5MB)")
	}
	for _, allowed := range h.allowedAvatarTypes {
		if contentType == allowed {
			return nil
		}
	}
	return fmt.Errorf("unsupported content type: %s", contentType)
}

func (h *Handler) UploadPlaylistAvatar(w http.ResponseWriter, r *http.Request) {
	const op = "handler.UploadPlaylistAvatar"

	userID, ok := middleware.GetUserID(r.Context())
	if !ok || userID == "" {
		response.UnauthorizedJSON(w)
		return
	}

	log := middleware.LoggerFromContext(r.Context())

	rawID := mux.Vars(r)["id"]
	playlistID, err := uuid.Parse(rawID)
	if err != nil {
		log.Errorf("[%s]: invalid playlist id: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	file, header, err := r.FormFile("avatar")
	if err != nil {
		log.Errorf("[%s]: failed to get file: %v", op, err)
		response.BadRequestJSON(w)
		return
	}
	defer file.Close()

	if err := h.validatePlaylistAvatar(header.Header.Get("Content-Type"), header.Size); err != nil {
		log.Errorf("[%s] validation error: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	res, err := h.service.UploadPlaylistAvatar(r.Context(), dto.UploadPlaylistAvatarRequest{
		PlaylistID:  playlistID,
		File:        file,
		Size:        header.Size,
		ContentType: header.Header.Get("Content-Type"),
	})
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, res)
}

func (h *Handler) DeletePlaylistAvatar(w http.ResponseWriter, r *http.Request) {
	const op = "handler.DeletePlaylistAvatar"
	log := middleware.LoggerFromContext(r.Context())

	userID, ok := middleware.GetUserID(r.Context())
	if !ok || userID == "" {
		response.UnauthorizedJSON(w)
		return
	}

	rawID := mux.Vars(r)["id"]
	playlistID, err := uuid.Parse(rawID)
	if err != nil {
		log.Errorf("[%s] invalid playlist id: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	req := dto.DeletePlaylistAvatarRequest{
		PlaylistID: playlistID,
	}
	if err := h.service.DeletePlaylistAvatar(r.Context(), req); err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func (h *Handler) GetMyPlaylists(w http.ResponseWriter, r *http.Request) {
	const op = "handler.GetMyPlaylists"
	log := middleware.LoggerFromContext(r.Context())

	rawUserID, ok := middleware.GetUserID(r.Context())
	if !ok || rawUserID == "" {
		log.Errorf("[%s]: missing userId", op)
		response.UnauthorizedJSON(w)
		return
	}

	userID, err := uuid.Parse(rawUserID)
	if err != nil {
		log.Errorf("[%s]: invalid userId: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	limit, offset := parsePagination(r.URL.Query())

	playlists, err := h.service.GetPlaylistsByUser(r.Context(), dto.GetPlaylistsByUserRequest{
		UserID: userID,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		h.handleError(w, err)
		return
	}

	full := make([]dto.Playlist, 0, len(playlists))
	for _, p := range playlists {
		id, _ := uuid.Parse(p.ID)
		pl, err := h.service.GetPlaylistWithTracks(r.Context(), id)
		if err == nil {
			full = append(full, *pl)
		}
	}

	response.JSON(w, http.StatusOK, full)
}

func (h *Handler) AddTrackToPlaylist(w http.ResponseWriter, r *http.Request) {
	const op = "handler.AddTrackToPlaylist"
	log := middleware.LoggerFromContext(r.Context())

	raw := mux.Vars(r)["id"]
	playlistID, err := uuid.Parse(raw)
	if err != nil {
		log.Errorf("[%s]: invalid playlist id: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	var req dto.AddTrackToPlaylistRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &req); err != nil {
		log.Errorf("[%s]: invalid body: %v", op, err)
		response.BadRequestJSON(w)
		return
	}
	req.PlaylistID = playlistID

	if err := h.service.AddTrackToPlaylist(r.Context(), req); err != nil {
		log.Errorf("[%s]: service: %v", op, err)
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) RemoveTrackFromPlaylist(w http.ResponseWriter, r *http.Request) {
	const op = "handler.RemoveTrackFromPlaylist"
	log := middleware.LoggerFromContext(r.Context())

	raw := mux.Vars(r)["id"]
	playlistID, err := uuid.Parse(raw)
	if err != nil {
		log.Errorf("[%s]: invalid playlist id: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	var req dto.RemoveTrackFromPlaylistRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &req); err != nil {
		log.Errorf("[%s]: invalid body: %v", op, err)
		response.BadRequestJSON(w)
		return
	}
	req.PlaylistID = playlistID

	if err := h.service.RemoveTrackFromPlaylist(r.Context(), req); err != nil {
		log.Errorf("[%s]: service: %v", op, err)
		h.handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{"status": "removed"})
}

func parsePagination(query url.Values) (uint64, uint64) {
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
