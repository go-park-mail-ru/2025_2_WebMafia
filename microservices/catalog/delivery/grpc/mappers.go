package grpc

import (
	"spotify/microservices/catalog/dto"
	pb "spotify/proto/catalog"
)

func ArtistToProto(a *dto.Artist) *pb.Artist {
	if a == nil {
		return nil
	}
	return &pb.Artist{
		Id:          a.ID,
		Name:        a.Name,
		AvatarUrl:   a.AvatarURL,
		HeaderUrl:   a.HeaderURL,
		Description: &a.Description,
		PlayCount:   a.PlayCount,
	}
}

func ArtistsToProto(artists []dto.Artist) []*pb.Artist {
	res := make([]*pb.Artist, 0, len(artists))
	for i := range artists {
		res = append(res, ArtistToProto(&artists[i]))
	}
	return res
}

func ArtistForAlbumToProto(a *dto.ArtistForAlbum) *pb.ArtistForAlbum {
	if a == nil {
		return nil
	}
	return &pb.ArtistForAlbum{
		Id:        a.ID,
		Name:      a.Name,
		AvatarUrl: a.AvatarURL,
		HeaderUrl: a.HeaderURL,
	}
}

func ArtistForTrackToProto(a *dto.ArtistForTrack) *pb.ArtistForTrack {
	if a == nil {
		return nil
	}
	return &pb.ArtistForTrack{
		Id:        a.ID,
		Name:      a.Name,
		AvatarUrl: a.AvatarURL,
	}
}

func AlbumToProto(a *dto.Album) *pb.Album {
	if a == nil {
		return nil
	}
	artists := make([]*pb.ArtistForAlbum, 0, len(a.Artists))
	for i := range a.Artists {
		artists = append(artists, ArtistForAlbumToProto(&a.Artists[i]))
	}
	return &pb.Album{
		Id:          a.ID,
		Title:       a.Title,
		Type:        a.Type,
		AvatarUrl:   a.AvatarURL,
		Description: &a.Description,
		ReleaseDate: a.ReleaseDate,
		Artists:     artists,
	}
}

func AlbumsToProto(albums []dto.Album) []*pb.Album {
	res := make([]*pb.Album, 0, len(albums))
	for i := range albums {
		res = append(res, AlbumToProto(&albums[i]))
	}
	return res
}

func AlbumForTrackToProto(a *dto.AlbumForTrack) *pb.AlbumForTrack {
	if a == nil {
		return nil
	}
	artists := make([]*pb.ArtistForTrack, 0, len(a.Artists))
	for i := range a.Artists {
		artists = append(artists, ArtistForTrackToProto(&a.Artists[i]))
	}
	return &pb.AlbumForTrack{
		Id:          a.ID,
		Title:       a.Title,
		AvatarUrl:   a.AvatarURL,
		ReleaseDate: a.ReleaseDate,
		Artists:     artists,
	}
}

func GenreToProto(g *dto.Genre) *pb.Genre {
	if g == nil {
		return nil
	}
	return &pb.Genre{
		Id:   g.ID,
		Name: g.Name,
	}
}

func TrackToProto(t *dto.Track) *pb.Track {
	if t == nil {
		return nil
	}
	artists := make([]*pb.ArtistForTrack, 0, len(t.Artists))
	for i := range t.Artists {
		artists = append(artists, ArtistForTrackToProto(&t.Artists[i]))
	}

	genres := make([]*pb.Genre, 0, len(t.Genres))
	for i := range t.Genres {
		genres = append(genres, GenreToProto(&t.Genres[i]))
	}

	return &pb.Track{
		Id:        t.ID,
		Title:     t.Title,
		DurationS: int32(t.DurationS),
		FileUrl:   t.FileURL,
		PlayCount: t.PlayCount,
		Artists:   artists,
		Album:     AlbumForTrackToProto(&t.Album),
		Genres:    genres,
	}
}

func TracksToProto(tracks []dto.Track) []*pb.Track {
	res := make([]*pb.Track, 0, len(tracks))
	for i := range tracks {
		res = append(res, TrackToProto(&tracks[i]))
	}
	return res
}
