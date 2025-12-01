package postgres

import (
	"context"
	"fmt"

	"spotify/internal/model"
	"spotify/microservices/catalog/dto"

	"github.com/google/uuid"
)

func (r *Repository) GetAlbumByID(ctx context.Context, id uuid.UUID) (*model.Album, error) {
	const op = "repository.GetAlbumByID"
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

func (r *Repository) GetAllAlbums(ctx context.Context, limit, offset uint64) ([]model.Album, error) {
	const op = "repository.GetAllAlbums"
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

func (r *Repository) GetAlbumsByIDs(ctx context.Context, ids []uuid.UUID) ([]model.Album, error) {
	const op = "repository.GetAlbumsByIDs"
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

func (r *Repository) GetAlbumsByArtistID(ctx context.Context, artistID uuid.UUID, limit, offset uint64) ([]model.Album, error) {
	const op = "repository.GetAlbumsByArtistID"
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

func (r *Repository) SearchAlbums(ctx context.Context, query string, limit uint64) ([]dto.AlbumSearchResult, error) {
	const op = "repository.SearchAlbums"

	sqlQuery := `
		SELECT album_id, title, type, avatar_url, artist_id, description, release_date, created_at, updated_at,
			ts_rank_cd(fts_vector, plainto_tsquery('simple', $1)) as rank
		FROM album
		WHERE fts_vector @@ plainto_tsquery('simple', $1)
		ORDER BY rank DESC
		LIMIT $2;`

	rows, err := r.db.QueryContext(ctx, sqlQuery, query, limit)
	if err != nil {
		return nil, fmt.Errorf("[%s]: query failed: %w", op, mapErrors(err))
	}
	defer rows.Close()

	results := make([]dto.AlbumSearchResult, 0, limit)
	for rows.Next() {
		var searchResult dto.AlbumSearchResult
		if err := rows.Scan(
			&searchResult.Album.ID,
			&searchResult.Album.Title,
			&searchResult.Album.Type,
			&searchResult.Album.AvatarURL,
			&searchResult.Album.ArtistID,
			&searchResult.Album.Description,
			&searchResult.Album.ReleaseDate,
			&searchResult.Album.CreatedAt,
			&searchResult.Album.UpdatedAt,
			&searchResult.Rank,
		); err != nil {
			return nil, fmt.Errorf("[%s]: scan failed: %w", op, mapErrors(err))
		}
		results = append(results, searchResult)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("[%s]: rows iteration failed: %w", op, mapErrors(err))
	}

	return results, nil
}
