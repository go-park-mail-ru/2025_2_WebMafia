package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"

	"spotify/internal/model"
)

func (r *Repository) CreatePlaylist(ctx context.Context, playlist model.Playlist, userID uuid.UUID) error {
	const op = "repository.CreatePlaylist"

	query := `INSERT INTO playlist (playlist_id, user_id, title, description, avatar_url, is_favorite)
	          VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.Conn.ExecContext(
		ctx,
		query,
		playlist.ID,
		userID,
		playlist.Title,
		playlist.Description,
		playlist.AvatarURL,
		playlist.IsFavorite,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*model.Playlist, error) {
	const op = "repository.GetByID"

	query := `SELECT playlist_id, user_id, title, description, avatar_url, is_favorite, created_at, updated_at
	          FROM playlist
	          WHERE playlist_id = $1`

	playlist, err := r.selectPlaylist(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return playlist, nil
}

func (r *Repository) GetAllByUser(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]model.Playlist, error) {
	const op = "repository.GetAllByUser"

	query := `SELECT playlist_id, user_id, title, description, avatar_url, is_favorite, created_at, updated_at
	          FROM playlist
	          WHERE user_id = $1
	          ORDER BY created_at DESC
	          LIMIT $2 OFFSET $3`

	rows, err := r.Conn.QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	playlists := make([]model.Playlist, 0)
	for rows.Next() {
		var playlist model.Playlist
		if err := rows.Scan(
			&playlist.ID,
			&playlist.UserID,
			&playlist.Title,
			&playlist.Description,
			&playlist.AvatarURL,
			&playlist.IsFavorite,
			&playlist.CreatedAt,
			&playlist.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("%s: scan failed: %w", op, err)
		}
		playlists = append(playlists, playlist)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: iteration failed: %w", op, err)
	}
	return playlists, nil
}

type PlaylistUpdate struct {
	Title       *string
	Description *string
	IsFavorite  *bool
}

func (r *Repository) UpdatePlaylist(ctx context.Context, id uuid.UUID, upd PlaylistUpdate) error {
	const op = "repository.UpdatePlaylist"

	setParts := []string{}
	args := []interface{}{}
	i := 1

	if upd.Title != nil {
		setParts = append(setParts, fmt.Sprintf("title = $%d", i))
		args = append(args, *upd.Title)
		i++
	}
	if upd.Description != nil {
		setParts = append(setParts, fmt.Sprintf("description = $%d", i))
		args = append(args, *upd.Description)
		i++
	}
	if upd.IsFavorite != nil {
		setParts = append(setParts, fmt.Sprintf("is_favorite = $%d", i))
		args = append(args, *upd.IsFavorite)
		i++
	}

	if len(setParts) == 0 {
		return fmt.Errorf("%s: no fields to update", op)
	}

	args = append(args, id)

	query := fmt.Sprintf("UPDATE playlist SET %s WHERE playlist_id = $%d", strings.Join(setParts, ", "), i)

	res, err := r.Conn.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: count failed: %w", op, err)
	}
	if ra == 0 {
		return fmt.Errorf("%s: %w", op, ErrNotFound)
	}

	return nil
}

func (r *Repository) DeletePlaylist(ctx context.Context, id uuid.UUID) error {
	const op = "repository.DeletePlaylist"

	query := `DELETE FROM playlist WHERE playlist_id = $1`

	res, err := r.Conn.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: count failed: %w", op, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, ErrNotFound)
	}

	return nil
}

func (r *Repository) GetFavoritePlaylist(ctx context.Context, userID uuid.UUID) (*model.Playlist, error) {
	query := `SELECT playlist_id, user_id, title, description, avatar_url, is_favorite, created_at, updated_at
	          FROM playlist
	          WHERE user_id = $1 AND is_favorite = true LIMIT 1`

	return r.selectPlaylist(ctx, query, userID)
}

func (r *Repository) AddTrackToPlaylist(ctx context.Context, playlistID uuid.UUID, trackID string) error {
	query := `INSERT INTO playlist_track (playlist_id, track_id)
	          VALUES ($1, $2)
	          ON CONFLICT DO NOTHING`

	_, err := r.Conn.ExecContext(ctx, query, playlistID, trackID)
	if err != nil {
		return fmt.Errorf("add track: %w", err)
	}
	return nil
}

func (r *Repository) RemoveTrackFromPlaylist(ctx context.Context, playlistID uuid.UUID, trackID string) error {
	const op = "repository.RemoveTrackFromPlaylist"

	query := `DELETE FROM playlist_track WHERE playlist_id = $1 AND track_id = $2`

	res, err := r.Conn.ExecContext(ctx, query, playlistID, trackID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if ra == 0 {
		return fmt.Errorf("%s: %w", op, ErrNotFound)
	}

	return nil
}

func (r *Repository) UpdatePlaylistAvatar(ctx context.Context, id uuid.UUID, avatar string) error {
	query := `UPDATE playlist SET avatar_url = $1 WHERE playlist_id = $2`

	_, err := r.Conn.ExecContext(ctx, query, avatar, id)
	if err != nil {
		return fmt.Errorf("update playlist avatar: %w", err)
	}
	return nil
}

func (r *Repository) GetTracksByPlaylist(ctx context.Context, playlistID uuid.UUID) ([]string, error) {
	const op = "repository.GetTracksByPlaylist"

	query := `SELECT track_id FROM playlist_track WHERE playlist_id = $1`

	rows, err := r.Conn.QueryContext(ctx, query, playlistID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	ids := make([]string, 0)
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("%s: scan: %w", op, err)
		}
		ids = append(ids, id)
	}

	return ids, nil
}

// ЛЮБИМЫЕ АЛЬБОМЫ
func (r *Repository) AddAlbumToFavorite(ctx context.Context, userID uuid.UUID, albumID string) error {
	const op = "repository.AddAlbumToFavorite"

	query := `INSERT INTO favorite_album (user_id, album_id)
	          VALUES ($1, $2)
	          ON CONFLICT DO NOTHING`

	_, err := r.Conn.ExecContext(ctx, query, userID, albumID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repository) RemoveAlbumFromFavorite(ctx context.Context, userID uuid.UUID, albumID string) error {
	const op = "repository.RemoveAlbumFromFavorite"

	query := `DELETE FROM favorite_album WHERE user_id = $1 AND album_id = $2`
	res, err := r.Conn.ExecContext(ctx, query, userID, albumID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if ra == 0 {
		return fmt.Errorf("%s: %w", op, ErrNotFound)
	}

	return nil
}

func (r *Repository) GetFavoriteAlbumIDs(ctx context.Context, userID uuid.UUID) ([]string, error) {
	const op = "repository.GetFavoriteAlbumIDs"

	query := `SELECT album_id FROM favorite_album WHERE user_id = $1 ORDER BY created_at DESC`

	rows, err := r.Conn.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	ids := make([]string, 0)
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("%s: scan: %w", op, err)
		}
		ids = append(ids, id)
	}

	return ids, nil
}

// ЛЮБИМЫЕ АРТИСТЫ
func (r *Repository) AddArtistToFavorite(ctx context.Context, userID uuid.UUID, artistID string) error {
	const op = "repository.AddArtistToFavorite"

	query := `INSERT INTO favorite_artist (user_id, artist_id)
	          VALUES ($1, $2) ON CONFLICT DO NOTHING`

	_, err := r.Conn.ExecContext(ctx, query, userID, artistID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repository) RemoveArtistFromFavorite(ctx context.Context, userID uuid.UUID, artistID string) error {
	const op = "repository.RemoveArtistFromFavorite"

	query := `DELETE FROM favorite_artist WHERE user_id = $1 AND artist_id = $2`

	res, err := r.Conn.ExecContext(ctx, query, userID, artistID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if ra == 0 {
		return fmt.Errorf("%s: %w", op, ErrNotFound)
	}

	return nil
}

func (r *Repository) GetFavoriteArtistIDs(ctx context.Context, userID uuid.UUID) ([]string, error) {
	const op = "repository.GetFavoriteArtistIDs"

	query := `SELECT artist_id FROM favorite_artist WHERE user_id = $1 ORDER BY created_at DESC`

	rows, err := r.Conn.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	ids := make([]string, 0)
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("%s: scan: %w", op, err)
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func (r *Repository) selectPlaylist(ctx context.Context, query string, args ...interface{}) (*model.Playlist, error) {
	row := r.Conn.QueryRowContext(ctx, query, args...)

	playlist := &model.Playlist{}
	err := row.Scan(
		&playlist.ID,
		&playlist.UserID,
		&playlist.Title,
		&playlist.Description,
		&playlist.AvatarURL,
		&playlist.IsFavorite,
		&playlist.CreatedAt,
		&playlist.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("playlist not found: %w", ErrNotFound)
	}

	if err != nil {
		return nil, fmt.Errorf("scan failed: %w", err)
	}

	return playlist, nil
}
