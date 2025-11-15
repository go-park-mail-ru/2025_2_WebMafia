package http

import (
	"encoding/json"
	"net/http"
	"spotify/internal/middleware"
	"spotify/internal/ticket/dto"
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

func (h *Handler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	const op = "handler.CreateTicket"
	log := middleware.LoggerFromContext(r.Context())

	userIDStr, ok := middleware.GetUserID(r.Context())
	if !ok {
		log.Warnf("[%s]: unauthorized", op)
		response.UnauthorizedJSON(w)
		return
	}
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		log.Warnf("[%s]: failed to parse userID: %v", op, err)
		response.UnauthorizedJSON(w)
		return
	}

	var req dto.CreateTicketRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Warnf("[%s]: failed to decode request body: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	ticket, err := h.service.CreateTicket(r.Context(), req, userID)
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		handleError(w, err)
		return
	}

	log.Infof("[%s]: ticket created successfully: %s", op, ticket.ID)
	response.JSON(w, http.StatusCreated, ticket)
}

func (h *Handler) GetUserTickets(w http.ResponseWriter, r *http.Request) {
	const op = "handler.GetUserTickets"
	log := middleware.LoggerFromContext(r.Context())

	userIDStr, ok := middleware.GetUserID(r.Context())
	if !ok {
		log.Warnf("[%s]: unauthorized", op)
		response.UnauthorizedJSON(w)
		return
	}
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		log.Warnf("[%s]: failed to parse userID: %v", op, err)
		response.UnauthorizedJSON(w)
		return
	}

	limit, offset := parsePagination(r)

	tickets, err := h.service.GetUserTickets(r.Context(), userID, limit, offset)
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		handleError(w, err)
		return
	}
	response.JSON(w, http.StatusOK, tickets)
}

func (h *Handler) UpdateTicket(w http.ResponseWriter, r *http.Request) {
	const op = "handler.UpdateTicket"
	log := middleware.LoggerFromContext(r.Context())

	userIDStr, ok := middleware.GetUserID(r.Context())
	if !ok {
		log.Warnf("[%s]: unauthorized", op)
		response.UnauthorizedJSON(w)
		return
	}
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		log.Warnf("[%s]: failed to parse userID: %v", op, err)
		response.UnauthorizedJSON(w)
		return
	}

	vars := mux.Vars(r)
	ticketID, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Warnf("[%s]: failed to parse ticket ID from URL: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	var req dto.UpdateTicketRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Warnf("[%s]: failed to decode request body: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	ticket, err := h.service.UpdateTicket(r.Context(), req, ticketID, userID)
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		handleError(w, err)
		return
	}

	log.Infof("[%s]: ticket updated successfully: %s", op, ticket.ID)
	response.JSON(w, http.StatusOK, ticket)
}

func (h *Handler) GetAllTickets(w http.ResponseWriter, r *http.Request) {
	const op = "handler.GetAllTickets"
	log := middleware.LoggerFromContext(r.Context())

	limit, offset := parsePagination(r)

	tickets, err := h.service.GetAllTickets(r.Context(), limit, offset)
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, tickets)
}

func (h *Handler) GetStatistics(w http.ResponseWriter, r *http.Request) {
	const op = "handler.GetStatistics"
	log := middleware.LoggerFromContext(r.Context())

	stats, err := h.service.GetStatistics(r.Context())
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, stats)
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
