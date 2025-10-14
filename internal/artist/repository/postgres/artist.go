package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"spotify/internal/artist/model"

	"github.com/google/uuid"
)

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*model.Artist, error) {
	query := `
		SELECT artist_id, artist_name, avatar_url, description, created_at, updated_at
		FROM artists
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("repository.GetByID: %w", err)
	}

	return &artist, nil
}

func (r *Repository) GetAll(ctx context.Context) ([]model.Artist, error) {
	query := `
		SELECT artist_id, artist_name, avatar_url, description, created_at, updated_at
		FROM artists
		ORDER BY artist_name ASC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository.GetAll: %w", err)
	}
	defer rows.Close()

	artists := make([]model.Artist, 0)
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
			return nil, fmt.Errorf("repository.GetAll: scan: %w", err)
		}
		artists = append(artists, artist)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("repository.GetAll: rows error: %w", err)
	}

	return artists, nil
}
