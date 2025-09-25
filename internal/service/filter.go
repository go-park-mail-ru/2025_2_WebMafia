package service

import (
	"spotify/internal/model"
	"strings"
)

func FilterTracksByArtist(tracks []*model.Track, artistID string) []*model.Track {
	if artistID == "" {
		return tracks
	}
	filtered := make([]*model.Track, 0)
	for _, t := range tracks {
		for _, a := range t.Artists {
			if a.ArtistID == artistID {
				filtered = append(filtered, t)
				break
			}
		}
	}
	return filtered
}

func FilterArtistsByName(artists []*model.Artist, name string) []*model.Artist {
	if name == "" {
		return artists
	}
	filtered := make([]*model.Artist, 0)
	for _, a := range artists {
		if strings.Contains(strings.ToLower(a.Name), strings.ToLower(name)) {
			filtered = append(filtered, a)
		}
	}
	return filtered
}

func FilterAlbumsByArtist(albums []*model.Album, artistID string) []*model.Album {
	if artistID == "" {
		return albums
	}
	filtered := make([]*model.Album, 0)
	for _, album := range albums {
		if album.ArtistID == artistID {
			filtered = append(filtered, album)
		}
	}
	return filtered
}
