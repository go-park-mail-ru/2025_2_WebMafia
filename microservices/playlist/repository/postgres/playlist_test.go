package postgres

import (
	"context"
	"database/sql"
	"errors"
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

	pl := model.Playlist{
		ID:    uuid.New(),
		Title: "P1",
	}
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
	uid := uuid.New()
	now := time.Now()

	rows := sqlmock.NewRows([]string{
		"playlist_id",
		"user_id",
		"title",
		"description",
		"avatar_url",
		"is_favorite",
		"created_at",
		"updated_at",
	}).AddRow(
		id,
		uid,
		"Title",
		"Desc",
		"url",
		false,
		now,
		now,
	)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT playlist_id`)).
		WithArgs(id).
		WillReturnRows(rows)

	res, err := repo.GetByID(context.Background(), id)
	assert.NoError(t, err)
	assert.Equal(t, "Title", res.Title)
	assert.Equal(t, uid, res.UserID)
}

func TestRepository_GetAllByUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	uid := uuid.New()
	now := time.Now()

	rows := sqlmock.NewRows([]string{
		"playlist_id",
		"user_id",
		"title",
		"description",
		"avatar_url",
		"is_favorite",
		"created_at",
		"updated_at",
	}).AddRow(
		uuid.New(),
		uid,
		"P1",
		"",
		"",
		false,
		now,
		now,
	)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT playlist_id`)).
		WithArgs(uid, uint64(10), uint64(0)).
		WillReturnRows(rows)

	res, err := repo.GetAllByUser(context.Background(), uid, 10, 0)
	assert.NoError(t, err)
	assert.Len(t, res, 1)
	assert.Equal(t, "P1", res[0].Title)
}

func TestRepository_UpdatePlaylist(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	id := uuid.New()
	title := "New"

	mock.ExpectExec(
		regexp.QuoteMeta(`UPDATE playlist SET title = $1 WHERE playlist_id = $2`)).
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

	rows := sqlmock.NewRows([]string{"track_id"}).AddRow("t1")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT track_id FROM playlist_track`)).
		WithArgs(id).
		WillReturnRows(rows)

	res, err := repo.GetTracksByPlaylist(context.Background(), id)
	assert.NoError(t, err)
	assert.Equal(t, []string{"t1"}, res)
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

func TestRepository_AddAlbumToFavorite(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	uid := uuid.New()

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO favorite_album`)).
		WithArgs(uid, "a1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.AddAlbumToFavorite(context.Background(), uid, "a1")
	assert.NoError(t, err)
}

func TestRepository_AddAlbumToFavorite_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	uid := uuid.New()

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO favorite_album`)).
		WithArgs(uid, "a1").
		WillReturnError(errors.New("db error"))

	err = repo.AddAlbumToFavorite(context.Background(), uid, "a1")
	assert.Error(t, err)
}

func TestRepository_RemoveAlbumFromFavorite(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	uid := uuid.New()

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM favorite_album`)).
		WithArgs(uid, "a1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.RemoveAlbumFromFavorite(context.Background(), uid, "a1")
	assert.NoError(t, err)
}

func TestRepository_RemoveAlbumFromFavorite_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	uid := uuid.New()

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM favorite_album`)).
		WithArgs(uid, "a1").
		WillReturnResult(sqlmock.NewResult(1, 0))

	err = repo.RemoveAlbumFromFavorite(context.Background(), uid, "a1")
	assert.ErrorIs(t, err, ErrNotFound)
}

func TestRepository_GetFavoriteAlbumIDs(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	uid := uuid.New()
	now := time.Now()

	a1 := uuid.New()
	a2 := uuid.New()

	rows := sqlmock.NewRows([]string{"album_id", "created_at"}).
		AddRow(a1, now).
		AddRow(a2, now)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT album_id, created_at FROM favorite_album`)).
		WithArgs(uid).
		WillReturnRows(rows)

	res, err := repo.GetFavoriteAlbumIDs(context.Background(), uid)
	assert.NoError(t, err)

	require.Len(t, res, 2)
	assert.Equal(t, uid, res[0].UserID)
	assert.Equal(t, a1, res[0].AlbumID)
	assert.Equal(t, now, res[0].CreatedAt)

	assert.Equal(t, a2, res[1].AlbumID)
	assert.Equal(t, now, res[1].CreatedAt)
}

func TestRepository_AddArtistToFavorite(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	uid := uuid.New()

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO favorite_artist`)).
		WithArgs(uid, "art1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.AddArtistToFavorite(context.Background(), uid, "art1")
	assert.NoError(t, err)
}

func TestRepository_RemoveArtistFromFavorite(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	uid := uuid.New()

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM favorite_artist`)).
		WithArgs(uid, "art1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.RemoveArtistFromFavorite(context.Background(), uid, "art1")
	assert.NoError(t, err)
}

func TestRepository_GetFavoriteArtistIDs(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	uid := uuid.New()
	now := time.Now()

	aid := uuid.New()

	rows := sqlmock.NewRows([]string{"artist_id", "created_at"}).
		AddRow(aid, now)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT artist_id, created_at FROM favorite_artist`)).
		WithArgs(uid).
		WillReturnRows(rows)

	res, err := repo.GetFavoriteArtistIDs(context.Background(), uid)
	assert.NoError(t, err)

	require.Len(t, res, 1)
	assert.Equal(t, uid, res[0].UserID)
	assert.Equal(t, aid, res[0].ArtistID)
	assert.Equal(t, now, res[0].CreatedAt)
}
