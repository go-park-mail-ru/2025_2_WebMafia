package service

import (
	"context"
	"fmt"
	"spotify/internal/artist/dto"

	"github.com/google/uuid"
)

func (s *Service) GetArtistByID(ctx context.Context, id uuid.UUID) (*dto.Artist, error) {
	const op = "service.GetArtistByID"
	artistModel, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get an artist by id: %w", op, mapError(err))
	}

	artistDTO := &dto.Artist{
		ID:        artistModel.ID.String(),
		Name:      artistModel.Name,
		AvatarURL: artistModel.AvatarURL,
	}

	return artistDTO, nil
}

func (s *Service) GetAllArtists(ctx context.Context, limit, offset uint64) ([]dto.Artist, error) {
	const op = "service.GetAllArtists"
	artistModels, err := s.repo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get all artists: %w", op, mapError(err))
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

func (s *Service) GetArtistsByIDs(ctx context.Context, ids []uuid.UUID) ([]dto.Artist, error) {
	const op = "service.GetArtistsByIDs"
	artistModels, err := s.repo.GetByIDs(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get an artists by ids: %w", op, mapError(err))
	}

	if len(artistModels) == 0 {
		return []dto.Artist{}, nil
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
