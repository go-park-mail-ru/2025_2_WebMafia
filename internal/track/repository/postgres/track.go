package postgres

import (
	"context"
	"database/sql"
	"spotify/internal/model"

	"github.com/google/uuid"
)

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*model.Track, error) {
	query := `
		SELECT track_id, title, duration_ms, file_url, description, created_at, updated_at
		FROM track
		WHERE track_id = $1`
	var track model.Track
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&track.ID,
		&track.Title,
		&track.DurationMs,
		&track.FileURL,
		&track.Description,
		&track.CreatedAt,
		&track.UpdatedAt,
	)
	if err != nil {
		return nil, mapErrors(err, "repository.GetByID")
	}
	return &track, nil
}

func (r *Repository) GetAll(ctx context.Context, limit, offset uint64) ([]model.Track, error) {
	query := `
		SELECT track_id, title, duration_ms, file_url, description, created_at, updated_at
		FROM track
		ORDER BY created_at
		LIMIT $1 OFFSET $2`
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, mapErrors(err, "repository.GetAll: query failed")
	}
	defer rows.Close()

	tracks, err := selectTracks(rows)
	if err != nil {
		return nil, mapErrors(err, "repository.GetAll: scan failed")
	}
	if err := rows.Err(); err != nil {
		return nil, mapErrors(err, "repository.GetAll: rows iteration failed")
	}
	return tracks, nil
}

func (r *Repository) GetByArtistID(ctx context.Context, artistID uuid.UUID, limit, offset uint64) ([]model.Track, error) {
	query := `
		SELECT t.track_id, t.title, t.duration_ms, t.file_url, t.description, t.created_at, t.updated_at
		FROM track t
		JOIN track_artist ta ON t.track_id = ta.track_id
		WHERE ta.artist_id = $1
		ORDER BY t.created_at
		LIMIT $2 OFFSET $3`
	rows, err := r.db.QueryContext(ctx, query, artistID, limit, offset)
	if err != nil {
		return nil, mapErrors(err, "repository.GetByArtistID: query failed")
	}
	defer rows.Close()

	tracks, err := selectTracks(rows)
	if err != nil {
		return nil, mapErrors(err, "repository.GetByArtistID: scan failed")
	}
	if err := rows.Err(); err != nil {
		return nil, mapErrors(err, "repository.GetByArtistID: rows iteration failed")
	}
	return tracks, nil
}

func (r *Repository) GetByAlbumID(ctx context.Context, albumID uuid.UUID, limit, offset uint64) ([]model.Track, error) {
	query := `
		SELECT t.track_id, t.title, t.duration_ms, t.file_url, t.description, t.created_at, t.updated_at
		FROM track t
		JOIN track_album ta ON t.track_id = ta.track_id
		WHERE ta.album_id = $1
		ORDER BY t.created_at
		LIMIT $2 OFFSET $3`
	rows, err := r.db.QueryContext(ctx, query, albumID, limit, offset)
	if err != nil {
		return nil, mapErrors(err, "repository.GetByAlbumID: query failed")
	}
	defer rows.Close()

	tracks, err := selectTracks(rows)
	if err != nil {
		return nil, mapErrors(err, "repository.GetByAlbumID: scan failed")
	}
	if err := rows.Err(); err != nil {
		return nil, mapErrors(err, "repository.GetByAlbumID: rows iteration failed")
	}
	return tracks, nil
}

func (r *Repository) GetByGenreID(ctx context.Context, genreID uuid.UUID, limit, offset uint64) ([]model.Track, error) {
	query := `
		SELECT t.track_id, t.title, t.duration_ms, t.file_url, t.description, t.created_at, t.updated_at
		FROM track t
		JOIN track_genre tg ON t.track_id = tg.track_id
		WHERE tg.genre_id = $1
		ORDER BY t.created_at
		LIMIT $2 OFFSET $3`
	rows, err := r.db.QueryContext(ctx, query, genreID, limit, offset)
	if err != nil {
		return nil, mapErrors(err, "repository.GetByGenreID: query failed")
	}
	defer rows.Close()

	tracks, err := selectTracks(rows)
	if err != nil {
		return nil, mapErrors(err, "repository.GetByGenreID: scan failed")
	}
	if err := rows.Err(); err != nil {
		return nil, mapErrors(err, "repository.GetByGenreID: rows iteration failed")
	}
	return tracks, nil
}

