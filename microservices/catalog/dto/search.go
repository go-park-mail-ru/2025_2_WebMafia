package dto

import "spotify/internal/model"

type ArtistSearchResult struct {
	Artist model.Artist
	Rank   float32
}

type AlbumSearchResult struct {
	Album model.Album
	Rank  float32
}

type TrackSearchResult struct {
	Track model.Track
	Rank  float32
}
