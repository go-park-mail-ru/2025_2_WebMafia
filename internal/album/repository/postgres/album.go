package postgres

import (
	"context"
	"fmt"
	"spotify/internal/model"

	"github.com/google/uuid"
)

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*model.Album, error) {
	const op = "repository.GetAll"
	query := `
		SELECT album_id, title, type, avatar_url, artist_id, description, release_date, created_at, updated_at
		FROM album
		WHERE album_id = $1`

	var album model.Album
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&album.ID,
		&album.Title,
		&album.Type,
		&album.AvatarURL,
		&album.ArtistID,
		&album.Description,
		&album.ReleaseDate,
		&album.CreatedAt,
		&album.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", op, mapErrors(err))
	}

	return &album, nil
}

func (r *Repository) GetAll(ctx context.Context, limit, offset uint64) ([]model.Album, error) {
	const op = "repository.GetAll"
	query := `
		SELECT album_id, title, type, avatar_url, artist_id, description, release_date, created_at, updated_at
		FROM album
		ORDER BY release_date DESC
		LIMIT $1 OFFSET $2`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[%s]: query failed: %w", op, mapErrors(err))
	}
	defer rows.Close()

	albums := make([]model.Album, 0, limit)
	for rows.Next() {
		var album model.Album
		if err := rows.Scan(
			&album.ID,
			&album.Title,
			&album.Type,
			&album.AvatarURL,
			&album.ArtistID,
			&album.Description,
			&album.ReleaseDate,
			&album.CreatedAt,
			&album.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("[%s]: scan failed: %w", op, mapErrors(err))
		}
		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("[%s]: rows iteration failed: %w", op, mapErrors(err))
	}

	return albums, nil
}

func (r *Repository) GetByIDs(ctx context.Context, ids []uuid.UUID) ([]model.Album, error) {
	const op = "repository.GetByIDs"
	if len(ids) == 0 {
		return nil, nil
	}

	query := `
		SELECT album_id, title, type, avatar_url, artist_id, description, release_date, created_at, updated_at
		FROM album
		WHERE album_id = ANY($1)`

	rows, err := r.db.QueryContext(ctx, query, ids)
	if err != nil {
		return nil, fmt.Errorf("[%s]: query failed: %w", op, mapErrors(err))
	}
	defer rows.Close()

	albums := make([]model.Album, 0, len(ids))
	for rows.Next() {
		var album model.Album
		if err := rows.Scan(
			&album.ID,
			&album.Title,
			&album.Type,
			&album.AvatarURL,
			&album.ArtistID,
			&album.Description,
			&album.ReleaseDate,
			&album.CreatedAt,
			&album.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("[%s]: scan failed: %w", op, mapErrors(err))
		}
		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("[%s]: rows iteration failed: %w", op, mapErrors(err))
	}

	return albums, nil
}

func (r *Repository) GetByArtistID(ctx context.Context, artistID uuid.UUID, limit, offset uint64) ([]model.Album, error) {
	const op = "repository.GetByArtistID"
	query := `
		SELECT album_id, title, type, avatar_url, artist_id, description, release_date, created_at, updated_at
		FROM album
		WHERE artist_id = $1
		ORDER BY release_date DESC
		LIMIT $2 OFFSET $3`

	rows, err := r.db.QueryContext(ctx, query, artistID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[%s]: query failed: %w", op, mapErrors(err))
	}
	defer rows.Close()

	albums := make([]model.Album, 0, limit)
	for rows.Next() {
		var album model.Album
		if err := rows.Scan(
			&album.ID,
			&album.Title,
			&album.Type,
			&album.AvatarURL,
			&album.ArtistID,
			&album.Description,
			&album.ReleaseDate,
			&album.CreatedAt,
			&album.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("[%s]: scan failed: %w", op, mapErrors(err))
		}
		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("[%s]: rows iteration failed: %w", op, mapErrors(err))
	}

	return albums, nil
}
