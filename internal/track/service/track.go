package service

import (
	"context"
	"fmt"
	"spotify/internal/model"
	"spotify/internal/track/dto"
	"sync"

	"github.com/google/uuid"
)

func (s *Service) GetTrackByID(ctx context.Context, id uuid.UUID) (*dto.Track, error) {
	const op = "service.GetTrackByID"
	trackModel, err := s.trackRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get a track by id: %w", op, mapError(err))
	}

	enrichedTracks, err := s.enrichTracks(ctx, []model.Track{*trackModel})
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to enrich track: %w", op, mapError(err))
	}
	if len(enrichedTracks) == 0 {
		return nil, fmt.Errorf("[%s]: track data is inconsistent: %w", op, mapError(err))
	}

	return &enrichedTracks[0], nil
}

func (s *Service) GetAllTracks(ctx context.Context, limit, offset uint64) ([]dto.Track, error) {
	const op = "service.GetAllTracks"
	trackModels, err := s.trackRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get tracks: %w", op, mapError(err))
	}
	return s.enrichTracks(ctx, trackModels)
}

func (s *Service) GetTracksByArtistID(ctx context.Context, artistID uuid.UUID, limit, offset uint64) ([]dto.Track, error) {
	const op = "service.GetTracksByArtistID"
	trackModels, err := s.trackRepo.GetByArtistID(ctx, artistID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get tracks: %w", op, mapError(err))
	}
	return s.enrichTracks(ctx, trackModels)
}

func (s *Service) GetTracksByAlbumID(ctx context.Context, albumID uuid.UUID, limit, offset uint64) ([]dto.Track, error) {
	const op = "service.GetTracksByAlbumID"
	trackModels, err := s.trackRepo.GetByAlbumID(ctx, albumID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get tracks: %w", op, mapError(err))
	}
	return s.enrichTracks(ctx, trackModels)
}

func (s *Service) GetTracksByGenreID(ctx context.Context, genreID uuid.UUID, limit, offset uint64) ([]dto.Track, error) {
	const op = "service.GetTracksByGenreID"
	trackModels, err := s.trackRepo.GetByGenreID(ctx, genreID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get tracks: %w", op, mapError(err))
	}
	return s.enrichTracks(ctx, trackModels)
}

func (s *Service) enrichTracks(ctx context.Context, tracks []model.Track) ([]dto.Track, error) {
	const op = "service.enrichTracks"
	if len(tracks) == 0 {
		return []dto.Track{}, nil
	}

	trackIDs := make([]uuid.UUID, len(tracks))
	for i, track := range tracks {
		trackIDs[i] = track.ID
	}

	var wg sync.WaitGroup
	var albumsMap map[uuid.UUID]dto.Album
	var artistsMap map[uuid.UUID][]dto.Artist
	var genresMap map[uuid.UUID][]dto.Genre
	var errAlbums, errArtists, errGenres error

	wg.Add(3)

	go func() {
		defer wg.Done()
		albumsMap, errAlbums = s.getAlbumsForTracks(ctx, trackIDs)
	}()

	go func() {
		defer wg.Done()
		artistsMap, errArtists = s.getArtistsForTracks(ctx, trackIDs)
	}()

	go func() {
		defer wg.Done()
		genresMap, errGenres = s.getGenresForTracks(ctx, trackIDs)
	}()

	wg.Wait()

	if errAlbums != nil {
		return nil, fmt.Errorf("[%s]: failed to get albums: %w", op, errAlbums)
	}
	if errArtists != nil {
		return nil, fmt.Errorf("[%s]: failed to get artists: %w", op, errArtists)
	}
	if errGenres != nil {
		return nil, fmt.Errorf("[%s]: failed to get genres: %w", op, errGenres)
	}

	enrichedTracks := make([]dto.Track, 0, len(tracks))
	for _, track := range tracks {
		album, albumOk := albumsMap[track.ID]
		artists, artistsOk := artistsMap[track.ID]
		genres := genresMap[track.ID]

		if !albumOk || !artistsOk {
			continue
		}

		enrichedTracks = append(enrichedTracks, dto.Track{
			ID:         track.ID.String(),
			Title:      track.Title,
			DurationMs: track.DurationMs,
			FileURL:    track.FileURL,
			Album:      album,
			Artists:    artists,
			Genres:     genres,
		})
	}

	return enrichedTracks, nil
}

