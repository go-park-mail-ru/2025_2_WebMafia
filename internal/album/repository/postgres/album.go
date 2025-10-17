package postgres

import (
	"context"
	"spotify/internal/model"

	"github.com/google/uuid"
)

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*model.Album, error) {
	query := `
		SELECT album_id, title, avatar_url, artist_id, description, release_date, created_at, updated_at
		FROM album
		WHERE album_id = $1`

	var album model.Album
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&album.ID,
		&album.Title,
		&album.AvatarURL,
		&album.ArtistID,
		&album.Description,
		&album.ReleaseDate,
		&album.CreatedAt,
		&album.UpdatedAt,
	)

	if err != nil {
		return nil, mapErrors(err, "repository.GetByID")
	}

	return &album, nil
}

func (r *Repository) GetAll(ctx context.Context, limit, offset uint64) ([]model.Album, error) {
	query := `
		SELECT album_id, title, avatar_url, artist_id, description, release_date, created_at, updated_at
		FROM album
		ORDER BY release_date DESC
		LIMIT $1 OFFSET $2`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, mapErrors(err, "repository.GetAll: query failed")
	}
	defer rows.Close()

	albums := make([]model.Album, 0, limit)
	for rows.Next() {
		var album model.Album
		if err := rows.Scan(
			&album.ID,
			&album.Title,
			&album.AvatarURL,
			&album.ArtistID,
			&album.Description,
			&album.ReleaseDate,
			&album.CreatedAt,
			&album.UpdatedAt,
		); err != nil {
			return nil, mapErrors(err, "repository.GetAll: scan failed")
		}
		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		return nil, mapErrors(err, "repository.GetAll: rows iteration failed")
	}

	return albums, nil
}

func (r *Repository) GetByIDs(ctx context.Context, ids []uuid.UUID) ([]model.Album, error) {
	if len(ids) == 0 {
		return []model.Album{}, nil
	}

	query := `
		SELECT album_id, title, avatar_url, artist_id, description, release_date, created_at, updated_at
		FROM album
		WHERE album_id = ANY($1)`

	rows, err := r.db.QueryContext(ctx, query, ids)
	if err != nil {
		return nil, mapErrors(err, "repository.GetByIDs: query failed")
	}
	defer rows.Close()

	albums := make([]model.Album, 0, len(ids))
	for rows.Next() {
		var album model.Album
		if err := rows.Scan(
			&album.ID,
			&album.Title,
			&album.AvatarURL,
			&album.ArtistID,
			&album.Description,
			&album.ReleaseDate,
			&album.CreatedAt,
			&album.UpdatedAt,
		); err != nil {
			return nil, mapErrors(err, "repository.GetByIDs: scan failed")
		}
		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		return nil, mapErrors(err, "repository.GetByIDs: rows iteration failed")
	}

	return albums, nil
}

func (r *Repository) GetByArtistID(ctx context.Context, artistID uuid.UUID, limit, offset uint64) ([]model.Album, error) {
	query := `
		SELECT album_id, title, avatar_url, artist_id, description, release_date, created_at, updated_at
		FROM album
		WHERE artist_id = $1
		ORDER BY release_date DESC
		LIMIT $2 OFFSET $3`

	rows, err := r.db.QueryContext(ctx, query, artistID, limit, offset)
	if err != nil {
		return nil, mapErrors(err, "repository.GetByArtistID: query failed")
	}
	defer rows.Close()

	albums := make([]model.Album, 0, limit)
	for rows.Next() {
		var album model.Album
		if err := rows.Scan(
			&album.ID,
			&album.Title,
			&album.AvatarURL,
			&album.ArtistID,
			&album.Description,
			&album.ReleaseDate,
			&album.CreatedAt,
			&album.UpdatedAt,
		); err != nil {
			return nil, mapErrors(err, "repository.GetByArtistID: scan failed")
		}
		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		return nil, mapErrors(err, "repository.GetByArtistID: rows iteration failed")
	}

	return albums, nil
}
