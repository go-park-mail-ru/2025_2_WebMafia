package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"spotify/internal/model"
)

func (m *Repository) Create(ctx context.Context, ticket model.Ticket) error {
	const op = "repository.Create"

	query := `INSERT INTO "ticket" (ticket_id, user_id, status, category, title, description)
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := m.Conn.ExecContext(ctx,
		query,
		ticket.ID, ticket.UserID, ticket.Status, ticket.Category,
		ticket.Description, ticket.Title,
	)

	if err != nil {
		return fmt.Errorf("%s: %w", op, mapErrors(err))
	}
	return nil

}

func (m *Repository) GetById(ctx context.Context, id uuid.UUID) (*model.Ticket, error) {
	const op = "repository.GetById"

	query := `
		SELECT ticket_id, user_id, status, category, title, description, created_at, updated_at
		FROM ticket
		WHERE ticket_id = $1`

	var ticket model.Ticket

	err := m.Conn.QueryRowContext(ctx, query, id).Scan(
		&ticket.ID,
		&ticket.UserID,
		&ticket.Status,
		&ticket.Category,
		&ticket.Title,
		&ticket.Description,
		&ticket.Rating,
		&ticket.CreatedAt,
		&ticket.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", op, mapErrors(err))
	}
	return &ticket, nil
}

func (m *Repository) GetByUserId(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]model.Ticket, error) {
	const op = "repository.GetByUserId"

	query := `
		SELECT ticket.ticket_id, ticket.user_id, ticket.status, ticket.category, ticket.title, ticket.description, ticket.rating, ticket.created_at, ticket.updated_at
		ORDER BY ticket.updated_at
		LIMIT $2 OFFSET $3`

	rows, err := m.Conn.QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[%s]: query failed: %w", op, mapErrors(err))
	}
	defer rows.Close()

	tickets, err := selectTicket(rows)
	if err != nil {
		return nil, fmt.Errorf("[%s]: scan failed: %w", op, mapErrors(err))
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("[%s]: rows iteration failed: %w", op, mapErrors(err))
	}
	return tickets, nil
}

func (m *Repository) Update(ctx context.Context, ticket model.Ticket) error {
	const op = "repository.Update"

	query := `
		UPDATE "ticket"
		SET title = $1,
			description = $2,
			rating = $3
		WHERE ticket_id = $4`

	_, err := m.Conn.ExecContext(ctx, query,
		ticket.Title,
		ticket.Description,
		ticket.Rating,
		ticket.ID,
	)

	if err != nil {
		return fmt.Errorf("%s: %w", op, mapErrors(err))
	}
	return nil
}

func (m *Repository) GetAll(ctx context.Context, limit, offset uint64) ([]model.Ticket, error) {
	const op = "repository.GetAll"

	query := `
		SELECT ticket_id, user_id, status, category, title, description, rating, created_at, updated_at
		FROM ticket
		ORDER BY updated_at
		LIMIT $1 OFFSET $2`
	rows, err := m.Conn.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[%s]: query failed: %w", op, err)
	}
	defer rows.Close()

	tickets, err := selectTicket(rows)
	if err != nil {
		return nil, fmt.Errorf("[%s]: scan failed: %w", op, mapErrors(err))
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("[%s]: rows iteration failed: %w", op, mapErrors(err))
	}
	return tickets, nil
}

func selectTicket(rows *sql.Rows) ([]model.Ticket, error) {
	tickets := make([]model.Ticket, 0)
	for rows.Next() {
		var ticket model.Ticket
		if err := rows.Scan(
			&ticket.ID,
			&ticket.UserID,
			&ticket.Status,
			&ticket.Category,
			&ticket.Title,
			&ticket.Description,
			&ticket.Rating,
			&ticket.CreatedAt,
			&ticket.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}
