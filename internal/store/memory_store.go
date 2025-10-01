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
		{GenreID: 1, Name: "Hip-Hop", CreatedAt: time.Now()},
		{GenreID: 2, Name: "Rock", CreatedAt: time.Now()},
		{GenreID: 3, Name: "Indie", CreatedAt: time.Now()},
		{GenreID: 4, Name: "Post-Punk", CreatedAt: time.Now()},
	}
	ms.genres = append(ms.genres, genres...)

	artists := []model.Artist{
		{
			ArtistID:  1,
			Name:      "Tyler, the Creator",
			AvatarURL: "image1.jpg",
			CreatedAt: time.Now(),
		},
		{
			ArtistID:  2,
			Name:      "Mac Miller",
			AvatarURL: "image2.jpg",
			CreatedAt: time.Now(),
		},
		{
			ArtistID:  3,
			Name:      "Jpegmafia",
			AvatarURL: "image3.jpg",
			CreatedAt: time.Now(),
		},
		{
			ArtistID:  4,
			Name:      "The roots",
			AvatarURL: "image4.jpg",
			CreatedAt: time.Now(),
		},
		{
			ArtistID:  5,
			Name:      "Iggy Pop",
			AvatarURL: "image5.jpg",
			CreatedAt: time.Now(),
		},
		{
			ArtistID:  6,
			Name:      "Молчат дома",
			AvatarURL: "image6.jpg",
			CreatedAt: time.Now(),
		},
		{
			ArtistID:  7,
			Name:      "Arctic Monkeys",
			AvatarURL: "image7.jpg",
			CreatedAt: time.Now(),
		},
		{
			ArtistID:  8,
			Name:      "Kali Uchis",
			AvatarURL: "image8.jpg",
			CreatedAt: time.Now(),
		},
		{
			ArtistID:  9,
			Name:      "Playboi Carti",
			AvatarURL: "image9.jpg",
			CreatedAt: time.Now(),
		},
		{
			ArtistID:  10,
			Name:      "The Weekend",
			AvatarURL: "image10.jpg",
			CreatedAt: time.Now(),
		},
	}
	ms.artists = append(ms.artists, artists...)

	albums := []model.Album{
		{
			AlbumID:     1,
			Title:       "Flower boy",
			AvatarURL:   "image11.jpg",
			ArtistID:    1,
			ReleaseDate: time.Date(2017, 7, 21, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
		},
		{
			AlbumID:     2,
			Title:       "Lust for life",
			AvatarURL:   "image12.jpg",
			ArtistID:    5,
			ReleaseDate: time.Date(1977, 3, 29, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
		},
		{
			AlbumID:     3,
			Title:       "All my heroes are cornballs",
			AvatarURL:   "image13.jpg",
			ArtistID:    3,
			ReleaseDate: time.Date(2019, 9, 13, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
		},
		{
			AlbumID:     4,
			Title:       "Things fall apart",
			AvatarURL:   "image14.jpg",
			ArtistID:    4,
			ReleaseDate: time.Date(1999, 2, 23, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
		},
		{
			AlbumID:     5,
			Title:       "Flower boy (Deluxe)",
			AvatarURL:   "image15.jpg",
			ArtistID:    1,
			ReleaseDate: time.Date(2017, 12, 1, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
		},
		{
			AlbumID:     6,
			Title:       "Veteran",
			AvatarURL:   "image16.jpg",
			ArtistID:    3,
			ReleaseDate: time.Date(2018, 1, 19, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
		},
		{
			AlbumID:     7,
			Title:       "AM",
			AvatarURL:   "image17.jpg",
			ArtistID:    7,
			ReleaseDate: time.Date(2013, 9, 9, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
		},
		{
			AlbumID:     8,
			Title:       "Isolation",
			AvatarURL:   "image18.jpg",
			ArtistID:    8,
			ReleaseDate: time.Date(2018, 4, 6, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
		},
		{
			AlbumID:     9,
			Title:       "Die Lit",
			AvatarURL:   "image19.jpg",
			ArtistID:    9,
			ReleaseDate: time.Date(2018, 5, 11, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
		},
		{
			AlbumID:     10,
			Title:       "After Hours",
			AvatarURL:   "image20.jpg",
			ArtistID:    10,
			ReleaseDate: time.Date(2020, 3, 20, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
		},
	}
	ms.albums = append(ms.albums, albums...)

	tracks := []model.Track{
		{
			TrackID:    1,
			Title:      "See you again",
			DurationMs: 180000,
			FileURL:    "tyler_see_you_again.mp3",
			CreatedAt:  time.Now(),
			Album:      albums[0],
			Artists:    []model.Artist{artists[0], artists[7]},
			Genres:     []model.Genre{genres[0]},
		},
		{
			TrackID:    2,
			Title:      "Kenan vs Kel",
			DurationMs: 165000,
			FileURL:    "peggy_kenan_vs_kel.mp3",
			CreatedAt:  time.Now(),
			Album:      albums[5],
			Artists:    []model.Artist{artists[2]},
			Genres:     []model.Genre{genres[0]},
		},
		{
			TrackID:    3,
			Title:      "Rather Lie",
			DurationMs: 195000,
			FileURL:    "carti_rather_lie.mp3",
			CreatedAt:  time.Now(),
			Album:      albums[8],
			Artists:    []model.Artist{artists[8]},
			Genres:     []model.Genre{genres[0]},
		},
		{
			TrackID:    4,
			Title:      "See you again (Remix)",
			DurationMs: 210000,
			FileURL:    "tyler_see_you_again_remix.mp3",
			CreatedAt:  time.Now(),
			Album:      albums[4],
			Artists:    []model.Artist{artists[0], artists[7]},
			Genres:     []model.Genre{genres[0]},
		},
		{
			TrackID:    5,
			Title:      "1539 N. Calvert",
			DurationMs: 172000,
			FileURL:    "peggy_1539_calvert.mp3",
			CreatedAt:  time.Now(),
			Album:      albums[5],
			Artists:    []model.Artist{artists[2]},
			Genres:     []model.Genre{genres[0]},
		},
		{
			TrackID:    6,
			Title:      "Shoota",
			DurationMs: 188000,
			FileURL:    "carti_shoota.mp3",
			CreatedAt:  time.Now(),
			Album:      albums[8],
			Artists:    []model.Artist{artists[8], artists[9]},
			Genres:     []model.Genre{genres[0]},
		},
		{
			TrackID:    7,
			Title:      "Baby I'm Bleeding",
			DurationMs: 159000,
			FileURL:    "peggy_baby_bleeding.mp3",
			CreatedAt:  time.Now(),
			Album:      albums[5],
			Artists:    []model.Artist{artists[2]},
			Genres:     []model.Genre{genres[0]},
		},
		{
			TrackID:    8,
			Title:      "Long Time",
			DurationMs: 203000,
			FileURL:    "carti_long_time.mp3",
			CreatedAt:  time.Now(),
			Album:      albums[8],
			Artists:    []model.Artist{artists[8]},
			Genres:     []model.Genre{genres[0]},
		},
		{
			TrackID:    9,
			Title:      "Thug Tears",
			DurationMs: 176000,
			FileURL:    "peggy_thug_tears.mp3",
			CreatedAt:  time.Now(),
			Album:      albums[5],
			Artists:    []model.Artist{artists[2]},
			Genres:     []model.Genre{genres[0]},
		},
		{
			TrackID:    10,
			Title:      "FlatBed Freestyle",
			DurationMs: 192000,
			FileURL:    "carti_flatbed.mp3",
			CreatedAt:  time.Now(),
			Album:      albums[8],
			Artists:    []model.Artist{artists[8]},
			Genres:     []model.Genre{genres[0]},
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
