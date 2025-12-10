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

const (
	wsSendBufferSize = 256
)

func (h *Handler) GetTrackComments(w http.ResponseWriter, r *http.Request) {
	const op = "handler.GetTrackComments"
	log := middleware.LoggerFromContext(r.Context())

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		log.Errorf("[%s]: id is missing in URL vars", op)
		response.BadRequestJSON(w)
		return
	}

	trackID, err := uuid.Parse(idStr)
	if err != nil {
		log.Warnf("[%s]: Failed to parse track ID from URL: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	limit, offset := parsePagination(r)

	comments, err := h.service.GetCommentsByTrackID(r.Context(), trackID, limit, offset)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			response.JSON(w, http.StatusOK, []interface{}{})
			return
		}
		log.Errorf("[%s]: service error: %v", op, err)
		response.InternalErrorJSON(w)
		return
	}

	response.JSON(w, http.StatusOK, comments)
}

func (h *Handler) ServeWS(w http.ResponseWriter, r *http.Request) {
	const op = "handler.ServeWS"
	log := middleware.LoggerFromContext(r.Context())

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		log.Errorf("[%s]: id is missing in URL vars", op)
		response.BadRequestJSON(w)
		return
	}

	trackID, err := uuid.Parse(idStr)
	if err != nil {
		log.Warnf("[%s]: Failed to parse track ID from URL: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	rawUserID, ok := middleware.GetUserID(r.Context())
	if !ok {
		log.Warnf("[%s]: unauthorized ws attempt", op)
		response.UnauthorizedJSON(w)
		return
	}
	userID, err := uuid.Parse(rawUserID)
	if err != nil {
		log.Errorf("[%s]: invalid userID in context: %v", op, err)
		response.InternalErrorJSON(w)
		return
	}

	conn, err := h.wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("[%s]: failed to upgrade connection: %v", op, err)
		return
	}

	client := &Client{
		hub:     h.hub,
		conn:    conn,
		send:    make(chan []byte, wsSendBufferSize),
		service: h.service,
		logger:  log,
		trackID: trackID,
		userID:  userID,
	}

	h.hub.register <- client

	go client.writePump()
	client.readPump()
}