func (s *Service) getAlbumsForTracks(ctx context.Context, trackIDs []uuid.UUID) (map[uuid.UUID]dto.Album, error) {
	const op = "service.getAlbumsForTracks"
	trackAlbumIDMap, err := s.trackRepo.GetAlbumIDsForTracks(ctx, trackIDs)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get album IDs from repo: %w", op, err)
	}

	albumIDs := make([]uuid.UUID, 0, len(trackAlbumIDMap))
	uniqueAlbumIDs := make(map[uuid.UUID]bool)
	for _, albumID := range trackAlbumIDMap {
		if !uniqueAlbumIDs[albumID] {
			uniqueAlbumIDs[albumID] = true
			albumIDs = append(albumIDs, albumID)
		}
	}
	if len(albumIDs) == 0 {
		return make(map[uuid.UUID]dto.Album), nil
	}

	albumServiceDTOs, err := s.albumService.GetAlbumsByIDs(ctx, albumIDs)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get albums from service: %w", op, err)
	}

	albumsMapByID := make(map[uuid.UUID]dto.Album, len(albumServiceDTOs))
	for _, albumServiceDTO := range albumServiceDTOs {
		albumUUID, err := uuid.Parse(albumServiceDTO.ID)
		if err != nil {
			continue
		}
		albumsMapByID[albumUUID] = dto.Album{
			ID:        albumServiceDTO.ID,
			Title:     albumServiceDTO.Title,
			AvatarURL: albumServiceDTO.AvatarURL,
		}
	}

	result := make(map[uuid.UUID]dto.Album, len(trackIDs))
	for trackID, albumID := range trackAlbumIDMap {
		if album, ok := albumsMapByID[albumID]; ok {
			result[trackID] = album
		}
	}
	return result, nil
}

func (s *Service) getArtistsForTracks(ctx context.Context, trackIDs []uuid.UUID) (map[uuid.UUID][]dto.Artist, error) {
	const op = "service.getArtistsForTracks"
	trackArtistIDsMap, err := s.trackRepo.GetArtistIDsForTracks(ctx, trackIDs)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get artist IDs from repo: %w", op, err)
	}

	uniqueArtistIDs := make(map[uuid.UUID]bool)
	for _, artistIDs := range trackArtistIDsMap {
		for _, artistID := range artistIDs {
			uniqueArtistIDs[artistID] = true
		}
	}
	allArtistIDs := make([]uuid.UUID, 0, len(uniqueArtistIDs))
	for id := range uniqueArtistIDs {
		allArtistIDs = append(allArtistIDs, id)
	}
	if len(allArtistIDs) == 0 {
		return make(map[uuid.UUID][]dto.Artist), nil
	}

	artistServiceDTOs, err := s.artistService.GetArtistsByIDs(ctx, allArtistIDs)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get artists from service: %w", op, err)
	}

	artistsMapByID := make(map[uuid.UUID]dto.Artist, len(artistServiceDTOs))
	for _, artistServiceDTO := range artistServiceDTOs {
		artistUUID, err := uuid.Parse(artistServiceDTO.ID)
		if err != nil {
			continue
		}
		artistsMapByID[artistUUID] = dto.Artist{
			ID:        artistServiceDTO.ID,
			Name:      artistServiceDTO.Name,
			AvatarURL: artistServiceDTO.AvatarURL,
		}
	}

	result := make(map[uuid.UUID][]dto.Artist, len(trackIDs))
	for trackID, artistIDs := range trackArtistIDsMap {
		for _, artistID := range artistIDs {
			if artist, ok := artistsMapByID[artistID]; ok {
				result[trackID] = append(result[trackID], artist)
			}
		}
	}
	return result, nil
}

func (s *Service) getGenresForTracks(ctx context.Context, trackIDs []uuid.UUID) (map[uuid.UUID][]dto.Genre, error) {
	const op = "service.getGenresForTracks"
	trackGenresMap, err := s.trackRepo.GetGenresForTracks(ctx, trackIDs)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get genres from repo: %w", op, err)
	}

	result := make(map[uuid.UUID][]dto.Genre, len(trackGenresMap))
	for trackID, genres := range trackGenresMap {
		dtoGenres := make([]dto.Genre, len(genres))
		for i, genre := range genres {
			dtoGenres[i] = dto.Genre{
				ID:   genre.ID.String(),
				Name: genre.Name,
			}
		}
		result[trackID] = dtoGenres
	}
	return result, nil
}
