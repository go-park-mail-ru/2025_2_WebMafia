package http

import (
	"context"
	"spotify/internal/ticket/dto"

	"github.com/google/uuid"
)

type IService interface {
	CreateTicket(ctx context.Context, req dto.CreateTicketRequest, userID uuid.UUID) (*dto.TicketResponse, error)
	GetUserTickets(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]dto.TicketResponse, error)
	UpdateTicket(ctx context.Context, req dto.UpdateTicketRequest, ticketID, userID uuid.UUID) (*dto.TicketResponse, error)
	GetAllTickets(ctx context.Context, limit, offset uint64) ([]dto.TicketResponse, error)
}

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{
		service: service,
	}
}
