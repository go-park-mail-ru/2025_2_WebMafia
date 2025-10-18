package service

import (
	"context"
	"spotify/internal/album/dto"
	artistDTO "spotify/internal/artist/dto"

	"github.com/google/uuid"
)

const (
	dateFormat = "2006-01-02"
)

func (s *Service) GetAlbumByID(ctx context.Context, id uuid.UUID) (*dto.Album, error) {
	albumModel, err := s.albumRepo.GetByID(ctx, id)
	if err != nil {
		return nil, mapError(err, "service.GetAlbumByID")
	}

	artist, err := s.artistService.GetArtistByID(ctx, albumModel.ArtistID)
	if err != nil {
		return nil, mapError(err, "service.GetAlbumByID: failed to get an artist")
	}

	albumDTO := &dto.Album{
		ID:          albumModel.ID.String(),
		Title:       albumModel.Title,
		AvatarURL:   albumModel.AvatarURL,
		ReleaseDate: albumModel.ReleaseDate.Format(dateFormat),
		Artists: []dto.Artist{
			{
				ID:        artist.ID,
				Name:      artist.Name,
				AvatarURL: artist.AvatarURL,
			},
		},
	}

	return albumDTO, nil
}

func (s *Service) GetAllAlbums(ctx context.Context, limit, offset uint64) ([]dto.Album, error) {
	albumModels, err := s.albumRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, mapError(err, "service.GetAllAlbums: failed to get all albums")
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

	artists, err := s.artistService.GetArtistsByIDs(ctx, artistIDs)
	if err != nil {
		return nil, mapError(err, "service.GetAllAlbums: failed to get an artist")
	}

	artistsMap := make(map[string]artistDTO.Artist, len(artists))
	for _, artist := range artists {
		artistsMap[artist.ID] = artist
	}

	albumDTOs := make([]dto.Album, 0, len(albumModels))
	for _, albumModel := range albumModels {
		artist, ok := artistsMap[albumModel.ArtistID.String()]
		if !ok {
			continue
		}

		albumDTOs = append(albumDTOs, dto.Album{
			ID:          albumModel.ID.String(),
			Title:       albumModel.Title,
			AvatarURL:   albumModel.AvatarURL,
			ReleaseDate: albumModel.ReleaseDate.Format(dateFormat),
			Artists: []dto.Artist{
				{
					ID:        artist.ID,
					Name:      artist.Name,
					AvatarURL: artist.AvatarURL,
				},
			},
		})
	}

	return albumDTOs, nil
}

func (s *Service) GetAlbumsByIDs(ctx context.Context, ids []uuid.UUID) ([]dto.Album, error) {
	albumModels, err := s.albumRepo.GetByIDs(ctx, ids)
	if err != nil {
		return nil, mapError(err, "service.GetAlbumsByIDs: failed to get albums by ids")
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

	artists, err := s.artistService.GetArtistsByIDs(ctx, artistIDs)
	if err != nil {
		return nil, mapError(err, "service.GetAlbumsByIDs: failed to get an artist")
	}

	artistsMap := make(map[string]artistDTO.Artist, len(artists))
	for _, artist := range artists {
		artistsMap[artist.ID] = artist
	}

	albumDTOs := make([]dto.Album, 0, len(albumModels))
	for _, albumModel := range albumModels {
		artist, ok := artistsMap[albumModel.ArtistID.String()]
		if !ok {
			continue
		}
		albumDTOs = append(albumDTOs, dto.Album{
			ID:          albumModel.ID.String(),
			Title:       albumModel.Title,
			AvatarURL:   albumModel.AvatarURL,
			ReleaseDate: albumModel.ReleaseDate.Format(dateFormat),
			Artists: []dto.Artist{{
				ID:        artist.ID,
				Name:      artist.Name,
				AvatarURL: artist.AvatarURL,
			}},
		})
	}

	return albumDTOs, nil
}

func (s *Service) GetAlbumsByArtistID(ctx context.Context, artistID uuid.UUID, limit, offset uint64) ([]dto.Album, error) {
	albumModels, err := s.albumRepo.GetByArtistID(ctx, artistID, limit, offset)
	if err != nil {
		return nil, mapError(err, "service.GetAlbumsByArtistID: failed to get albums")
	}
	if len(albumModels) == 0 {
		return []dto.Album{}, nil
	}

	artist, err := s.artistService.GetArtistByID(ctx, artistID)
	if err != nil {
		return nil, mapError(err, "service.GetAlbumsByArtistID: failed to get artist")
	}

	albumDTOs := make([]dto.Album, 0, len(albumModels))
	for _, albumModel := range albumModels {
		albumDTOs = append(albumDTOs, dto.Album{
			ID:          albumModel.ID.String(),
			Title:       albumModel.Title,
			AvatarURL:   albumModel.AvatarURL,
			ReleaseDate: albumModel.ReleaseDate.Format(dateFormat),
			Artists: []dto.Artist{
				{
					ID:        artist.ID,
					Name:      artist.Name,
					AvatarURL: artist.AvatarURL,
				},
			},
		})
	}

	return albumDTOs, nil
}
