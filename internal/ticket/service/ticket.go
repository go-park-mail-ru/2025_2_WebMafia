package service

import (
	"context"
	"database/sql"
	"fmt"
	"spotify/internal/model"
	"spotify/internal/ticket/dto"
	"time"

	"github.com/google/uuid"
)

func (s *Service) CreateTicket(ctx context.Context, req dto.CreateTicketRequest, userID uuid.UUID) (*dto.TicketResponse, error) {
	const op = "service.CreateTicket"
	ticket := model.Ticket{
		ID:          uuid.New(),
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		Category:    req.Category,
		Status:      "Открыто",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.Create(ctx, ticket); err != nil {
		return nil, fmt.Errorf("[%s]: %w", op, mapError(err))
	}

	return toTicketResponse(&ticket), nil
}

func (s *Service) GetUserTickets(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]dto.TicketResponse, error) {
	const op = "service.GetUserTickets"
	tickets, err := s.repo.GetByUserID(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", op, mapError(err))
	}

	responses := make([]dto.TicketResponse, len(tickets))
	for i, ticket := range tickets {
		responses[i] = *toTicketResponse(&ticket)
	}
	return responses, nil
}

func (s *Service) GetAllTickets(ctx context.Context, limit, offset uint64) ([]dto.TicketResponse, error) {
	const op = "service.GetAllTickets"
	tickets, err := s.repo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", op, mapError(err))
	}

	responses := make([]dto.TicketResponse, len(tickets))
	for i, ticket := range tickets {
		responses[i] = *toTicketResponse(&ticket)
	}
	return responses, nil
}

func (s *Service) UpdateTicket(ctx context.Context, req dto.UpdateTicketRequest, ticketID, userID uuid.UUID) (*dto.TicketResponse, error) {
	const op = "service.UpdateTicket"
	ticket, err := s.repo.GetByID(ctx, ticketID)
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", op, mapError(err))
	}

	if ticket.UserID != userID {
		return nil, ErrForbidden
	}

	updated := false

	if req.Title != nil || req.Description != nil {
		if ticket.Status != "Открыто" {
			return nil, ErrInvalidStateForAction
		}
		if req.Title != nil {
			ticket.Title = *req.Title
		}
		if req.Description != nil {
			ticket.Description = *req.Description
		}
		updated = true
	}

	if req.Status != nil {
		if *req.Status != "Закрыто" || ticket.Status != "Открыто" {
			return nil, ErrInvalidStateForAction
		}
		ticket.Status = *req.Status
		updated = true
	}

	if req.Rating != nil {
		if ticket.Status != "Закрыто" {
			return nil, ErrInvalidStateForAction
		}
		if *req.Rating < 1 || *req.Rating > 5 {
			return nil, ErrInvalidRating
		}
		ticket.Rating = sql.NullInt32{Int32: int32(*req.Rating), Valid: true}
		updated = true
	}

	if updated {
		ticket.UpdatedAt = time.Now()
		if err := s.repo.Update(ctx, *ticket); err != nil {
			return nil, fmt.Errorf("[%s]: %w", op, mapError(err))
		}
	}

	return toTicketResponse(ticket), nil
}

func (s *Service) UpdateTicketStatusByAdmin(ctx context.Context, ticketID uuid.UUID, status string) (*dto.TicketResponse, error) {
	const op = "service.UpdateTicketStatusByAdmin"

	if status != "В работе" && status != "Закрыто" {
		return nil, ErrInvalidStateForAction
	}

	_, err := s.repo.GetByID(ctx, ticketID)
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", op, mapError(err))
	}

	if err := s.repo.UpdateStatus(ctx, ticketID, status); err != nil {
		return nil, fmt.Errorf("[%s]: %w", op, mapError(err))
	}

	updatedTicket, err := s.repo.GetByID(ctx, ticketID)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get updated ticket: %w", op, mapError(err))
	}

	return toTicketResponse(updatedTicket), nil
}

func (s *Service) GetStatistics(ctx context.Context) (*dto.TicketStatistics, error) {
	const op = "service.GetStatistics"

	stats, err := s.repo.GetStatistics(ctx)
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", op, mapError(err))
	}

	return stats, nil
}

func toTicketResponse(ticket *model.Ticket) *dto.TicketResponse {
	resp := &dto.TicketResponse{
		ID:          ticket.ID.String(),
		UserID:      ticket.UserID.String(),
		Title:       ticket.Title,
		Description: ticket.Description,
		Category:    ticket.Category,
		Status:      ticket.Status,
		CreatedAt:   ticket.CreatedAt,
		UpdatedAt:   ticket.UpdatedAt,
	}

	if ticket.Rating.Valid {
		rating := ticket.Rating.Int32
		resp.Rating = &rating
	}

	return resp
}
