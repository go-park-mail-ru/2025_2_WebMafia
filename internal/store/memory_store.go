package store

import (
	"context"
	"spotify/internal/model"
	"sync"
	"time"

	"github.com/google/uuid"
)

type MemoryStore struct {
	mu         *sync.RWMutex
	users      map[uuid.UUID]*model.User
	artists    []model.Artist
	tracks     []model.Track
	albums     []model.Album
	genres     []model.Genre
	nextArtist uint64
	nextAlbum  uint64
	nextTrack  uint64
	nextGenre  uint64
}

func NewMemoryStore() *MemoryStore {
	store := &MemoryStore{
		mu:      &sync.RWMutex{},
		users:   make(map[uuid.UUID]*model.User),
		artists: make([]model.Artist, 0),
		tracks:  make([]model.Track, 0),
		albums:  make([]model.Album, 0),
		genres:  make([]model.Genre, 0),
	}
	store.initMockData()
	return store
}

func (s *MemoryStore) CreateUser(ctx context.Context, user model.User) (*model.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, u := range s.users {
		if u.Login == user.Login || u.Email == user.Email {
			return nil, ErrUserAlreadyExists
		}
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()

	s.users[user.ID] = &user

	return &user, nil
}

func (s *MemoryStore) GetUserByLogin(ctx context.Context, login string) (*model.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, u := range s.users {
		if u.Login == login {
			user := *u
			return &user, nil
		}
	}
	return nil, ErrUserNotFound
}

func (s *MemoryStore) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, u := range s.users {
		if u.Email == email {
			user := *u
			return &user, nil
		}
	}
	return nil, ErrUserNotFound
}

func (s *MemoryStore) GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	u, ok := s.users[id]
	if !ok {
		return nil, ErrUserNotFound
	}
	user := *u
	return &user, nil
}

func (ms *MemoryStore) initMockData() {

	genres := []model.Genre{
		{GenreID: 1, Name: "genre1", CreatedAt: time.Now()},
		{GenreID: 2, Name: "genre2", CreatedAt: time.Now()},
		{GenreID: 3, Name: "genre3", CreatedAt: time.Now()},
		{GenreID: 4, Name: "genre4", CreatedAt: time.Now()},
	}
	ms.genres = append(ms.genres, genres...)

	artists := []model.Artist{
		{
			ArtistID:  1,
			Name:      "artist1",
			AvatarURL: "/artists/artist1",
			CreatedAt: time.Now(),
		},
		{
			ArtistID:  2,
			Name:      "artist2",
			AvatarURL: "/artists/artist2",
			CreatedAt: time.Now(),
		},
		{
			ArtistID:  3,
			Name:      "artist3",
			AvatarURL: "/artists/artist3",
			CreatedAt: time.Now(),
		},
		{
			ArtistID:  4,
			Name:      "artist4",
			AvatarURL: "/artists/artist4",
			CreatedAt: time.Now(),
		},
	}
	ms.artists = append(ms.artists, artists...)

	albums := []model.Album{
		{
			AlbumID:     1,
			Title:       "Some Album",
			AvatarURL:   "/albums/some_album",
			ArtistID:    1,
			ReleaseDate: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
		},
		{
			AlbumID:     2,
			Title:       "Album 2",
			AvatarURL:   "/albums/album2",
			ArtistID:    2,
			ReleaseDate: time.Date(2024, 10, 20, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
		},
		{
			AlbumID:     3,
			Title:       "Album 3",
			AvatarURL:   "/albums/album3",
			ArtistID:    3,
			ReleaseDate: time.Date(2024, 8, 10, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
		},
		{
			AlbumID:     4,
			Title:       "Album 4",
			AvatarURL:   "/albums/album4",
			ArtistID:    4,
			ReleaseDate: time.Date(2022, 5, 2, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
		},
	}
	ms.albums = append(ms.albums, albums...)

	tracks := []model.Track{
		{
			TrackID:    1,
			Title:      "Some Track",
			DurationMs: 0,
			FileURL:    "tracks/some_track",
			CreatedAt:  time.Now(),
			Album:      albums[0],
			Artists:    []model.Artist{artists[0]},
			Genres:     []model.Genre{genres[0]},
		},
		{
			TrackID:    2,
			Title:      "Track 2",
			DurationMs: 0,
			FileURL:    "tracks/track2",
			CreatedAt:  time.Now(),
			Album:      albums[0],
			Artists:    []model.Artist{artists[0]},
			Genres:     []model.Genre{genres[0]},
		},
		{
			TrackID:    3,
			Title:      "Track 3",
			DurationMs: 0,
			FileURL:    "tracks/track3",
			CreatedAt:  time.Now(),
			Album:      albums[1],
			Artists:    []model.Artist{artists[1]},
			Genres:     []model.Genre{genres[1]},
		},
		{
			TrackID:    4,
			Title:      "Track 4",
			DurationMs: 0,
			FileURL:    "tracks/track4",
			CreatedAt:  time.Now(),
			Album:      albums[2],
			Artists:    []model.Artist{artists[2]},
			Genres:     []model.Genre{genres[2]},
		},
		{
			TrackID:    5,
			Title:      "Track 5",
			DurationMs: 0,
			FileURL:    "tracks/track5",
			CreatedAt:  time.Now(),
			Album:      albums[3],
			Artists:    []model.Artist{artists[3]},
			Genres:     []model.Genre{genres[3]},
		},
	}
	ms.tracks = append(ms.tracks, tracks...)

	ms.nextArtist = uint64(len(artists) + 1)
	ms.nextAlbum = uint64(len(albums) + 1)
	ms.nextTrack = uint64(len(tracks) + 1)
	ms.nextGenre = uint64(len(genres) + 1)
}

func (ms *MemoryStore) GetAllTracks() ([]model.Track, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	return ms.tracks, nil
}

func (ms *MemoryStore) GetAllArtists() ([]model.Artist, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	return ms.artists, nil
}

func (ms *MemoryStore) GetAllAlbums() ([]model.Album, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	return ms.albums, nil
}
