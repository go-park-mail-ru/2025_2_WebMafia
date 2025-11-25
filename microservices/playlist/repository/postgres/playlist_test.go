package postgres

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"spotify/internal/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepository_CreatePlaylist(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := New(db)

	pl := model.Playlist{ID: uuid.New(), Title: "P1"}
	uid := uuid.New()

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO playlist`)).
		WithArgs(pl.ID, uid, pl.Title, pl.Description, pl.AvatarURL, pl.IsFavorite).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreatePlaylist(context.Background(), pl, uid)
	assert.NoError(t, err)
}

func TestRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := New(db)
	id := uuid.New()

	rows := sqlmock.NewRows([]string{"playlist_id", "title", "description", "avatar_url", "is_favorite", "created_at", "updated_at"}).
		AddRow(id, "Title", "Desc", "url", false, time.Now(), time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT playlist_id`)).WithArgs(id).WillReturnRows(rows)

	res, err := repo.GetByID(context.Background(), id)
	assert.NoError(t, err)
	assert.Equal(t, "Title", res.Title)
}

func TestRepository_GetAllByUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := New(db)
	uid := uuid.New()

	rows := sqlmock.NewRows([]string{"playlist_id", "title", "description", "avatar_url", "is_favorite", "created_at", "updated_at"}).
		AddRow(uuid.New(), "P1", "", "", false, time.Now(), time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT playlist_id`)).
		WithArgs(uid, uint64(10), uint64(0)).
		WillReturnRows(rows)

	res, err := repo.GetAllByUser(context.Background(), uid, 10, 0)
	assert.NoError(t, err)
	assert.Len(t, res, 1)
}

func TestRepository_UpdatePlaylist(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := New(db)
	id := uuid.New()
	title := "New"

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE playlist SET title = $1 WHERE playlist_id = $2`)).
		WithArgs(title, id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdatePlaylist(context.Background(), id, PlaylistUpdate{Title: &title})
	assert.NoError(t, err)
}

func TestRepository_DeletePlaylist(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := New(db)
	id := uuid.New()

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM playlist`)).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeletePlaylist(context.Background(), id)
	assert.NoError(t, err)
}

func TestRepository_AddTrackToPlaylist(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := New(db)
	id := uuid.New()

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO playlist_track`)).
		WithArgs(id, "t1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.AddTrackToPlaylist(context.Background(), id, "t1")
	assert.NoError(t, err)
}

func TestRepository_RemoveTrackFromPlaylist(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := New(db)
	id := uuid.New()

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM playlist_track`)).
		WithArgs(id, "t1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.RemoveTrackFromPlaylist(context.Background(), id, "t1")
	assert.NoError(t, err)
}

func TestRepository_GetTracksByPlaylist(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := New(db)
	id := uuid.New()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT track_id FROM playlist_track`)).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"track_id"}).AddRow("t1"))

	res, err := repo.GetTracksByPlaylist(context.Background(), id)
	assert.NoError(t, err)
	assert.Equal(t, "t1", res[0])
}

func TestRepository_UpdatePlaylistAvatar(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := New(db)
	id := uuid.New()

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE playlist SET avatar_url`)).
		WithArgs("url", id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdatePlaylistAvatar(context.Background(), id, "url")
	assert.NoError(t, err)
}

func TestRepository_GetFavoritePlaylist(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := New(db)
	uid := uuid.New()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT playlist_id`)).
		WithArgs(uid).
		WillReturnError(sql.ErrNoRows)

	_, err = repo.GetFavoritePlaylist(context.Background(), uid)
	assert.ErrorIs(t, err, ErrNotFound)
}
