package postgres

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepository_GetAlbumByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"album_id", "title", "type", "avatar_url", "artist_id", "description", "release_date", "created_at", "updated_at"}).
			AddRow(id, "Title", "LP", "url", uuid.New(), sql.NullString{String: "desc", Valid: true}, time.Now(), time.Now(), time.Now())

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT album_id, title`)).
			WithArgs(id).
			WillReturnRows(rows)

		res, err := repo.GetAlbumByID(context.Background(), id)
		assert.NoError(t, err)
		assert.Equal(t, "Title", res.Title)
	})

	t.Run("not found", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT album_id`)).
			WithArgs(id).
			WillReturnError(sql.ErrNoRows)

		_, err := repo.GetAlbumByID(context.Background(), id)
		assert.ErrorIs(t, err, ErrNotFound)
	})
}

func TestRepository_GetAllAlbums(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"album_id", "title", "type", "avatar_url", "artist_id", "description", "release_date", "created_at", "updated_at"}).
			AddRow(uuid.New(), "Title", "LP", "url", uuid.New(), sql.NullString{}, time.Now(), time.Now(), time.Now())

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT album_id, title`)).
			WithArgs(uint64(10), uint64(0)).
			WillReturnRows(rows)

		res, err := repo.GetAllAlbums(context.Background(), 10, 0)
		assert.NoError(t, err)
		assert.Len(t, res, 1)
	})
}

func TestRepository_GetAlbumsByArtistID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"album_id", "title", "type", "avatar_url", "artist_id", "description", "release_date", "created_at", "updated_at"}).
			AddRow(uuid.New(), "Title", "LP", "url", id, sql.NullString{}, time.Now(), time.Now(), time.Now())

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT album_id, title`)).
			WithArgs(id, uint64(10), uint64(0)).
			WillReturnRows(rows)

		res, err := repo.GetAlbumsByArtistID(context.Background(), id, 10, 0)
		assert.NoError(t, err)
		assert.Len(t, res, 1)
	})
}

func TestRepository_GetAlbumsByIDs(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"album_id", "title", "type", "avatar_url", "artist_id", "description", "release_date", "created_at", "updated_at"}).
			AddRow(id, "Title", "LP", "url", uuid.New(), sql.NullString{}, time.Now(), time.Now(), time.Now())

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT album_id, title`)).
			WithArgs([]uuid.UUID{id}).
			WillReturnRows(rows)

		res, err := repo.GetAlbumsByIDs(context.Background(), []uuid.UUID{id})
		assert.NoError(t, err)
		assert.Len(t, res, 1)
	})

	t.Run("db error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT album_id`)).
			WithArgs([]uuid.UUID{id}).
			WillReturnError(errors.New("db error"))

		_, err := repo.GetAlbumsByIDs(context.Background(), []uuid.UUID{id})
		assert.Error(t, err)
	})
}

func TestRepository_SearchAlbums(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"album_id", "title", "type", "avatar_url", "artist_id", "description", "release_date", "created_at", "updated_at", "rank"}).
			AddRow(uuid.New(), "Title", "LP", "url", uuid.New(), sql.NullString{}, time.Now(), time.Now(), time.Now(), 0.8)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT album_id, title`)).
			WithArgs("query", uint64(10)).
			WillReturnRows(rows)

		res, err := repo.SearchAlbums(context.Background(), "query", 10)
		assert.NoError(t, err)
		assert.Len(t, res, 1)
		assert.Equal(t, float32(0.8), res[0].Rank)
	})
}
