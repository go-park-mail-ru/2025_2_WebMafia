package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"spotify/internal/album/model"

	"github.com/google/uuid"
)

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*model.Album, *model.Artist, error) {
	query := `
		SELECT
			a.album_id, a.title, a.avatar_url, a.description, a.release_date, a.created_at, a.updated_at,
			ar.artist_id, ar.artist_name, ar.avatar_url, ar.description, ar.created_at, ar.updated_at
		FROM albums a
		JOIN artists ar ON a.artist_id = ar.artist_id
		WHERE a.album_id = $1`

	var album model.Album
	var artist model.Artist

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&album.ID,
		&album.Title,
		&album.AvatarURL,
		&album.Description,
		&album.ReleaseDate,
		&album.CreatedAt,
		&album.UpdatedAt,
		&artist.ID,
		&artist.Name,
		&artist.AvatarURL,
		&artist.Description,
		&artist.CreatedAt,
		&artist.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, ErrNotFound
		}
		return nil, nil, fmt.Errorf("repository.GetByID: queryrow or scan fail %w", err)
	}
	album.ArtistID = artist.ID

	return &album, &artist, nil
}

func (r *Repository) GetAll(ctx context.Context) ([]model.Album, []model.Artist, error) {
	query := `
		SELECT
			a.album_id, a.title, a.avatar_url, a.description, a.release_date, a.created_at, a.updated_at,
			ar.artist_id, ar.artist_name, ar.avatar_url, ar.description, ar.created_at, ar.updated_at
		FROM albums a
		JOIN artists ar ON a.artist_id = ar.artist_id
		ORDER BY a.release_date DESC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, nil, fmt.Errorf("repository.GetAll: query fail: %w", err)
	}

	defer rows.Close()

	albums := make([]model.Album, 0)
	artists := make([]model.Artist, 0)

	for rows.Next() {
		var album model.Album
		var artist model.Artist
		if err := rows.Scan(
			&album.ID,
			&album.Title,
			&album.AvatarURL,
			&album.Description,
			&album.ReleaseDate,
			&album.CreatedAt,
			&album.UpdatedAt,
			&artist.ID,
			&artist.Name,
			&artist.AvatarURL,
			&artist.Description,
			&artist.CreatedAt,
			&artist.UpdatedAt,
		); err != nil {
			return nil, nil, fmt.Errorf("repository.GetAll: row scan fail: %w", err)
		}
		album.ArtistID = artist.ID
		albums = append(albums, album)
		artists = append(artists, artist)
	}

	if err := rows.Err(); err != nil {
		return nil, nil, fmt.Errorf("repository.GetAll: rows iteration fail: %w", err)
	}

	return albums, artists, nil
}
