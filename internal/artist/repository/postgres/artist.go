package postgres

import (
	"context"
	"fmt"
	"spotify/internal/model"

	"github.com/google/uuid"
)

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*model.Artist, error) {
	const op = "repository.GetByID"
	query := `
		SELECT artist_id, artist_name, avatar_url, description, created_at, updated_at
		FROM artist
		WHERE artist_id = $1`

	var artist model.Artist

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&artist.ID,
		&artist.Name,
		&artist.AvatarURL,
		&artist.Description,
		&artist.CreatedAt,
		&artist.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", op, mapErrors(err))
	}

	return &artist, nil
}

func (r *Repository) GetByIDs(ctx context.Context, ids []uuid.UUID) ([]model.Artist, error) {
	const op = "repository.GetByIDs"
	if len(ids) == 0 {
		return nil, nil
	}

	query := `
		SELECT artist_id, artist_name, avatar_url, description, created_at, updated_at
		FROM artist
		WHERE artist_id = ANY($1)`

	rows, err := r.db.QueryContext(ctx, query, ids)
	if err != nil {
		return nil, fmt.Errorf("[%s]: query failed: %w", op, mapErrors(err))
	}
	defer rows.Close()

	artists := make([]model.Artist, 0, len(ids))
	for rows.Next() {
		var artist model.Artist
		if err := rows.Scan(
			&artist.ID,
			&artist.Name,
			&artist.AvatarURL,
			&artist.Description,
			&artist.CreatedAt,
			&artist.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("[%s]: scan failed: %w", op, mapErrors(err))
		}
		artists = append(artists, artist)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("[%s]: rows iteration failed: %w", op, mapErrors(err))
	}

	return artists, nil
}

func (r *Repository) GetAll(ctx context.Context, limit, offset uint64) ([]model.Artist, error) {
	const op = "repository.GetAll"
	query := `
		SELECT artist_id, artist_name, avatar_url, description, created_at, updated_at
		FROM artist
		ORDER BY artist_name
		LIMIT $1 OFFSET $2`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[%s]: query failed: %w", op, mapErrors(err))
	}
	defer rows.Close()

	artists := make([]model.Artist, 0, limit)
	for rows.Next() {
		var artist model.Artist
		if err := rows.Scan(
			&artist.ID,
			&artist.Name,
			&artist.AvatarURL,
			&artist.Description,
			&artist.CreatedAt,
			&artist.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("[%s]: scan failed: %w", op, mapErrors(err))
		}
		artists = append(artists, artist)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("[%s]: rows iteration failed: %w", op, mapErrors(err))
	}

	return artists, nil
}
