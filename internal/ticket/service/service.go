package service

import (
	"context"
	"spotify/internal/model"
	"spotify/internal/ticket/dto"

	"github.com/google/uuid"
)

type IRepository interface {
	Create(ctx context.Context, ticket model.Ticket) error
	GetByID(ctx context.Context, ticketID uuid.UUID) (*model.Ticket, error)
	GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]model.Ticket, error)
	Update(ctx context.Context, ticket model.Ticket) error
	GetAll(ctx context.Context, limit, offset uint64) ([]model.Ticket, error)
	UpdateStatus(ctx context.Context, ticketID uuid.UUID, status string) error
	GetStatistics(ctx context.Context) (*dto.TicketStatistics, error)
}

type Service struct {
	repo IRepository
}

func New(repo IRepository) *Service {
	return &Service{
		repo: repo,
	}
}
