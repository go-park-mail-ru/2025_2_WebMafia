package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	model "spotify/internal/models"

	"github.com/google/uuid"
)

const baseQuery = `
		SELECT
			t.track_id, t.title, t.duration_ms, t.file_url, t.description, t.created_at, t.updated_at,
			al.album_id, al.title, al.avatar_url, al.description, al.release_date, al.created_at, al.updated_at,
			COALESCE(
				(SELECT json_agg(json_build_object(
					'artist_id', ar.artist_id,
					'artist_name', ar.artist_name,
					'avatar_url', ar.avatar_url
				))
				FROM artists ar
				JOIN track_artist ta ON ar.artist_id = ta.artist_id
				WHERE ta.track_id = t.track_id),
			'[]') AS artists,
			COALESCE(
				(SELECT json_agg(json_build_object(
					'genre_id', g.genre_id,
					'genre_name', g.genre_name
				))
				FROM genres g
				JOIN track_genre tg ON g.genre_id = tg.genre_id
				WHERE tg.track_id = t.track_id),
			'[]') AS genres
		FROM tracks t
		LEFT JOIN track_album tal ON t.track_id = tal.track_id
		LEFT JOIN albums al ON tal.album_id = al.album_id
	`

func (r *Repository) selectTracks(ctx context.Context, whereClause string, args ...interface{}) ([]model.Track, []model.Album, [][]model.Artist, [][]model.Genre, error) {

	query := baseQuery + whereClause

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("repository.selectTracks: query: %w", err)
	}
	defer rows.Close()

	var tracks []model.Track
	var albums []model.Album
	var artistsSlice [][]model.Artist
	var genresSlice [][]model.Genre

	for rows.Next() {
		var track model.Track
		var album model.Album
		var artistsJSON, genresJSON []byte

		if err := rows.Scan(
			&track.ID, &track.Title, &track.DurationMs, &track.FileURL, &track.Description, &track.CreatedAt, &track.UpdatedAt,
			&album.ID, &album.Title, &album.AvatarURL, &album.Description, &album.ReleaseDate, &album.CreatedAt, &album.UpdatedAt,
			&artistsJSON,
			&genresJSON,
		); err != nil {
			return nil, nil, nil, nil, fmt.Errorf("repository.selectTracks: scan: %w", err)
		}

		var currentArtists []model.Artist
		if err := json.Unmarshal(artistsJSON, &currentArtists); err != nil {
			return nil, nil, nil, nil, fmt.Errorf("repository.selectTracks: unmarshal artists: %w", err)
		}

		var currentGenres []model.Genre
		if err := json.Unmarshal(genresJSON, &currentGenres); err != nil {
			return nil, nil, nil, nil, fmt.Errorf("repository.selectTracks: unmarshal genres: %w", err)
		}

		tracks = append(tracks, track)
		albums = append(albums, album)
		artistsSlice = append(artistsSlice, currentArtists)
		genresSlice = append(genresSlice, currentGenres)
	}

	if err := rows.Err(); err != nil {
		return nil, nil, nil, nil, fmt.Errorf("repository.selectTracks: rows error: %w", err)
	}

	return tracks, albums, artistsSlice, genresSlice, nil
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*model.Track, *model.Album, []model.Artist, []model.Genre, error) {
	whereClause := " WHERE t.track_id = $1"
	tracks, albums, artists, genres, err := r.selectTracks(ctx, whereClause, id)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	if len(tracks) == 0 {
		return nil, nil, nil, nil, ErrNotFound
	}
	return &tracks[0], &albums[0], artists[0], genres[0], nil
}

func (r *Repository) GetAll(ctx context.Context) ([]model.Track, []model.Album, [][]model.Artist, [][]model.Genre, error) {
	return r.selectTracks(ctx, "")
}

func (r *Repository) GetByArtistID(ctx context.Context, artistID uuid.UUID) ([]model.Track, []model.Album, [][]model.Artist, [][]model.Genre, error) {
	whereClause := `
		JOIN track_artist ta_filter ON t.track_id = ta_filter.track_id
		WHERE ta_filter.artist_id = $1`
	return r.selectTracks(ctx, whereClause, artistID)
}

func (r *Repository) GetByAlbumID(ctx context.Context, albumID uuid.UUID) ([]model.Track, []model.Album, [][]model.Artist, [][]model.Genre, error) {
	whereClause := " WHERE al.album_id = $1"
	return r.selectTracks(ctx, whereClause, albumID)
}

func (r *Repository) GetByGenreID(ctx context.Context, genreID uuid.UUID) ([]model.Track, []model.Album, [][]model.Artist, [][]model.Genre, error) {
	whereClause := `
		JOIN track_genre tg_filter ON t.track_id = tg_filter.track_id
		WHERE tg_filter.genre_id = $1`
	return r.selectTracks(ctx, whereClause, genreID)
}
