package service

import (
	"context"
	"errors"
	"fmt"
	"spotify/internal/track/dto"
	"spotify/internal/track/model"
	"spotify/internal/track/repository/postgres"

	"github.com/google/uuid"
)

func mapModelsToDTOs(tracks []model.Track, albums []model.Album, artistsSlice [][]model.Artist, genresSlice [][]model.Genre) []dto.Track {
	trackDTOs := make([]dto.Track, len(tracks))

	for i := range tracks {

		artistDTOs := make([]dto.Artist, len(artistsSlice[i]))
		for j, artist := range artistsSlice[i] {
			artistDTOs[j] = dto.Artist{
				ID:        artist.ID.String(),
				Name:      artist.Name,
				AvatarURL: artist.AvatarURL,
			}
		}

		genreDTOs := make([]dto.Genre, len(genresSlice[i]))
		for j, genre := range genresSlice[i] {
			genreDTOs[j] = dto.Genre{
				ID:   genre.ID.String(),
				Name: genre.Name,
			}
		}

		trackDTOs[i] = dto.Track{
			ID:         tracks[i].ID.String(),
			Title:      tracks[i].Title,
			DurationMs: tracks[i].DurationMs,
			FileURL:    tracks[i].FileURL,
			Album: dto.Album{
				ID:          albums[i].ID.String(),
				Title:       albums[i].Title,
				AvatarURL:   albums[i].AvatarURL,
				ReleaseDate: albums[i].ReleaseDate.Format("2006-01-02"),
				Artists:     artistDTOs,
			},
			Artists: artistDTOs,
			Genres:  genreDTOs,
		}
	}

	return trackDTOs
}

func (s *Service) GetTrackByID(ctx context.Context, id uuid.UUID) (*dto.Track, error) {
	track, album, artists, genres, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, postgres.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("service.GetTrackByID: %w", err)
	}
	resultDTOs := mapModelsToDTOs([]model.Track{*track}, []model.Album{*album}, [][]model.Artist{artists}, [][]model.Genre{genres})
	return &resultDTOs[0], nil
}

func (s *Service) GetAllTracks(ctx context.Context) ([]dto.Track, error) {
	tracks, albums, artists, genres, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("service.GetAllTracks: %w", err)
	}
	return mapModelsToDTOs(tracks, albums, artists, genres), nil
}

func (s *Service) GetTracksByArtistID(ctx context.Context, artistID uuid.UUID) ([]dto.Track, error) {
	tracks, albums, artists, genres, err := s.repo.GetByArtistID(ctx, artistID)
	if err != nil {
		return nil, fmt.Errorf("service.GetTracksByArtistID: %w", err)
	}
	return mapModelsToDTOs(tracks, albums, artists, genres), nil
}

func (s *Service) GetTracksByAlbumID(ctx context.Context, albumID uuid.UUID) ([]dto.Track, error) {
	tracks, albums, artists, genres, err := s.repo.GetByAlbumID(ctx, albumID)
	if err != nil {
		return nil, fmt.Errorf("service.GetTracksByAlbumID: %w", err)
	}
	return mapModelsToDTOs(tracks, albums, artists, genres), nil
}

func (s *Service) GetTracksByGenreID(ctx context.Context, genreID uuid.UUID) ([]dto.Track, error) {
	tracks, albums, artists, genres, err := s.repo.GetByGenreID(ctx, genreID)
	if err != nil {
		return nil, fmt.Errorf("service.GetTracksByGenreID: %w", err)
	}
	return mapModelsToDTOs(tracks, albums, artists, genres), nil
}
