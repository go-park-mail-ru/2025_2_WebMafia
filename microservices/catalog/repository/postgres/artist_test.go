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

func TestRepository_GetArtistByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"artist_id", "artist_name", "avatar_url", "header_url", "description", "created_at", "updated_at"}).
			AddRow(id, "Name", "url", "h_url", sql.NullString{String: "desc", Valid: true}, time.Now(), time.Now())

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT artist_id, artist_name`)).
			WithArgs(id).
			WillReturnRows(rows)

		res, err := repo.GetArtistByID(context.Background(), id)
		assert.NoError(t, err)
		assert.Equal(t, "Name", res.Name)
	})

	t.Run("not found", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT artist_id`)).
			WithArgs(id).
			WillReturnError(sql.ErrNoRows)

		_, err := repo.GetArtistByID(context.Background(), id)
		assert.ErrorIs(t, err, ErrNotFound)
	})
}

func TestRepository_GetAllArtists(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"artist_id", "artist_name", "avatar_url", "header_url", "description", "created_at", "updated_at"}).
			AddRow(uuid.New(), "Name", "url", "h", sql.NullString{}, time.Now(), time.Now())

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT artist_id, artist_name`)).
			WithArgs(uint64(10), uint64(0)).
			WillReturnRows(rows)

		res, err := repo.GetAllArtists(context.Background(), 10, 0)
		assert.NoError(t, err)
		assert.Len(t, res, 1)
	})
}

func TestRepository_GetArtistsByIDs(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"artist_id", "artist_name", "avatar_url", "header_url", "description", "created_at", "updated_at"}).
			AddRow(id, "Name", "url", "h", sql.NullString{}, time.Now(), time.Now())

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT artist_id, artist_name`)).
			WithArgs([]uuid.UUID{id}).
			WillReturnRows(rows)

		res, err := repo.GetArtistsByIDs(context.Background(), []uuid.UUID{id})
		assert.NoError(t, err)
		assert.Len(t, res, 1)
	})

	t.Run("db error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT artist_id`)).
			WithArgs([]uuid.UUID{id}).
			WillReturnError(errors.New("db error"))

		_, err := repo.GetArtistsByIDs(context.Background(), []uuid.UUID{id})
		assert.Error(t, err)
	})
}

func TestRepository_SearchArtists(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"artist_id", "artist_name", "avatar_url", "header_url", "description", "created_at", "updated_at", "rank"}).
			AddRow(uuid.New(), "Name", "url", "h", sql.NullString{}, time.Now(), time.Now(), 0.5)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT artist_id, artist_name`)).
			WithArgs("query", uint64(10)).
			WillReturnRows(rows)

		res, err := repo.SearchArtists(context.Background(), "query", 10)
		assert.NoError(t, err)
		assert.Len(t, res, 1)
		assert.Equal(t, float32(0.5), res[0].Rank)
	})
}
