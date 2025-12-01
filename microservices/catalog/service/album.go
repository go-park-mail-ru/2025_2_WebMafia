package service

import (
	"context"
	"fmt"

	"spotify/microservices/catalog/dto"

	"github.com/google/uuid"
)

const (
	dateFormat = "2006-01-02"
)

func (s *Service) GetAlbumByID(ctx context.Context, id uuid.UUID) (*dto.Album, error) {
	const op = "service.GetAlbumByID"
	albumModel, err := s.repo.GetAlbumByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get an album: %w", op, mapError(err))
	}

	artist, err := s.GetArtistByID(ctx, albumModel.ArtistID)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get an artist: %w", op, mapError(err))
	}

	albumDTO := &dto.Album{
		ID:          albumModel.ID.String(),
		Title:       albumModel.Title,
		Type:        albumModel.Type,
		AvatarURL:   albumModel.AvatarURL,
		ReleaseDate: albumModel.ReleaseDate.Format(dateFormat),
		Artists: []dto.ArtistForAlbum{
			{
				ID:        artist.ID,
				Name:      artist.Name,
				AvatarURL: artist.AvatarURL,
			},
		},
	}
	if albumModel.Description.Valid {
		albumDTO.Description = albumModel.Description.String
	}

	return albumDTO, nil
}

func (s *Service) GetAllAlbums(ctx context.Context, limit, offset uint64) ([]dto.Album, error) {
	const op = "service.GetAllAlbums"
	albumModels, err := s.repo.GetAllAlbums(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get all albums: %w", op, mapError(err))
	}
	if len(albumModels) == 0 {
		return []dto.Album{}, nil
	}

	artistIDs := make([]uuid.UUID, 0, len(albumModels))
	uniqueArtistIDs := make(map[uuid.UUID]bool)
	for _, album := range albumModels {
		if !uniqueArtistIDs[album.ArtistID] {
			uniqueArtistIDs[album.ArtistID] = true
			artistIDs = append(artistIDs, album.ArtistID)
		}
	}

	artists, err := s.GetArtistsByIDs(ctx, artistIDs)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get an artist: %w", op, mapError(err))
	}

	artistsMap := make(map[string]dto.Artist, len(artists))
	for _, artist := range artists {
		artistsMap[artist.ID] = artist
	}

	albumDTOs := make([]dto.Album, 0, len(albumModels))
	for _, albumModel := range albumModels {
		artist, ok := artistsMap[albumModel.ArtistID.String()]
		if !ok {
			continue
		}

		album := dto.Album{
			ID:          albumModel.ID.String(),
			Title:       albumModel.Title,
			Type:        albumModel.Type,
			AvatarURL:   albumModel.AvatarURL,
			ReleaseDate: albumModel.ReleaseDate.Format(dateFormat),
			Artists: []dto.ArtistForAlbum{
				{
					ID:        artist.ID,
					Name:      artist.Name,
					AvatarURL: artist.AvatarURL,
				},
			},
		}
		if albumModel.Description.Valid {
			album.Description = albumModel.Description.String
		}
		albumDTOs = append(albumDTOs, album)
	}

	return albumDTOs, nil
}

func (s *Service) GetAlbumsByIDs(ctx context.Context, ids []uuid.UUID) ([]dto.Album, error) {
	const op = "service.GetAlbumsByIDs"
	albumModels, err := s.repo.GetAlbumsByIDs(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get albums by ids: %w", op, mapError(err))
	}
	if len(albumModels) == 0 {
		return []dto.Album{}, nil
	}

	artistIDs := make([]uuid.UUID, 0, len(albumModels))
	uniqueArtistIDs := make(map[uuid.UUID]bool)
	for _, album := range albumModels {
		if !uniqueArtistIDs[album.ArtistID] {
			uniqueArtistIDs[album.ArtistID] = true
			artistIDs = append(artistIDs, album.ArtistID)
		}
	}

	artists, err := s.GetArtistsByIDs(ctx, artistIDs)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get an artist: %w", op, mapError(err))
	}

	artistsMap := make(map[string]dto.Artist, len(artists))
	for _, artist := range artists {
		artistsMap[artist.ID] = artist
	}

	albumDTOs := make([]dto.Album, 0, len(albumModels))
	for _, albumModel := range albumModels {
		artist, ok := artistsMap[albumModel.ArtistID.String()]
		if !ok {
			continue
		}
		album := dto.Album{
			ID:          albumModel.ID.String(),
			Title:       albumModel.Title,
			Type:        albumModel.Type,
			AvatarURL:   albumModel.AvatarURL,
			ReleaseDate: albumModel.ReleaseDate.Format(dateFormat),
			Artists: []dto.ArtistForAlbum{
				{
					ID:        artist.ID,
					Name:      artist.Name,
					AvatarURL: artist.AvatarURL,
				},
			},
		}
		if albumModel.Description.Valid {
			album.Description = albumModel.Description.String
		}
		albumDTOs = append(albumDTOs, album)
	}

	return albumDTOs, nil
}

