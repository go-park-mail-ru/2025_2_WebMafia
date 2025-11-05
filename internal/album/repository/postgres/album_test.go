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

func newMockAlbum() *model.Album {
	return &model.Album{
		ID:          uuid.New(),
		Title:       "Test Album",
		Type:        "Альбом",
		AvatarURL:   "http://example.com/album.jpg",
		ArtistID:    uuid.New(),
		Description: sql.NullString{String: "A test album", Valid: true},
		ReleaseDate: time.Now().Truncate(24 * time.Hour),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func TestAlbumRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()
	repo := New(db)

	query := regexp.QuoteMeta(`SELECT album_id, title, type, avatar_url, artist_id, description, release_date, created_at, updated_at FROM album WHERE album_id = $1`)
	columns := []string{"album_id", "title", "type", "avatar_url", "artist_id", "description", "release_date", "created_at", "updated_at"}
	mockAlbum := newMockAlbum()

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows(columns).
			AddRow(mockAlbum.ID, mockAlbum.Title, mockAlbum.Type, mockAlbum.AvatarURL, mockAlbum.ArtistID, mockAlbum.Description, mockAlbum.ReleaseDate, mockAlbum.CreatedAt, mockAlbum.UpdatedAt)
		mock.ExpectQuery(query).WithArgs(mockAlbum.ID).WillReturnRows(rows)

		album, err := repo.GetByID(context.Background(), mockAlbum.ID)
		assert.NoError(t, err)
		assert.Equal(t, mockAlbum, album)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("not found", func(t *testing.T) {
		mock.ExpectQuery(query).WithArgs(mockAlbum.ID).WillReturnError(sql.ErrNoRows)

		_, err := repo.GetByID(context.Background(), mockAlbum.ID)
		assert.ErrorIs(t, err, ErrNotFound)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestAlbumRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := New(db)

	query := regexp.QuoteMeta(`SELECT album_id, title, type, avatar_url, artist_id, description, release_date, created_at, updated_at FROM album ORDER BY release_date DESC LIMIT $1 OFFSET $2`)
	columns := []string{"album_id", "title", "type", "avatar_url", "artist_id", "description", "release_date", "created_at", "updated_at"}
	mockAlbum := newMockAlbum()

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows(columns).
			AddRow(mockAlbum.ID, mockAlbum.Title, mockAlbum.Type, mockAlbum.AvatarURL, mockAlbum.ArtistID, mockAlbum.Description, mockAlbum.ReleaseDate, mockAlbum.CreatedAt, mockAlbum.UpdatedAt)
		mock.ExpectQuery(query).WithArgs(uint64(10), uint64(0)).WillReturnRows(rows)

		albums, err := repo.GetAll(context.Background(), 10, 0)
		assert.NoError(t, err)
		require.Len(t, albums, 1)
		assert.Equal(t, *mockAlbum, albums[0])
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("db error", func(t *testing.T) {
		mock.ExpectQuery(query).WithArgs(uint64(10), uint64(0)).WillReturnError(errors.New("db error"))
		_, err := repo.GetAll(context.Background(), 10, 0)
		assert.Error(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestAlbumRepository_GetByArtistID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()
	repo := New(db)

	query := regexp.QuoteMeta(`SELECT album_id, title, type, avatar_url, artist_id, description, release_date, created_at, updated_at FROM album WHERE artist_id = $1 ORDER BY release_date DESC LIMIT $2 OFFSET $3`)
	columns := []string{"album_id", "title", "type", "avatar_url", "artist_id", "description", "release_date", "created_at", "updated_at"}
	mockAlbum := newMockAlbum()

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows(columns).
			AddRow(mockAlbum.ID, mockAlbum.Title, mockAlbum.Type, mockAlbum.AvatarURL, mockAlbum.ArtistID, mockAlbum.Description, mockAlbum.ReleaseDate, mockAlbum.CreatedAt, mockAlbum.UpdatedAt)
		mock.ExpectQuery(query).WithArgs(mockAlbum.ArtistID, uint64(10), uint64(0)).WillReturnRows(rows)

		albums, err := repo.GetByArtistID(context.Background(), mockAlbum.ArtistID, 10, 0)
		assert.NoError(t, err)
		require.Len(t, albums, 1)
		assert.Equal(t, *mockAlbum, albums[0])
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("db error", func(t *testing.T) {
		mock.ExpectQuery(query).WithArgs(mockAlbum.ArtistID, uint64(10), uint64(0)).WillReturnError(errors.New("db error"))
		_, err := repo.GetByArtistID(context.Background(), mockAlbum.ArtistID, 10, 0)
		assert.Error(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestAlbumRepository_GetByIDs(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(pgxUUIDValueConverter{}))
	require.NoError(t, err)
	defer db.Close()
	repo := New(db)

	query := regexp.QuoteMeta(`SELECT album_id, title, type, avatar_url, artist_id, description, release_date, created_at, updated_at FROM album WHERE album_id = ANY($1)`)
	columns := []string{"album_id", "title", "type", "avatar_url", "artist_id", "description", "release_date", "created_at", "updated_at"}
	album1 := newMockAlbum()
	album2 := newMockAlbum()
	ids := []uuid.UUID{album1.ID, album2.ID}

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows(columns).
			AddRow(album1.ID, album1.Title, album1.Type, album1.AvatarURL, album1.ArtistID, album1.Description, album1.ReleaseDate, album1.CreatedAt, album1.UpdatedAt).
			AddRow(album2.ID, album2.Title, album2.Type, album2.AvatarURL, album2.ArtistID, album2.Description, album2.ReleaseDate, album2.CreatedAt, album2.UpdatedAt)
		mock.ExpectQuery(query).WithArgs(ids).WillReturnRows(rows)

		albums, err := repo.GetByIDs(context.Background(), ids)
		assert.NoError(t, err)
		require.Len(t, albums, 2)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("empty ids slice", func(t *testing.T) {
		albums, err := repo.GetByIDs(context.Background(), []uuid.UUID{})
		assert.NoError(t, err)
		assert.Nil(t, albums)
	})

	t.Run("db error", func(t *testing.T) {
		mock.ExpectQuery(query).WithArgs(ids).WillReturnError(errors.New("db error"))
		_, err := repo.GetByIDs(context.Background(), ids)
		assert.Error(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
