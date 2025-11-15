package service

import (
	"context"
	"spotify/internal/model"

	"github.com/google/uuid"
)

type IRepository interface {
	Create(ctx context.Context, ticket model.SupportTicket) error
	GetByID(ctx context.Context, ticketID uuid.UUID) (*model.SupportTicket, error)
	GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]model.SupportTicket, error)
	Update(ctx context.Context, ticket model.SupportTicket) error
	GetAll(ctx context.Context, limit, offset uint64) ([]model.SupportTicket, error)
}

type Service struct {
	repo IRepository
}

func New(repo IRepository) *Service {
	return &Service{
		repo: repo,
	}
}
