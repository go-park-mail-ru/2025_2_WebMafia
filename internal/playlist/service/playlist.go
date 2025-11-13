package service

import (
	"context"
	"errors"
	"fmt"
	"spotify/internal/model"
	"spotify/internal/playlist/dto"
	"time"

	"github.com/google/uuid"
)

func (s *Service) CreatePlaylist(ctx context.Context, req dto.CreatePlaylistRequest) (*dto.Playlist, error) {
	const op = "service.CreatePlaylist"

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, fmt.Errorf("[%s]: invalid user ID: %w", op, err)
	}

	playlist := model.Playlist{
		ID:          uuid.New(),
		Title:       req.Title,
		Description: req.Description,
		IsFavorite:  false,
		AvatarURL:   "",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.CreatePlaylist(ctx, playlist, userID); err != nil {
		return nil, mapRepositoryError(err)
	}

	return &dto.Playlist{
		ID:          playlist.ID.String(),
		Title:       playlist.Title,
		Description: playlist.Description,
		IsFavorite:  playlist.IsFavorite,
		AvatarURL:   playlist.AvatarURL,
	}, nil
}

func (s *Service) GetPlaylist(ctx context.Context, req dto.GetPlaylistRequest) (*dto.Playlist, error) {
	const op = "service.GetPlaylist"

	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, fmt.Errorf("[%s]: invalid playlist ID: %w", op, err)
	}

	playlist, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, mapRepositoryError(err)
	}

	return &dto.Playlist{
		ID:          playlist.ID.String(),
		Title:       playlist.Title,
		Description: playlist.Description,
		IsFavorite:  playlist.IsFavorite,
		AvatarURL:   playlist.AvatarURL,
	}, nil
}

func (s *Service) GetPlaylistsByUser(ctx context.Context, req dto.GetPlaylistsByUserRequest) ([]dto.Playlist, error) {
	const op = "service.GetPlaylistsByUser"

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, fmt.Errorf("[%s]: invalid user ID: %w", op, err)
	}

	playlists, err := s.repo.GetAllByUser(ctx, userID, req.Limit, req.Offset)
	if err != nil {
		return nil, mapRepositoryError(err)
	}

	dtoList := make([]dto.Playlist, 0, len(playlists))

	for _, p := range playlists {
		dtoList = append(dtoList, dto.Playlist{
			ID:          p.ID.String(),
			Title:       p.Title,
			Description: p.Description,
			IsFavorite:  p.IsFavorite,
			AvatarURL:   p.AvatarURL,
		})
	}

	return dtoList, nil
}

func (s *Service) UpdatePlaylist(ctx context.Context, req dto.UpdatePlaylistRequest) (*dto.Playlist, error) {
	const op = "service.UpdatePlaylist"

	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, fmt.Errorf("[%s]: invalid playlist ID: %w", op, err)
	}

	playlist := model.Playlist{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		IsFavorite:  req.IsFavorite,
		AvatarURL:   "",
	}

	if err := s.repo.UpdatePlaylist(ctx, playlist); err != nil {
		return nil, mapRepositoryError(err)
	}

	return &dto.Playlist{
		ID:          playlist.ID.String(),
		Title:       playlist.Title,
		Description: playlist.Description,
		IsFavorite:  playlist.IsFavorite,
		AvatarURL:   playlist.AvatarURL,
	}, nil
}

func (s *Service) DeletePlaylist(ctx context.Context, req dto.DeletePlaylistRequest) error {
	const op = "service.DeletePlaylist"

	id, err := uuid.Parse(req.ID)
	if err != nil {
		return fmt.Errorf("[%s]: invalid playlist ID: %w", op, err)
	}

	if err := s.repo.DeletePlaylist(ctx, id); err != nil {
		return mapRepositoryError(err)
	}

	return nil
}

func (s *Service) getOrCreateFavorite(ctx context.Context, user uuid.UUID) (*model.Playlist, error) {
	p, err := s.repo.GetFavoritePlaylist(ctx, user)

	if err != nil {
		err = mapRepositoryError(err)
		if !errors.Is(err, ErrNotFound) {
			return nil, err
		}
	}
	if p != nil {
		return p, nil
	}

	playlist := model.Playlist{
		ID:          uuid.New(),
		Title:       "Понравившиеся треки",
		Description: "",
		AvatarURL:   "",
		IsFavorite:  true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.CreatePlaylist(ctx, playlist, user); err != nil {
		return nil, mapRepositoryError(err)
	}

	return &playlist, nil
}

func (s *Service) GetFavoritePlaylist(ctx context.Context, req dto.GetFavoritePlaylistRequest) (*dto.Playlist, error) {
	user, err := uuid.Parse(req.UserID)

	if err != nil {
		return nil, fmt.Errorf("invalid userID: %w", err)
	}

	playlist, err := s.getOrCreateFavorite(ctx, user)
	if err != nil {
		return nil, err
	}

	return &dto.Playlist{
		ID:          playlist.ID.String(),
		Title:       playlist.Title,
		Description: playlist.Description,
		IsFavorite:  playlist.IsFavorite,
		AvatarURL:   playlist.AvatarURL,
	}, nil
}

func (s *Service) AddTrackToFavorite(ctx context.Context, req dto.AddTrackToFavoriteRequest) error {
	user, err := uuid.Parse(req.UserID)
	if err != nil {
		return fmt.Errorf("invalid userID: %w", err)
	}

	fav, err := s.getOrCreateFavorite(ctx, user)
	if err != nil {
		return err
	}
	if err := s.repo.AddTrackToPlaylist(ctx, fav.ID, req.TrackID); err != nil {
		return mapRepositoryError(err)
	}
	return nil
}
