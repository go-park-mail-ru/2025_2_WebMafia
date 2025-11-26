package postgres

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"spotify/internal/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type pgxUUIDValueConverter struct{}

func (c pgxUUIDValueConverter) ConvertValue(v interface{}) (driver.Value, error) {
	if id, ok := v.(uuid.UUID); ok {
		return id, nil
	}
	if ids, ok := v.([]uuid.UUID); ok {
		return ids, nil
	}
	return v, nil
}

func TestRepository_GetTrackByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"track_id", "title", "duration_s", "file_url", "play_count", "description", "created_at", "updated_at"}).
			AddRow(id, "Title", 180, "url", 0, sql.NullString{String: "desc", Valid: true}, time.Now(), time.Now())

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT track_id, title, duration_s, file_url, play_count, description, created_at, updated_at FROM track WHERE track_id = $1`)).
			WithArgs(id).
			WillReturnRows(rows)

		res, err := repo.GetTrackByID(context.Background(), id)
		assert.NoError(t, err)
		assert.Equal(t, "Title", res.Title)
	})

	t.Run("not found", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT track_id`)).
			WithArgs(id).
			WillReturnError(sql.ErrNoRows)

		_, err := repo.GetTrackByID(context.Background(), id)
		assert.ErrorIs(t, err, ErrNotFound)
	})
}

func TestRepository_GetAllTracks(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"track_id", "title", "duration_s", "file_url", "play_count", "description", "created_at", "updated_at"}).
			AddRow(uuid.New(), "Title", 180, "url", 0, sql.NullString{}, time.Now(), time.Now())

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT track_id, title, duration_s, file_url, play_count, description, created_at, updated_at FROM track ORDER BY play_count DESC, created_at DESC LIMIT $1 OFFSET $2`)).
			WithArgs(uint64(10), uint64(0)).
			WillReturnRows(rows)

		res, err := repo.GetAllTracks(context.Background(), 10, 0)
		assert.NoError(t, err)
		assert.Len(t, res, 1)
	})
}

func TestRepository_GetTracksByReferences(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	id := uuid.New()

	newRows := func() *sqlmock.Rows {
		return sqlmock.NewRows([]string{"track_id", "title", "duration_s", "file_url", "play_count", "description", "created_at", "updated_at"}).
			AddRow(uuid.New(), "Title", 180, "url", 0, sql.NullString{}, time.Now(), time.Now())
	}

	t.Run("GetTracksByArtistID", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT t.track_id, t.title`)).
			WithArgs(id, uint64(10), uint64(0)).
			WillReturnRows(newRows())
		tracks, err := repo.GetTracksByArtistID(context.Background(), id, 10, 0)
		assert.NoError(t, err)
		assert.Len(t, tracks, 1)
	})

	t.Run("GetTracksByAlbumID", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT t.track_id, t.title`)).
			WithArgs(id, uint64(10), uint64(0)).
			WillReturnRows(newRows())
		tracks, err := repo.GetTracksByAlbumID(context.Background(), id, 10, 0)
		assert.NoError(t, err)
		assert.Len(t, tracks, 1)
	})

	t.Run("GetTracksByGenreID", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT t.track_id, t.title`)).
			WithArgs(id, uint64(10), uint64(0)).
			WillReturnRows(newRows())
		tracks, err := repo.GetTracksByGenreID(context.Background(), id, 10, 0)
		assert.NoError(t, err)
		assert.Len(t, tracks, 1)
	})
}

func TestRepository_GetAlbumIDsForTracks(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	trackID := uuid.New()
	albumID := uuid.New()

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"track_id", "album_id"}).
			AddRow(trackID, albumID)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT track_id, album_id FROM track_album WHERE track_id = ANY($1)`)).
			WithArgs([]uuid.UUID{trackID}).
			WillReturnRows(rows)

		res, err := repo.GetAlbumIDsForTracks(context.Background(), []uuid.UUID{trackID})
		assert.NoError(t, err)
		assert.Equal(t, albumID, res[trackID])
	})
}

func TestRepository_GetArtistIDsForTracks(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	trackID := uuid.New()
	artistID := uuid.New()

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"track_id", "artist_id"}).
			AddRow(trackID, artistID)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT track_id, artist_id FROM track_artist WHERE track_id = ANY($1)`)).
			WithArgs([]uuid.UUID{trackID}).
			WillReturnRows(rows)

		res, err := repo.GetArtistIDsForTracks(context.Background(), []uuid.UUID{trackID})
		assert.NoError(t, err)
		assert.Equal(t, artistID, res[trackID][0])
	})
}

func TestRepository_GetGenresForTracks(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	trackID := uuid.New()
	genre := model.Genre{ID: uuid.New(), Name: "Rock", Description: sql.NullString{String: "Desc", Valid: true}, CreatedAt: time.Now()}

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"track_id", "genre_id", "genre_name", "description", "created_at"}).
			AddRow(trackID, genre.ID, genre.Name, genre.Description, genre.CreatedAt)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT tg.track_id, g.genre_id, g.genre_name, g.description, g.created_at`)).
			WithArgs([]uuid.UUID{trackID}).
			WillReturnRows(rows)

		res, err := repo.GetGenresForTracks(context.Background(), []uuid.UUID{trackID})
		assert.NoError(t, err)
		assert.Equal(t, genre.Name, res[trackID][0].Name)
	})
}

func TestRepository_GetTotalPlaysByArtistIDs(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	artistID := uuid.New()

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"artist_id", "total_plays"}).
			AddRow(artistID, 100)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT ta.artist_id, COALESCE(SUM(t.play_count), 0)`)).
			WithArgs([]uuid.UUID{artistID}).
			WillReturnRows(rows)

		res, err := repo.GetTotalPlaysByArtistIDs(context.Background(), []uuid.UUID{artistID})
		assert.NoError(t, err)
		assert.Equal(t, int64(100), res[artistID])
	})
}

func TestRepository_IncrementPlayCount(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE track SET play_count = play_count + 1`)).
			WithArgs(id).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.IncrementPlayCount(context.Background(), id)
		assert.NoError(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE track`)).
			WithArgs(id).
			WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.IncrementPlayCount(context.Background(), id)
		assert.Equal(t, ErrNotFound, err)
	})
}

func TestRepository_SearchTracks(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"track_id", "title", "duration_s", "file_url", "play_count", "description", "created_at", "updated_at", "rank"}).
			AddRow(uuid.New(), "Title", 180, "url", 0, sql.NullString{}, time.Now(), time.Now(), 0.9)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT track_id, title`)).
			WithArgs("query", uint64(10)).
			WillReturnRows(rows)

		res, err := repo.SearchTracks(context.Background(), "query", 10)
		assert.NoError(t, err)
		assert.Len(t, res, 1)
		assert.Equal(t, float32(0.9), res[0].Rank)
	})
}
