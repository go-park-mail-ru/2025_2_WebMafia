package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"spotify/internal/model"
	"spotify/internal/ticket/dto"
	"time"

	"github.com/google/uuid"
)

func (m *Repository) Create(ctx context.Context, ticket model.Ticket) error {
	const op = "repository.Create"

	query := `INSERT INTO ticket (ticket_id, user_id, status, category, title, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := m.Conn.ExecContext(ctx,
		query,
		ticket.ID, ticket.UserID, ticket.Status, ticket.Category,
		ticket.Title, ticket.Description, ticket.CreatedAt, ticket.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (m *Repository) GetByID(ctx context.Context, id uuid.UUID) (*model.Ticket, error) {
	const op = "repository.GetByID"

	query := `
		SELECT ticket_id, user_id, title, description, category, status, rating, created_at, updated_at
		FROM ticket
		WHERE ticket_id = $1`

	var ticket model.Ticket
	err := m.Conn.QueryRowContext(ctx, query, id).Scan(
		&ticket.ID,
		&ticket.UserID,
		&ticket.Title,
		&ticket.Description,
		&ticket.Category,
		&ticket.Status,
		&ticket.Rating,
		&ticket.CreatedAt,
		&ticket.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", op, mapErrors(err))
	}
	return &ticket, nil
}

func (m *Repository) GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]model.Ticket, error) {
	const op = "repository.GetByUserID"

	query := `
		SELECT ticket_id, user_id, title, description, category, status, rating, created_at, updated_at
		FROM ticket
		WHERE user_id = $1
		ORDER BY updated_at DESC
		LIMIT $2 OFFSET $3`

	rows, err := m.Conn.QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[%s]: query failed: %w", op, mapErrors(err))
	}
	defer rows.Close()

	return selectTickets(rows)
}

func (m *Repository) Update(ctx context.Context, ticket model.Ticket) error {
	const op = "repository.Update"

	query := `
		UPDATE ticket
		SET title = $1,
			description = $2,
			status = $3,
			rating = $4,
			updated_at = $5
		WHERE ticket_id = $6`

	_, err := m.Conn.ExecContext(ctx, query,
		ticket.Title,
		ticket.Description,
		ticket.Status,
		ticket.Rating,
		ticket.UpdatedAt,
		ticket.ID,
	)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (m *Repository) GetAll(ctx context.Context, limit, offset uint64) ([]model.Ticket, error) {
	const op = "repository.GetAll"

	query := `
		SELECT ticket_id, user_id, title, description, category, status, rating, created_at, updated_at
		FROM ticket
		ORDER BY updated_at DESC
		LIMIT $1 OFFSET $2`
	rows, err := m.Conn.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[%s]: query failed: %w", op, mapErrors(err))
	}
	defer rows.Close()

	return selectTickets(rows)
}

func (m *Repository) UpdateStatus(ctx context.Context, ticketID uuid.UUID, status string) error {
	const op = "repository.UpdateStatus"

	query := `
		UPDATE ticket
		SET status = $1,
			updated_at = $2
		WHERE ticket_id = $3`

	result, err := m.Conn.ExecContext(ctx, query, status, time.Now(), ticketID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: could not get rows affected: %w", op, err)
	}
	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (m *Repository) GetStatistics(ctx context.Context) (*dto.TicketStatistics, error) {
	const op = "repository.GetStatistics"

	queryStats := `
		SELECT
			COUNT(*),
			COUNT(*) FILTER (WHERE status = 'Открыто'),
			COUNT(*) FILTER (WHERE status = 'В работе'),
			COUNT(*) FILTER (WHERE status = 'Закрыто')
		FROM ticket`

	stats := &dto.TicketStatistics{}
	err := m.Conn.QueryRowContext(ctx, queryStats).Scan(
		&stats.TotalTickets,
		&stats.OpenTickets,
		&stats.InProgress,
		&stats.ClosedTickets,
	)
	if err != nil {
		return nil, fmt.Errorf("[%s]: query failed for main stats: %w", op, mapErrors(err))
	}

	queryCategory := `
		SELECT category, COUNT(*)
		FROM ticket
		GROUP BY category`

	rows, err := m.Conn.QueryContext(ctx, queryCategory)
	if err != nil {
		return nil, fmt.Errorf("[%s]: query failed for category stats: %w", op, mapErrors(err))
	}
	defer rows.Close()

	stats.ByCategory = make(map[string]int)
	for rows.Next() {
		var category string
		var count int
		if err := rows.Scan(&category, &count); err != nil {
			return nil, fmt.Errorf("[%s]: scan failed for category stats: %w", op, err)
		}
		stats.ByCategory[category] = count
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("[%s]: rows iteration failed for category stats: %w", op, err)
	}

	return stats, nil
}

func selectTickets(rows *sql.Rows) ([]model.Ticket, error) {
	tickets := make([]model.Ticket, 0)
	for rows.Next() {
		var ticket model.Ticket
		if err := rows.Scan(
			&ticket.ID,
			&ticket.UserID,
			&ticket.Title,
			&ticket.Description,
			&ticket.Category,
			&ticket.Status,
			&ticket.Rating,
			&ticket.CreatedAt,
			&ticket.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tickets, nil
}
