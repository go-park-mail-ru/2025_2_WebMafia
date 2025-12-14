package http

import (
	"context"
	"errors"
	"net/http"
	"spotify/internal/middleware"
	"spotify/microservices/catalog/dto"
	"spotify/microservices/catalog/service"
	"spotify/pkg/response"
	"spotify/pkg/ws"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
)

const (
	maxTextLength      = 200
	saveCommentTimeout = 5 * time.Second
)

func (h *Handler) GetTrackComments(w http.ResponseWriter, r *http.Request) {
	const op = "handler.GetTrackComments"
	log := middleware.LoggerFromContext(r.Context())

	idStr, ok := mux.Vars(r)["id"]
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

	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		log.Errorf("[%s]: id is missing in URL vars", op)
		response.BadRequestJSON(w)
		return
	}

	trackUUID, err := uuid.Parse(idStr)
	if err != nil {
		log.Warnf("[%s]: Failed to parse track ID from URL: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	rawUserID, ok := middleware.GetUserID(r.Context())
	if !ok {
		log.Errorf("[%s]: userID missing in context", op)
		response.InternalErrorJSON(w)
		return
	}

	responseHeader := http.Header{}
	if protocol := r.Header.Get("Sec-WebSocket-Protocol"); protocol != "" {
		responseHeader.Add("Sec-WebSocket-Protocol", protocol)
	}
	conn, err := h.wsUpgrader.Upgrade(w, r, responseHeader)
	if err != nil {
		log.Errorf("[%s]: failed to upgrade connection: %v", op, err)
		return
	}

	client := ws.NewClient(
		trackUUID.String(),
		rawUserID,
		h.hub,
		conn,
		log,
		h,
		h.wsConfig,
	)

	h.hub.Register(client)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Errorf("WritePump panic: %v", r)
			}
		}()
		client.WritePump()
	}()
	client.ReadPump()
}

func (h *Handler) HandleMessage(ctx context.Context, client *ws.Client, message []byte) {
	const op = "handler.HandleMessage"
	log := middleware.LoggerFromContext(ctx)

	var req dto.PostCommentRequest
	if err := easyjson.Unmarshal(message, &req); err != nil {
		log.Warnf("[%s]: invalid json format from client %s: %v", op, client.ID(), err)
		h.sendErrorJSON(client, response.ErrBadRequest)
		return
	}

	if len(req.Text) == 0 || len(req.Text) > maxTextLength {
		log.Warnf("[%s]: invalid text length from client %s: %d", op, client.ID(), len(req.Text))
		h.sendErrorJSON(client, response.ErrBadRequest)
		return
	}

	req.TrackID = client.Topic()
	userID, _ := uuid.Parse(client.ID())

	saveCtx, cancel := context.WithTimeout(ctx, saveCommentTimeout)
	defer cancel()

	createdComment, err := h.service.PostComment(saveCtx, userID, req)
	if err != nil {
		log.Errorf("[%s]: failed to save comment: %v", op, err)
		h.sendErrorJSON(client, response.ErrInternalServer)
		return
	}

	responseBytes, err := easyjson.Marshal(createdComment)
	if err != nil {
		log.Errorf("[%s]: failed to marshal response: %v", op, err)
		return
	}

	h.hub.BroadcastTo(client.Topic(), responseBytes)
}

func (h *Handler) sendErrorJSON(client *ws.Client, errResp response.ErrorResponse) {
	data, err := easyjson.Marshal(errResp)
	if err != nil {
		return
	}
	client.SendMessage(data)
}
