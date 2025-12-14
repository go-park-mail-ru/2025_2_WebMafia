package postgres

import (
	"context"
	"fmt"
	"spotify/internal/model"

	"github.com/google/uuid"
)

func (r *Repository) CreateComment(ctx context.Context, comment model.Comment) error {
	const op = "repository.CreateComment"

	query := `
		INSERT INTO comment (comment_id, track_id, user_id, text, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(ctx, query,
		comment.ID,
		comment.TrackID,
		comment.UserID,
		comment.Text,
		comment.CreatedAt,
		comment.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("[%s]: %w", op, mapErrors(err))
	}

	return nil
}

func (r *Repository) GetCommentsByTrackID(ctx context.Context, trackID uuid.UUID, limit, offset uint64) ([]model.Comment, error) {
	const op = "repository.GetCommentsByTrackID"

	query := `
		SELECT comment_id, track_id, user_id, text, created_at, updated_at
		FROM comment
		WHERE track_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, trackID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[%s]: query failed: %w", op, mapErrors(err))
	}
	defer rows.Close()

	comments := make([]model.Comment, 0, limit)
	for rows.Next() {
		var c model.Comment
		if err := rows.Scan(
			&c.ID,
			&c.TrackID,
			&c.UserID,
			&c.Text,
			&c.CreatedAt,
			&c.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("[%s]: scan failed: %w", op, mapErrors(err))
		}
		comments = append(comments, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("[%s]: rows iteration failed: %w", op, mapErrors(err))
	}

	if len(comments) == 0 {
		return nil, fmt.Errorf("[%s]: no comments found: %w", op, ErrNotFound)
	}

	return comments, nil
}
