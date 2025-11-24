package service

import (
	"context"
	"fmt"
	"sort"
	"spotify/microservices/catalog/dto"

	"github.com/google/uuid"
)

func (s *Service) GetArtistByID(ctx context.Context, id uuid.UUID) (*dto.Artist, error) {
	const op = "service.GetArtistByID"
	artistModel, err := s.repo.GetArtistByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get an artist by id: %w", op, mapError(err))
	}

	playCount, err := s.GetTotalPlaysByArtistID(ctx, artistModel.ID)
	if err != nil {
		playCount = 0
	}

	artistDTO := &dto.Artist{
		ID:        artistModel.ID.String(),
		Name:      artistModel.Name,
		AvatarURL: artistModel.AvatarURL,
		HeaderURL: artistModel.HeaderURL,
		PlayCount: playCount,
	}

	if artistModel.Description.Valid {
		artistDTO.Description = artistModel.Description.String
	}

	return artistDTO, nil
}

func (s *Service) GetAllArtists(ctx context.Context, limit, offset uint64) ([]dto.Artist, error) {
	const op = "service.GetAllArtists"
	artistModels, err := s.repo.GetAllArtists(ctx, limit, offset)
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

	playsMap, err := s.GetTotalPlaysByArtistIDs(ctx, artistIDs)
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

	sort.Slice(artistDTOs, func(i, j int) bool {
		return artistDTOs[i].PlayCount > artistDTOs[j].PlayCount
	})

	return artistDTOs, nil
}

func (s *Service) GetArtistsByIDs(ctx context.Context, ids []uuid.UUID) ([]dto.Artist, error) {
	const op = "service.GetArtistsByIDs"
	artistModels, err := s.repo.GetArtistsByIDs(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get an artists by ids: %w", op, mapError(err))
	}

	if len(artistModels) == 0 {
		return []dto.Artist{}, nil
	}

	playsMap, err := s.GetTotalPlaysByArtistIDs(ctx, ids)
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

	sort.Slice(artistDTOs, func(i, j int) bool {
		return artistDTOs[i].PlayCount > artistDTOs[j].PlayCount
	})

	return artistDTOs, nil
}

func (s *Service) SearchArtists(ctx context.Context, query string, limit uint64) ([]dto.ArtistSearch, error) {
	const op = "service.SearchArtists"

	repoResults, err := s.repo.SearchArtists(ctx, query, limit)
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to search artists in repository: %w", op, mapError(err))
	}

	if len(repoResults) == 0 {
		return []dto.ArtistSearch{}, nil
	}

	artistIDs := make([]uuid.UUID, len(repoResults))
	for i, result := range repoResults {
		artistIDs[i] = result.Artist.ID
	}

	playsMap, err := s.GetTotalPlaysByArtistIDs(ctx, artistIDs)
	if err != nil {
		playsMap = make(map[uuid.UUID]int64)
	}

	dtoResults := make([]dto.ArtistSearch, len(repoResults))
	for i, result := range repoResults {
		dtoResults[i] = dto.ArtistSearch{
			Artist: dto.Artist{
				ID:        result.Artist.ID.String(),
				Name:      result.Artist.Name,
				AvatarURL: result.Artist.AvatarURL,
				HeaderURL: result.Artist.HeaderURL,
				PlayCount: playsMap[result.Artist.ID],
			},
			Rank: result.Rank,
		}
		if result.Artist.Description.Valid {
			dtoResults[i].Artist.Description = result.Artist.Description.String
		}
	}

	return dtoResults, nil
}
