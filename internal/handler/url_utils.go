package handler

import (
	"fmt"
	"net/http"
	"spotify/internal/model"
)

func enrichArtistURLs(cfg StaticURLProvider, r *http.Request, artist *model.Artist) {
	base := cfg.GetStaticBaseURL(r)
	if artist.AvatarURL != "" {
		artist.AvatarURL = fmt.Sprintf("%s/img/%s", base, artist.AvatarURL)
	}
}

func enrichAlbumURLs(cfg StaticURLProvider, r *http.Request, album *model.Album) {
	base := cfg.GetStaticBaseURL(r)
	if album.AvatarURL != "" {
		album.AvatarURL = fmt.Sprintf("%s/img/%s", base, album.AvatarURL)
	}
	for i := range album.Tracks {
		enrichTrackURLs(cfg, r, &album.Tracks[i])
	}
}

func enrichTrackURLs(cfg StaticURLProvider, r *http.Request, track *model.Track) {
	base := cfg.GetStaticBaseURL(r)
	if track.FileURL != "" {
		track.FileURL = fmt.Sprintf("%s/img/%s", base, track.FileURL)
	}
	for i := range track.Artists {
		enrichArtistURLs(cfg, r, &track.Artists[i])
	}
	if track.Album.AlbumID != 0 {
		enrichAlbumURLs(cfg, r, &track.Album)
	}
}
