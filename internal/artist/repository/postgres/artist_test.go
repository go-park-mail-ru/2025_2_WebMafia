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

func newMockArtist() *model.Artist {
	return &model.Artist{
		ID:          uuid.New(),
		Name:        "Test Artist",
		AvatarURL:   "http://example.com/avatar.jpg",
		HeaderURL:   "http://example.com/header.jpg",
		Description: sql.NullString{String: "A test artist", Valid: true},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func TestArtistRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	mockArtist := newMockArtist()
	query := regexp.QuoteMeta(`SELECT artist_id, artist_name, avatar_url, header_url, description, created_at, updated_at FROM artist WHERE artist_id = $1`)
	columns := []string{"artist_id", "artist_name", "avatar_url", "header_url", "description", "created_at", "updated_at"}

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows(columns).
			AddRow(mockArtist.ID, mockArtist.Name, mockArtist.AvatarURL, mockArtist.HeaderURL, mockArtist.Description, mockArtist.CreatedAt, mockArtist.UpdatedAt)
		mock.ExpectQuery(query).WithArgs(mockArtist.ID).WillReturnRows(rows)

		artist, err := repo.GetByID(context.Background(), mockArtist.ID)

		assert.NoError(t, err)
		assert.Equal(t, mockArtist, artist)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("not found", func(t *testing.T) {
		mock.ExpectQuery(query).WithArgs(mockArtist.ID).WillReturnError(sql.ErrNoRows)

		artist, err := repo.GetByID(context.Background(), mockArtist.ID)

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrNotFound)
		assert.Nil(t, artist)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("db error", func(t *testing.T) {
		expectedError := sql.ErrConnDone
		mock.ExpectQuery(query).WithArgs(mockArtist.ID).WillReturnError(expectedError)

		_, err := repo.GetByID(context.Background(), mockArtist.ID)

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArtistRepository_GetByIDs(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	artist1 := newMockArtist()
	artist2 := newMockArtist()
	ids := []uuid.UUID{artist1.ID, artist2.ID}

	query := regexp.QuoteMeta(`SELECT artist_id, artist_name, avatar_url, header_url, description, created_at, updated_at FROM artist WHERE artist_id = ANY($1)`)
	columns := []string{"artist_id", "artist_name", "avatar_url", "header_url", "description", "created_at", "updated_at"}

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows(columns).
			AddRow(artist1.ID, artist1.Name, artist1.AvatarURL, artist1.HeaderURL, artist1.Description, artist1.CreatedAt, artist1.UpdatedAt).
			AddRow(artist2.ID, artist2.Name, artist2.AvatarURL, artist2.HeaderURL, artist2.Description, artist2.CreatedAt, artist2.UpdatedAt)
		mock.ExpectQuery(query).WithArgs(ids).WillReturnRows(rows)

		artists, err := repo.GetByIDs(context.Background(), ids)

		assert.NoError(t, err)
		require.Len(t, artists, 2)
		assert.Contains(t, artists, *artist1)
		assert.Contains(t, artists, *artist2)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("empty ids slice", func(t *testing.T) {
		artists, err := repo.GetByIDs(context.Background(), []uuid.UUID{})
		assert.NoError(t, err)
		assert.Nil(t, artists)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("db error on query", func(t *testing.T) {
		expectedError := errors.New("db error")
		mock.ExpectQuery(query).WithArgs(ids).WillReturnError(expectedError)

		artists, err := repo.GetByIDs(context.Background(), ids)

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Nil(t, artists)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("db error on rows scan", func(t *testing.T) {
		rows := sqlmock.NewRows(columns).
			AddRow(artist1.ID, artist1.Name, artist1.AvatarURL, artist1.HeaderURL, artist1.Description, artist1.CreatedAt, "not a date")
		mock.ExpectQuery(query).WithArgs(ids).WillReturnRows(rows)

		artists, err := repo.GetByIDs(context.Background(), ids)

		assert.Error(t, err)
		assert.Nil(t, artists)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("db error on rows iteration", func(t *testing.T) {
		expectedError := errors.New("rows error")
		rows := sqlmock.NewRows(columns).
			AddRow(artist1.ID, artist1.Name, artist1.AvatarURL, artist1.HeaderURL, artist1.Description, artist1.CreatedAt, artist1.UpdatedAt).
			RowError(0, expectedError)
		mock.ExpectQuery(query).WithArgs(ids).WillReturnRows(rows)

		artists, err := repo.GetByIDs(context.Background(), ids)

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Nil(t, artists)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArtistRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	mockArtist := newMockArtist()
	query := regexp.QuoteMeta(`SELECT artist_id, artist_name, avatar_url, header_url,  description, created_at, updated_at FROM artist ORDER BY artist_name LIMIT $1 OFFSET $2`)
	columns := []string{"artist_id", "artist_name", "avatar_url", "header_url", "description", "created_at", "updated_at"}

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows(columns).
			AddRow(mockArtist.ID, mockArtist.Name, mockArtist.AvatarURL, mockArtist.HeaderURL, mockArtist.Description, mockArtist.CreatedAt, mockArtist.UpdatedAt)
		mock.ExpectQuery(query).WithArgs(uint64(10), uint64(0)).WillReturnRows(rows)

		artists, err := repo.GetAll(context.Background(), 10, 0)

		assert.NoError(t, err)
		require.Len(t, artists, 1)
		assert.Equal(t, *mockArtist, artists[0])
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("success - empty result", func(t *testing.T) {
		rows := sqlmock.NewRows(columns)
		mock.ExpectQuery(query).WithArgs(uint64(10), uint64(0)).WillReturnRows(rows)

		artists, err := repo.GetAll(context.Background(), 10, 0)

		assert.NoError(t, err)
		assert.Len(t, artists, 0)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("db error", func(t *testing.T) {
		expectedError := sql.ErrConnDone
		mock.ExpectQuery(query).WithArgs(uint64(10), uint64(0)).WillReturnError(expectedError)

		_, err := repo.GetAll(context.Background(), 10, 0)

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
