package service

import (
	"context"
	"errors"
	"fmt"
	"spotify/internal/album/dto"
	"spotify/internal/album/repository/postgres"

	"github.com/google/uuid"
)

func (s *Service) GetAlbumByID(ctx context.Context, id uuid.UUID) (*dto.Album, error) {
	albumModel, artistModel, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, postgres.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("service.GetAlbumByID: %w", err)
	}

	albumDTO := &dto.Album{
		ID:          albumModel.ID.String(),
		Title:       albumModel.Title,
		AvatarURL:   albumModel.AvatarURL,
		ReleaseDate: albumModel.ReleaseDate.Format("2006-01-02"),
		Artists: []dto.Artist{
			{
				ID:        artistModel.ID.String(),
				Name:      artistModel.Name,
				AvatarURL: artistModel.AvatarURL,
			},
		},
	}

	return albumDTO, nil
}

func (s *Service) GetAllAlbums(ctx context.Context) ([]dto.Album, error) {
	albumModels, artistModels, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("service.GetAllAlbums: %w", err)
	}

	albumDTOs := make([]dto.Album, len(albumModels))

	for i := range albumModels {
		albumDTOs[i] = dto.Album{
			ID:          albumModels[i].ID.String(),
			Title:       albumModels[i].Title,
			AvatarURL:   albumModels[i].AvatarURL,
			ReleaseDate: albumModels[i].ReleaseDate.Format("2006-01-02"),
			Artists: []dto.Artist{
				{
					ID:        artistModels[i].ID.String(),
					Name:      artistModels[i].Name,
					AvatarURL: artistModels[i].AvatarURL,
				},
			},
		}
	}

	return albumDTOs, nil
}
