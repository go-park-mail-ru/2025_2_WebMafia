package service

import (
	"context"
	"errors"
	"fmt"
	"spotify/internal/artist/dto"
	"spotify/internal/artist/repository/postgres"

	"github.com/google/uuid"
)

func (s *Service) GetArtistByID(ctx context.Context, id uuid.UUID) (*dto.Artist, error) {
	artistModel, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, postgres.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("service.GetArtistByID: %w", err)
	}

	artistDTO := &dto.Artist{
		ID:        artistModel.ID.String(),
		Name:      artistModel.Name,
		AvatarURL: artistModel.AvatarURL,
	}

	return artistDTO, nil
}

func (s *Service) GetAllArtists(ctx context.Context) ([]dto.Artist, error) {
	artistModels, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("service.GetAllArtists: %w", err)
	}

	artistDTOs := make([]dto.Artist, len(artistModels))
	for i, artist := range artistModels {
		artistDTOs[i] = dto.Artist{
			ID:        artist.ID.String(),
			Name:      artist.Name,
			AvatarURL: artist.AvatarURL,
		}
	}

	return artistDTOs, nil
}
