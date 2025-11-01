package postgres

import (
	"context"
	"database/sql"
	"database/sql/driver"
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

func newMockTrack() *model.Track {
	return &model.Track{
		ID:          uuid.New(),
		Title:       "Test Track",
		DurationMs:  180000,
		FileURL:     "http://example.com/track.mp3",
		Description: sql.NullString{String: "A test track", Valid: true},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func TestTrackRepository_SimpleGets(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	mockTrack := newMockTrack()
	columns := []string{"track_id", "title", "duration_ms", "file_url", "description", "created_at", "updated_at"}
	baseQuery := regexp.QuoteMeta(`SELECT track_id, title, duration_ms, file_url, description, created_at, updated_at FROM track`)
	joinQuery := regexp.QuoteMeta(`SELECT t.track_id, t.title, t.duration_ms, t.file_url, t.description, t.created_at, t.updated_at FROM track t`)

	t.Run("GetByID", func(t *testing.T) {
		rows := sqlmock.NewRows(columns).
			AddRow(mockTrack.ID, mockTrack.Title, mockTrack.DurationMs, mockTrack.FileURL, mockTrack.Description, mockTrack.CreatedAt, mockTrack.UpdatedAt)
		mock.ExpectQuery(baseQuery).WithArgs(mockTrack.ID).WillReturnRows(rows)

		track, err := repo.GetByID(context.Background(), mockTrack.ID)
		assert.NoError(t, err)
		assert.Equal(t, mockTrack, track)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetAll", func(t *testing.T) {
		rows := sqlmock.NewRows(columns).
			AddRow(mockTrack.ID, mockTrack.Title, mockTrack.DurationMs, mockTrack.FileURL, mockTrack.Description, mockTrack.CreatedAt, mockTrack.UpdatedAt)
		mock.ExpectQuery(baseQuery).WithArgs(uint64(10), uint64(0)).WillReturnRows(rows)

		tracks, err := repo.GetAll(context.Background(), 10, 0)
		assert.NoError(t, err)
		require.Len(t, tracks, 1)
		assert.Equal(t, *mockTrack, tracks[0])
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetByArtistID", func(t *testing.T) {
		artistID := uuid.New()
		rows := sqlmock.NewRows(columns).
			AddRow(mockTrack.ID, mockTrack.Title, mockTrack.DurationMs, mockTrack.FileURL, mockTrack.Description, mockTrack.CreatedAt, mockTrack.UpdatedAt)
		mock.ExpectQuery(joinQuery).WithArgs(artistID, uint64(10), uint64(0)).WillReturnRows(rows)

		tracks, err := repo.GetByArtistID(context.Background(), artistID, 10, 0)
		assert.NoError(t, err)
		require.Len(t, tracks, 1)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetByAlbumID", func(t *testing.T) {
		albumID := uuid.New()
		rows := sqlmock.NewRows(columns).
			AddRow(mockTrack.ID, mockTrack.Title, mockTrack.DurationMs, mockTrack.FileURL, mockTrack.Description, mockTrack.CreatedAt, mockTrack.UpdatedAt)
		mock.ExpectQuery(joinQuery).WithArgs(albumID, uint64(10), uint64(0)).WillReturnRows(rows)

		tracks, err := repo.GetByAlbumID(context.Background(), albumID, 10, 0)
		assert.NoError(t, err)
		require.Len(t, tracks, 1)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetByGenreID", func(t *testing.T) {
		genreID := uuid.New()
		rows := sqlmock.NewRows(columns).
			AddRow(mockTrack.ID, mockTrack.Title, mockTrack.DurationMs, mockTrack.FileURL, mockTrack.Description, mockTrack.CreatedAt, mockTrack.UpdatedAt)
		mock.ExpectQuery(joinQuery).WithArgs(genreID, uint64(10), uint64(0)).WillReturnRows(rows)

		tracks, err := repo.GetByGenreID(context.Background(), genreID, 10, 0)
		assert.NoError(t, err)
		require.Len(t, tracks, 1)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetByID returns error on db failure", func(t *testing.T) {
		expectedError := sql.ErrConnDone
		mock.ExpectQuery(baseQuery).WithArgs(mockTrack.ID).WillReturnError(expectedError)

		_, err := repo.GetByID(context.Background(), mockTrack.ID)
		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestTrackRepository_JunctionTableGets(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()
	repo := New(db)

	trackID1 := uuid.New()
	trackID2 := uuid.New()
	trackIDs := []uuid.UUID{trackID1, trackID2}

	t.Run("GetAlbumIDsForTracks", func(t *testing.T) {
		albumID1 := uuid.New()
		query := regexp.QuoteMeta(`SELECT track_id, album_id FROM track_album WHERE track_id = ANY($1)`)
		rows := sqlmock.NewRows([]string{"track_id", "album_id"}).AddRow(trackID1, albumID1)
		mock.ExpectQuery(query).WithArgs(trackIDs).WillReturnRows(rows)

		result, err := repo.GetAlbumIDsForTracks(context.Background(), trackIDs)
		assert.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, albumID1, result[trackID1])
		_, ok := result[trackID2]
		assert.False(t, ok)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetArtistIDsForTracks", func(t *testing.T) {
		artistID1 := uuid.New()
		artistID2 := uuid.New()
		query := regexp.QuoteMeta(`SELECT track_id, artist_id FROM track_artist WHERE track_id = ANY($1)`)
		rows := sqlmock.NewRows([]string{"track_id", "artist_id"}).
			AddRow(trackID1, artistID1).
			AddRow(trackID1, artistID2)
		mock.ExpectQuery(query).WithArgs(trackIDs).WillReturnRows(rows)

		result, err := repo.GetArtistIDsForTracks(context.Background(), trackIDs)
		assert.NoError(t, err)
		require.NotNil(t, result)
		assert.ElementsMatch(t, []uuid.UUID{artistID1, artistID2}, result[trackID1])
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetGenresForTracks", func(t *testing.T) {
		genre := model.Genre{ID: uuid.New(), Name: "Rock"}
		query := regexp.QuoteMeta(`SELECT tg.track_id, g.genre_id, g.genre_name, g.description, g.created_at FROM genre g JOIN track_genre tg ON g.genre_id = tg.genre_id WHERE tg.track_id = ANY($1)`)
		rows := sqlmock.NewRows([]string{"track_id", "genre_id", "genre_name", "description", "created_at"}).
			AddRow(trackID1, genre.ID, genre.Name, genre.Description, genre.CreatedAt)
		mock.ExpectQuery(query).WithArgs(trackIDs).WillReturnRows(rows)

		result, err := repo.GetGenresForTracks(context.Background(), trackIDs)
		assert.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, []model.Genre{genre}, result[trackID1])
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetAlbumIDsForTracks - empty input", func(t *testing.T) {
		result, err := repo.GetAlbumIDsForTracks(context.Background(), []uuid.UUID{})
		assert.NoError(t, err)
		assert.Nil(t, result)
	})

	t.Run("GetArtistIDsForTracks - db error", func(t *testing.T) {
		query := regexp.QuoteMeta(`SELECT track_id, artist_id FROM track_artist WHERE track_id = ANY($1)`)
		mock.ExpectQuery(query).WithArgs(trackIDs).WillReturnError(errors.New("db error"))
		_, err := repo.GetArtistIDsForTracks(context.Background(), trackIDs)
		assert.Error(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetGenresForTracks - scan error", func(t *testing.T) {
		query := regexp.QuoteMeta(`SELECT tg.track_id, g.genre_id, g.genre_name, g.description, g.created_at FROM genre g JOIN track_genre tg ON g.genre_id = tg.genre_id WHERE tg.track_id = ANY($1)`)
		rows := sqlmock.NewRows([]string{"track_id", "genre_id"}).AddRow(trackID1, "not-a-uuid")
		mock.ExpectQuery(query).WithArgs(trackIDs).WillReturnRows(rows)

		_, err := repo.GetGenresForTracks(context.Background(), trackIDs)
		assert.Error(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
