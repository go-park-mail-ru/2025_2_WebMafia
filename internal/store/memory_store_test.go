package store

import (
	"context"
	model2 "spotify/internal/user/model"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemoryStore_CreateUser_Success(t *testing.T) {
	store := NewMemoryStore()

	user := model2.User{
		Login:        "user_login",
		Email:        "new_user@test.com",
		PasswordHash: "some_password",
	}

	created, err := store.CreateUser(context.Background(), user)
	require.NoError(t, err)

	assert.NotEqual(t, uuid.Nil, created.ID)
	assert.Equal(t, user.Login, created.Login)
	assert.Equal(t, user.Email, created.Email)
	assert.Equal(t, user.PasswordHash, created.PasswordHash)
	assert.WithinDuration(t, time.Now(), created.CreatedAt, time.Second)
	assert.WithinDuration(t, time.Now(), created.UpdatedAt, time.Second)
}

func TestMemoryStore_CreateUser_DuplicateLogin(t *testing.T) {
	store := NewMemoryStore()

	user1 := model2.User{
		Login:        "user_login",
		Email:        "new_user1@test.com",
		PasswordHash: "some_password1",
	}

	user2 := model2.User{
		Login:        "user_login",
		Email:        "new_user2@test.com",
		PasswordHash: "some_password2",
	}

	_, err := store.CreateUser(context.Background(), user1)
	require.NoError(t, err)

	_, err = store.CreateUser(context.Background(), user2)
	assert.ErrorIs(t, err, ErrUserAlreadyExists)
}

func TestMemoryStore_CreateUser_DuplicateEmail(t *testing.T) {
	store := NewMemoryStore()

	user1 := model2.User{
		Login:        "user_login1",
		Email:        "new_user1@test.com",
		PasswordHash: "some_password1",
	}

	user2 := model2.User{
		Login:        "user_login",
		Email:        "new_user1@test.com",
		PasswordHash: "some_password1",
	}

	_, err := store.CreateUser(context.Background(), user1)
	require.NoError(t, err)

	_, err = store.CreateUser(context.Background(), user2)
	assert.ErrorIs(t, err, ErrUserAlreadyExists)
}

func TestMemoryStore_GetUserByLogin_Success(t *testing.T) {
	store := NewMemoryStore()

	user := model2.User{
		Login:        "user_login",
		Email:        "new_user@test.com",
		PasswordHash: "some_password",
	}

	created, err := store.CreateUser(context.Background(), user)
	require.NoError(t, err)

	found, err := store.GetUserByLogin(context.Background(), "user_login")
	require.NoError(t, err)

	assert.Equal(t, created.ID, found.ID)
	assert.Equal(t, user.Login, found.Login)
	assert.Equal(t, user.Email, found.Email)
	assert.Equal(t, user.PasswordHash, found.PasswordHash)
}

func TestMemoryStore_GetUserByLogin_NotFound(t *testing.T) {
	store := NewMemoryStore()

	_, err := store.GetUserByLogin(context.Background(), "noexist")
	assert.ErrorIs(t, err, ErrUserNotFound)
}

func TestMemoryStore_GetUserByEmail_Success(t *testing.T) {
	store := NewMemoryStore()

	user := model2.User{
		Login:        "user_login",
		Email:        "new_user@test.com",
		PasswordHash: "some_password",
	}

	created, err := store.CreateUser(context.Background(), user)
	require.NoError(t, err)

	found, err := store.GetUserByEmail(context.Background(), "new_user@test.com")
	require.NoError(t, err)

	assert.Equal(t, created.ID, found.ID)
	assert.Equal(t, user.Login, found.Login)
	assert.Equal(t, user.Email, found.Email)
}

func TestMemoryStore_GetUserByEmail_NotFound(t *testing.T) {
	store := NewMemoryStore()

	_, err := store.GetUserByEmail(context.Background(), "noexist@test.com")
	assert.ErrorIs(t, err, ErrUserNotFound)
}

func TestMemoryStore_GetUserByID(t *testing.T) {
	store := NewMemoryStore()

	user := model2.User{
		Login:        "user_login",
		Email:        "new_user@test.com",
		PasswordHash: "some_password",
	}

	created, err := store.CreateUser(context.Background(), user)
	require.NoError(t, err)

	found, err := store.GetUserByID(context.Background(), created.ID)
	require.NoError(t, err)

	assert.Equal(t, created.ID, found.ID)
	assert.Equal(t, user.Login, found.Login)
	assert.Equal(t, user.Email, found.Email)
}

func TestMemoryStore_GetUserByID_NotFound(t *testing.T) {
	store := NewMemoryStore()

	randomID := uuid.New()
	_, err := store.GetUserByID(context.Background(), randomID)
	assert.ErrorIs(t, err, ErrUserNotFound)
}

func TestMemoryStore_GetAllTracks(t *testing.T) {
	store := NewMemoryStore()

	tracks, err := store.GetAllTracks()
	require.NoError(t, err)

	assert.Greater(t, len(tracks), 0)

	for _, track := range tracks {
		assert.NotEmpty(t, track.Title)
		assert.Greater(t, track.DurationMs, 0)
		assert.NotEmpty(t, track.FileURL)
		assert.NotZero(t, track.TrackID)
		assert.NotNil(t, track.Album)
		assert.NotEmpty(t, track.Artists)
		assert.NotEmpty(t, track.Genres)
	}
}

func TestMemoryStore_GetAllArtists(t *testing.T) {
	store := NewMemoryStore()

	artists, err := store.GetAllArtists()
	require.NoError(t, err)

	assert.Greater(t, len(artists), 0)

	for _, artist := range artists {
		assert.NotEmpty(t, artist.Name)
		assert.NotZero(t, artist.ArtistID)
		assert.NotEmpty(t, artist.AvatarURL)
		assert.NotZero(t, artist.CreatedAt)
	}
}

func TestMemoryStore_GetAllAlbums(t *testing.T) {
	store := NewMemoryStore()

	albums, err := store.GetAllAlbums()
	require.NoError(t, err)

	assert.Greater(t, len(albums), 0)

	for _, album := range albums {
		assert.NotEmpty(t, album.Title)
		assert.NotZero(t, album.AlbumID)
		assert.NotEmpty(t, album.AvatarURL)
		assert.NotZero(t, album.ArtistID)
		assert.NotZero(t, album.ReleaseDate)
		assert.NotZero(t, album.CreatedAt)
	}
}
