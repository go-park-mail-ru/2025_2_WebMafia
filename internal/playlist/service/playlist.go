package service

import (
	"context"
	"errors"
	"spotify/internal/model"
	"spotify/internal/playlist/dto"
	"time"

	"github.com/google/uuid"
)

func (s *Service) CreatePlaylist(ctx context.Context, req dto.CreatePlaylistRequest) (*dto.Playlist, error) {
	const op = "service.CreatePlaylist"

	playlist := model.Playlist{
		ID:          uuid.New(),
		Title:       req.Title,
		Description: req.Description,
		IsFavorite:  false,
		AvatarURL:   "",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.CreatePlaylist(ctx, playlist, req.UserID); err != nil {
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

	playlist, err := s.repo.GetByID(ctx, req.ID)
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

	playlists, err := s.repo.GetAllByUser(ctx, req.UserID, req.Limit, req.Offset)
	if err != nil {
		return nil, mapRepositoryError(err)
	}

	res := make([]dto.Playlist, 0, len(playlists))
	for _, p := range playlists {
		res = append(res, dto.Playlist{
			ID:          p.ID.String(),
			Title:       p.Title,
			Description: p.Description,
			IsFavorite:  p.IsFavorite,
			AvatarURL:   p.AvatarURL,
		})
	}

	return res, nil
}

func (s *Service) UpdatePlaylist(ctx context.Context, req dto.UpdatePlaylistRequest) (*dto.Playlist, error) {
	const op = "service.UpdatePlaylist"

	playlist, err := s.repo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, mapRepositoryError(err)
	}

	if req.Title != nil {
		playlist.Title = *req.Title
	}
	if req.Description != nil {
		playlist.Description = *req.Description
	}
	if req.IsFavorite != nil {
		playlist.IsFavorite = *req.IsFavorite
	}

	if err := s.repo.UpdatePlaylist(ctx, *playlist); err != nil {
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

	return mapRepositoryError(s.repo.DeletePlaylist(ctx, req.ID))
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
	playlist, err := s.getOrCreateFavorite(ctx, req.UserID)
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
	fav, err := s.getOrCreateFavorite(ctx, req.UserID)
	if err != nil {
		return err
	}

	return mapRepositoryError(s.repo.AddTrackToPlaylist(ctx, fav.ID, req.TrackID))
}
