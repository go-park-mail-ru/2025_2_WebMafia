package mock_store

import "spotify/internal/model"

func (ms *MockStore) GetAllTracks() ([]*model.Track, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	return ms.tracks, nil
}

func (ms *MockStore) GetAllArtists() ([]*model.Artist, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	return ms.artists, nil
}

func (ms *MockStore) GetAllAlbums() ([]*model.Album, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	return ms.albums, nil
}
