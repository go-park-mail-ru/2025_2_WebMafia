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

	playCount, err := s.trackService.GetTotalPlaysByArtistID(ctx, artistModel.ID)
	if err != nil {
		playCount = 0
	}

	artistDTO := &dto.Artist{
		ID:        artistModel.ID.String(),
		Name:      artistModel.Name,
		AvatarURL: artistModel.AvatarURL,
		PlayCount: playCount,
	}

	if artistModel.Description.Valid {
		artistDTO.Description = artistModel.Description.String
	}

	return artistDTO, nil
}

func (s *Service) GetAllArtists(ctx context.Context, limit, offset uint64) ([]dto.Artist, error) {
	const op = "service.GetAllArtists"
	artistModels, err := s.repo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get all artists: %w", op, mapError(err))
	}

	if len(artistModels) == 0 {
		return []dto.Artist{}, nil
	}

	artistIDs := make([]uuid.UUID, len(artistModels))
	for i, artist := range artistModels {
		artistIDs[i] = artist.ID
	}

	playsMap, err := s.trackService.GetTotalPlaysByArtistIDs(ctx, artistIDs)
	if err != nil {
		playsMap = make(map[uuid.UUID]int64)
	}

	artistDTOs := make([]dto.Artist, len(artistModels))
	for i, artist := range artistModels {
		artistDTOs[i] = dto.Artist{
			ID:        artist.ID.String(),
			Name:      artist.Name,
			AvatarURL: artist.AvatarURL,
			PlayCount: playsMap[artist.ID],
		}
		if artist.Description.Valid {
			artistDTOs[i].Description = artist.Description.String
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

	playsMap, err := s.trackService.GetTotalPlaysByArtistIDs(ctx, ids)
	if err != nil {
		playsMap = make(map[uuid.UUID]int64)
	}

	artistDTOs := make([]dto.Artist, len(artistModels))
	for i, artist := range artistModels {
		artistDTOs[i] = dto.Artist{
			ID:        artist.ID.String(),
			Name:      artist.Name,
			AvatarURL: artist.AvatarURL,
			PlayCount: playsMap[artist.ID],
		}
		if artist.Description.Valid {
			artistDTOs[i].Description = artist.Description.String
		}
	}

	return artistDTOs, nil
}
