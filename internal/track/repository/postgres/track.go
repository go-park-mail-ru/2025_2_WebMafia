package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"spotify/internal/model"

	"github.com/google/uuid"
)

const maxIDsInBatch = 100

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*model.Track, error) {
	const op = "repository.GetByID"
	query := `
		SELECT track_id, title, duration_s, file_url, description, created_at, updated_at
		FROM track
		WHERE track_id = $1`
	var track model.Track
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&track.ID,
		&track.Title,
		&track.DurationS,
		&track.FileURL,
		&track.PlayCount,
		&track.Description,
		&track.CreatedAt,
		&track.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", op, mapErrors(err))
	}
	return &track, nil
}

func (r *Repository) GetAll(ctx context.Context, limit, offset uint64) ([]model.Track, error) {
	const op = "repository.GetAll"
	query := `
		SELECT track_id, title, duration_s, file_url, play_count, description, created_at, updated_at
		FROM track
		ORDER BY created_at
		LIMIT $1 OFFSET $2`
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[%s]: query failed: %w", op, mapErrors(err))
	}
	defer rows.Close()

	tracks, err := selectTracks(rows)
	if err != nil {
		return nil, fmt.Errorf("[%s]: scan failed: %w", op, mapErrors(err))
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("[%s]: rows iteration failed: %w", op, mapErrors(err))
	}
	return tracks, nil
}

func (r *Repository) GetByArtistID(ctx context.Context, artistID uuid.UUID, limit, offset uint64) ([]model.Track, error) {
	const op = "repository.GetByArtistID"
	query := `
		SELECT t.track_id, t.title, t.duration_s, t.file_url, t.play_count, t.description, t.created_at, t.updated_at
		FROM track t
		JOIN track_artist ta ON t.track_id = ta.track_id
		WHERE ta.artist_id = $1
		ORDER BY t.created_at
		LIMIT $2 OFFSET $3`
	rows, err := r.db.QueryContext(ctx, query, artistID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[%s]: query failed: %w", op, mapErrors(err))
	}
	defer rows.Close()

	tracks, err := selectTracks(rows)
	if err != nil {
		return nil, fmt.Errorf("[%s]: scan failed: %w", op, mapErrors(err))
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("[%s]: rows iteration failed: %w", op, mapErrors(err))
	}
	return tracks, nil
}

func (r *Repository) GetByAlbumID(ctx context.Context, albumID uuid.UUID, limit, offset uint64) ([]model.Track, error) {
	const op = "repository.GetByAlbumID"
	query := `
		SELECT t.track_id, t.title, t.duration_s, t.file_url, t.play_count, t.description, t.created_at, t.updated_at
		FROM track t
		JOIN track_album ta ON t.track_id = ta.track_id
		WHERE ta.album_id = $1
		ORDER BY t.created_at
		LIMIT $2 OFFSET $3`
	rows, err := r.db.QueryContext(ctx, query, albumID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[%s]: query failed: %w", op, mapErrors(err))
	}
	defer rows.Close()

	tracks, err := selectTracks(rows)
	if err != nil {
		return nil, fmt.Errorf("[%s]: scan failed: %w", op, mapErrors(err))
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("[%s]: rows iteration failed: %w", op, mapErrors(err))
	}
	return tracks, nil
}

func (r *Repository) GetByGenreID(ctx context.Context, genreID uuid.UUID, limit, offset uint64) ([]model.Track, error) {
	const op = "repository.GetByGenreID"
	query := `
		SELECT t.track_id, t.title, t.duration_s, t.file_url, t.play_count, t.description, t.created_at, t.updated_at
		FROM track t
		JOIN track_genre tg ON t.track_id = tg.track_id
		WHERE tg.genre_id = $1
		ORDER BY t.created_at
		LIMIT $2 OFFSET $3`
	rows, err := r.db.QueryContext(ctx, query, genreID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[%s]: query failed: %w", op, mapErrors(err))
	}
	defer rows.Close()

	tracks, err := selectTracks(rows)
	if err != nil {
		return nil, fmt.Errorf("[%s]: scan failed: %w", op, mapErrors(err))
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("[%s]: rows iteration failed: %w", op, mapErrors(err))
	}
	return tracks, nil
}

func (r *Repository) GetAlbumIDsForTracks(ctx context.Context, trackIDs []uuid.UUID) (map[uuid.UUID]uuid.UUID, error) {
	const op = "repository.GetAlbumIDsForTracks"
	if len(trackIDs) == 0 {
		return nil, nil
	}
	query := `SELECT track_id, album_id FROM track_album WHERE track_id = ANY($1)`
	rows, err := r.db.QueryContext(ctx, query, trackIDs)
	if err != nil {
		return nil, fmt.Errorf("[%s]: query failed: %w", op, mapErrors(err))
	}
	defer rows.Close()

	result := make(map[uuid.UUID]uuid.UUID)
	for rows.Next() {
		var trackID, albumID uuid.UUID
		if err := rows.Scan(&trackID, &albumID); err != nil {
			return nil, fmt.Errorf("[%s]: scan failed: %w", op, mapErrors(err))
		}
		result[trackID] = albumID
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("[%s]: rows iteration failed: %w", op, mapErrors(err))
	}
	return result, nil
}