func (r *Repository) GetAlbumIDsForTracks(ctx context.Context, trackIDs []uuid.UUID) (map[uuid.UUID]uuid.UUID, error) {
	if len(trackIDs) == 0 {
		return make(map[uuid.UUID]uuid.UUID), nil
	}
	query := `SELECT track_id, album_id FROM track_album WHERE track_id = ANY($1)`
	rows, err := r.db.QueryContext(ctx, query, trackIDs)
	if err != nil {
		return nil, mapErrors(err, "repository.GetAlbumIDsForTracks: query failed")
	}
	defer rows.Close()

	result := make(map[uuid.UUID]uuid.UUID)
	for rows.Next() {
		var trackID, albumID uuid.UUID
		if err := rows.Scan(&trackID, &albumID); err != nil {
			return nil, mapErrors(err, "repository.GetAlbumIDsForTracks: scan failed")
		}
		result[trackID] = albumID
	}
	if err := rows.Err(); err != nil {
		return nil, mapErrors(err, "repository.GetAlbumIDsForTracks: rows iteration failed")
	}
	return result, nil
}

func (r *Repository) GetArtistIDsForTracks(ctx context.Context, trackIDs []uuid.UUID) (map[uuid.UUID][]uuid.UUID, error) {
	if len(trackIDs) == 0 {
		return make(map[uuid.UUID][]uuid.UUID), nil
	}
	query := `SELECT track_id, artist_id FROM track_artist WHERE track_id = ANY($1)`
	rows, err := r.db.QueryContext(ctx, query, trackIDs)
	if err != nil {
		return nil, mapErrors(err, "repository.GetArtistIDsForTracks: query failed")
	}
	defer rows.Close()

	result := make(map[uuid.UUID][]uuid.UUID)
	for rows.Next() {
		var trackID, artistID uuid.UUID
		if err := rows.Scan(&trackID, &artistID); err != nil {
			return nil, mapErrors(err, "repository.GetArtistIDsForTracks: scan failed")
		}
		result[trackID] = append(result[trackID], artistID)
	}
	if err := rows.Err(); err != nil {
		return nil, mapErrors(err, "repository.GetArtistIDsForTracks: rows iteration failed")
	}
	return result, nil
}

func (r *Repository) GetGenresForTracks(ctx context.Context, trackIDs []uuid.UUID) (map[uuid.UUID][]model.Genre, error) {
	if len(trackIDs) == 0 {
		return make(map[uuid.UUID][]model.Genre), nil
	}
	query := `
		SELECT tg.track_id, g.genre_id, g.genre_name, g.description, g.created_at
		FROM genre g
		JOIN track_genre tg ON g.genre_id = tg.genre_id
		WHERE tg.track_id = ANY($1)`
	rows, err := r.db.QueryContext(ctx, query, trackIDs)
	if err != nil {
		return nil, mapErrors(err, "repository.GetGenresForTracks: query failed")
	}
	defer rows.Close()

	result := make(map[uuid.UUID][]model.Genre)
	for rows.Next() {
		var trackID uuid.UUID
		var genre model.Genre
		if err := rows.Scan(
			&trackID,
			&genre.ID,
			&genre.Name,
			&genre.Description,
			&genre.CreatedAt,
		); err != nil {
			return nil, mapErrors(err, "repository.GetGenresForTracks: scan failed")
		}
		result[trackID] = append(result[trackID], genre)
	}
	if err := rows.Err(); err != nil {
		return nil, mapErrors(err, "repository.GetGenresForTracks: rows iteration failed")
	}
	return result, nil
}

func selectTracks(rows *sql.Rows) ([]model.Track, error) {
	tracks := make([]model.Track, 0)
	for rows.Next() {
		var track model.Track
		if err := rows.Scan(
			&track.ID,
			&track.Title,
			&track.DurationMs,
			&track.FileURL,
			&track.Description,
			&track.CreatedAt,
			&track.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}
	return tracks, nil
}
