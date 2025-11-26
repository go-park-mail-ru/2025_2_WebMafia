package grpc

import (
	"testing"

	"spotify/microservices/catalog/dto"

	"github.com/stretchr/testify/assert"
)

func TestMappers(t *testing.T) {
	t.Run("ArtistToProto", func(t *testing.T) {
		assert.Nil(t, ArtistToProto(nil))

		in := &dto.Artist{
			ID:          "1",
			Name:        "Name",
			AvatarURL:   "url",
			HeaderURL:   "header",
			Description: "desc",
			PlayCount:   100,
		}
		out := ArtistToProto(in)
		assert.Equal(t, in.ID, out.Id)
		assert.Equal(t, in.Name, out.Name)
		assert.Equal(t, in.PlayCount, out.PlayCount)
		assert.Equal(t, in.Description, *out.Description)
	})

	t.Run("ArtistsToProto", func(t *testing.T) {
		in := []dto.Artist{{ID: "1"}}
		out := ArtistsToProto(in)
		assert.Len(t, out, 1)
		assert.Equal(t, "1", out[0].Id)
	})

	t.Run("AlbumToProto", func(t *testing.T) {
		assert.Nil(t, AlbumToProto(nil))

		in := &dto.Album{
			ID:          "1",
			Title:       "Title",
			Type:        "Album",
			AvatarURL:   "url",
			Description: "desc",
			ReleaseDate: "2023-01-01",
			Artists:     []dto.ArtistForAlbum{{ID: "a1"}},
		}
		out := AlbumToProto(in)
		assert.Equal(t, in.ID, out.Id)
		assert.Equal(t, in.Title, out.Title)
		assert.Equal(t, in.Description, *out.Description)
		assert.Len(t, out.Artists, 1)
		assert.Equal(t, "a1", out.Artists[0].Id)
	})

	t.Run("AlbumsToProto", func(t *testing.T) {
		in := []dto.Album{{ID: "1"}}
		out := AlbumsToProto(in)
		assert.Len(t, out, 1)
	})

	t.Run("TrackToProto", func(t *testing.T) {
		assert.Nil(t, TrackToProto(nil))

		in := &dto.Track{
			ID:        "1",
			Title:     "T",
			DurationS: 100,
			FileURL:   "url",
			PlayCount: 50,
			Artists:   []dto.ArtistForTrack{{ID: "a1"}},
			Album:     dto.AlbumForTrack{ID: "al1"},
			Genres:    []dto.Genre{{ID: "g1"}},
		}
		out := TrackToProto(in)
		assert.Equal(t, in.ID, out.Id)
		assert.Equal(t, int32(in.DurationS), out.DurationS)
		assert.Equal(t, in.PlayCount, out.PlayCount)
		assert.Len(t, out.Artists, 1)
		assert.Equal(t, "al1", out.Album.Id)
		assert.Len(t, out.Genres, 1)
	})

	t.Run("TracksToProto", func(t *testing.T) {
		in := []dto.Track{{ID: "1"}}
		out := TracksToProto(in)
		assert.Len(t, out, 1)
	})

	t.Run("AlbumForTrackToProto nil", func(t *testing.T) {
		assert.Nil(t, AlbumForTrackToProto(nil))
	})

	t.Run("ArtistForAlbumToProto nil", func(t *testing.T) {
		assert.Nil(t, ArtistForAlbumToProto(nil))
	})

	t.Run("ArtistForTrackToProto nil", func(t *testing.T) {
		assert.Nil(t, ArtistForTrackToProto(nil))
	})

	t.Run("GenreToProto nil", func(t *testing.T) {
		assert.Nil(t, GenreToProto(nil))
	})
}