func (r *Repository) GetArtistIDsForTracks(ctx context.Context, trackIDs []uuid.UUID) (map[uuid.UUID][]uuid.UUID, error) {
	const op = "repository.GetArtistIDsForTracks"
	if len(trackIDs) == 0 {
		return nil, nil
	}
	query := `SELECT track_id, artist_id FROM track_artist WHERE track_id = ANY($1)`
	rows, err := r.db.QueryContext(ctx, query, trackIDs)
	if err != nil {
		return nil, fmt.Errorf("[%s]: query failed: %w", op, mapErrors(err))
	}
	defer rows.Close()

	result := make(map[uuid.UUID][]uuid.UUID)
	for rows.Next() {
		var trackID, artistID uuid.UUID
		if err := rows.Scan(&trackID, &artistID); err != nil {
			return nil, fmt.Errorf("[%s]: scan failed: %w", op, mapErrors(err))
		}
		result[trackID] = append(result[trackID], artistID)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("[%s]: rows iteration failed: %w", op, mapErrors(err))
	}
	return result, nil
}

func (r *Repository) GetGenresForTracks(ctx context.Context, trackIDs []uuid.UUID) (map[uuid.UUID][]model.Genre, error) {
	const op = "repository.GetGenresForTracks"
	if len(trackIDs) == 0 {
		return nil, nil
	}
	query := `
		SELECT tg.track_id, g.genre_id, g.genre_name, g.description, g.created_at
		FROM genre g
		JOIN track_genre tg ON g.genre_id = tg.genre_id
		WHERE tg.track_id = ANY($1)`
	rows, err := r.db.QueryContext(ctx, query, trackIDs)
	if err != nil {
		return nil, fmt.Errorf("[%s]: query failed: %w", op, mapErrors(err))
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
			return nil, fmt.Errorf("[%s]: scan failed: %w", op, mapErrors(err))
		}
		result[trackID] = append(result[trackID], genre)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("[%s]: rows iteration failed: %w", op, mapErrors(err))
	}
	return result, nil
}

func (r *Repository) IncrementPlayCount(ctx context.Context, trackID uuid.UUID) error {
	const op = "repository.IncrementPlayCount"
	query := `UPDATE track SET play_count = play_count + 1 WHERE track_id = $1`

	result, err := r.db.ExecContext(ctx, query, trackID)
	if err != nil {
		return fmt.Errorf("[%s]: %w", op, mapErrors(err))
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("[%s]: could not get rows affected: %w", op, mapErrors(err))
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *Repository) GetTotalPlaysByArtistID(ctx context.Context, artistID uuid.UUID) (int64, error) {
	const op = "repository.GetTotalPlaysByArtistID"
	query := `
		SELECT COALESCE(SUM(t.play_count), 0)
		FROM track t
		JOIN track_artist ta ON t.track_id = ta.track_id
		WHERE ta.artist_id = $1`

	var totalPlays int64
	err := r.db.QueryRowContext(ctx, query, artistID).Scan(&totalPlays)
	if err != nil {
		return 0, fmt.Errorf("[%s]: %w", op, mapErrors(err))
	}

	return totalPlays, nil
}

func (r *Repository) GetTotalPlaysByArtistIDs(ctx context.Context, artistIDs []uuid.UUID) (map[uuid.UUID]int64, error) {
	const op = "repository.GetTotalPlaysByArtistIDs"
	if len(artistIDs) == 0 {
		return nil, nil
	}

	if len(artistIDs) > maxIDsInBatch {
		return nil, fmt.Errorf("[%s]: too many IDs requested", op)
	}

	query := `
        SELECT ta.artist_id, COALESCE(SUM(t.play_count), 0) as total_plays
        FROM track_artist ta
        JOIN track t ON t.track_id = ta.track_id
        WHERE ta.artist_id = ANY($1)
        GROUP BY ta.artist_id`

	rows, err := r.db.QueryContext(ctx, query, artistIDs)
	if err != nil {
		return nil, fmt.Errorf("[%s]: query failed: %w", op, mapErrors(err))
	}
	defer rows.Close()

	playsMap := make(map[uuid.UUID]int64)
	for rows.Next() {
		var artistID uuid.UUID
		var totalPlays int64
		if err := rows.Scan(&artistID, &totalPlays); err != nil {
			return nil, fmt.Errorf("[%s]: scan failed: %w", op, mapErrors(err))
		}
		playsMap[artistID] = totalPlays
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("[%s]: rows iteration failed: %w", op, mapErrors(err))
	}

	for _, id := range artistIDs {
		if _, ok := playsMap[id]; !ok {
			playsMap[id] = 0
		}
	}

	return playsMap, nil
}

func selectTracks(rows *sql.Rows) ([]model.Track, error) {
	tracks := make([]model.Track, 0)
	for rows.Next() {
		var track model.Track
		if err := rows.Scan(
			&track.ID,
			&track.Title,
			&track.DurationS,
			&track.FileURL,
			&track.PlayCount,
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