func (s *Service) GetAlbumsByArtistID(ctx context.Context, artistID uuid.UUID, limit, offset uint64) ([]dto.Album, error) {
	const op = "service.GetAlbumsByArtistID"
	albumModels, err := s.repo.GetAlbumsByArtistID(ctx, artistID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get albums: %w", op, mapError(err))
	}
	if len(albumModels) == 0 {
		return []dto.Album{}, nil
	}

	artist, err := s.GetArtistByID(ctx, artistID)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get artist: %w", op, mapError(err))
	}

	albumDTOs := make([]dto.Album, 0, len(albumModels))
	for _, albumModel := range albumModels {
		album := dto.Album{
			ID:          albumModel.ID.String(),
			Title:       albumModel.Title,
			Type:        albumModel.Type,
			AvatarURL:   albumModel.AvatarURL,
			ReleaseDate: albumModel.ReleaseDate.Format(dateFormat),
			Artists: []dto.ArtistForAlbum{
				{
					ID:        artist.ID,
					Name:      artist.Name,
					AvatarURL: artist.AvatarURL,
				},
			},
		}
		if albumModel.Description.Valid {
			album.Description = albumModel.Description.String
		}
		albumDTOs = append(albumDTOs, album)
	}

	return albumDTOs, nil
}

func (s *Service) SearchAlbums(ctx context.Context, query string, limit uint64) ([]dto.AlbumSearch, error) {
	const op = "service.SearchAlbums"

	repoResults, err := s.repo.SearchAlbums(ctx, query, limit)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to search albums in repository: %w", op, mapError(err))
	}

	if len(repoResults) == 0 {
		return []dto.AlbumSearch{}, nil
	}

	artistIDs := make([]uuid.UUID, 0, len(repoResults))
	uniqueArtistIDs := make(map[uuid.UUID]bool)
	for _, result := range repoResults {
		if !uniqueArtistIDs[result.Album.ArtistID] {
			uniqueArtistIDs[result.Album.ArtistID] = true
			artistIDs = append(artistIDs, result.Album.ArtistID)
		}
	}

	artists, err := s.GetArtistsByIDs(ctx, artistIDs)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get artists for albums: %w", op, mapError(err))
	}

	artistsMap := make(map[string]dto.Artist, len(artists))
	for _, artist := range artists {
		artistsMap[artist.ID] = artist
	}

	dtoResults := make([]dto.AlbumSearch, len(repoResults))
	for i, result := range repoResults {
		artist, ok := artistsMap[result.Album.ArtistID.String()]
		if !ok {
			continue
		}

		dtoResults[i] = dto.AlbumSearch{
			Album: dto.Album{
				ID:          result.Album.ID.String(),
				Title:       result.Album.Title,
				Type:        result.Album.Type,
				AvatarURL:   result.Album.AvatarURL,
				ReleaseDate: result.Album.ReleaseDate.Format(dateFormat),
				Artists: []dto.ArtistForAlbum{
					{
						ID:        artist.ID,
						Name:      artist.Name,
						AvatarURL: artist.AvatarURL,
					},
				},
			},
			Rank: result.Rank,
		}
		if result.Album.Description.Valid {
			dtoResults[i].Album.Description = result.Album.Description.String
		}
	}

	return dtoResults, nil
}
