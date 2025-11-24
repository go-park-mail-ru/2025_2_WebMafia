package service

import (
	"context"
	"errors"
	"fmt"
	"spotify/internal/model"
	"spotify/microservices/playlist/dto"
	"spotify/microservices/playlist/repository/postgres"
	pbCatalog "spotify/proto/catalog"
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
	playlist, err := s.repo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, mapRepositoryError(err)
	}

	repoUpd := postgres.PlaylistUpdate{
		Title:       req.Title,
		Description: req.Description,
		IsFavorite:  req.IsFavorite,
	}

	err = s.repo.UpdatePlaylist(ctx, req.ID, repoUpd)
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

func (s *Service) validateTrack(ctx context.Context, trackID string) error {
	if trackID == "" {
		return errors.New("empty track id")
	}

	_, err := s.catalog.GetTrackByID(ctx, &pbCatalog.GetTrackByIDRequest{
		Id: trackID,
	})
	if err != nil {
		return fmt.Errorf("track not found: %w", ErrNotFound)
	}

	return nil
}

func (s *Service) AddTrackToFavorite(ctx context.Context, req dto.AddTrackToFavoriteRequest) error {
	if err := s.validateTrack(ctx, req.TrackID); err != nil {
		return err
	}

	fav, err := s.getOrCreateFavorite(ctx, req.UserID)
	if err != nil {
		return err
	}

	return mapRepositoryError(s.repo.AddTrackToPlaylist(ctx, fav.ID, req.TrackID))
}

func (s *Service) UploadPlaylistAvatar(ctx context.Context, req dto.UploadPlaylistAvatarRequest) (*dto.UploadPlaylistAvatarResponse, error) {
	playlist, err := s.repo.GetByID(ctx, req.PlaylistID)
	if err != nil {
		return nil, mapRepositoryError(err)
	}
	objectName, err := s.storage.UploadAvatar(ctx, req.File, req.Size, req.ContentType)
	if err != nil {
		return nil, err
	}

	if playlist.AvatarURL != "" {
		_ = s.storage.DeleteAvatar(ctx, playlist.AvatarURL)
	}

	if err := s.repo.UpdatePlaylistAvatar(ctx, req.PlaylistID, objectName); err != nil {
		_ = s.storage.DeleteAvatar(ctx, objectName)
		return nil, mapRepositoryError(err)
	}
	return &dto.UploadPlaylistAvatarResponse{URL: objectName}, nil
}

func (s *Service) DeletePlaylistAvatar(ctx context.Context, req dto.DeletePlaylistAvatarRequest) error {
	playlist, err := s.repo.GetByID(ctx, req.PlaylistID)
	if err != nil {
		return mapRepositoryError(err)
	}

	if playlist.AvatarURL != "" {
		_ = s.storage.DeleteAvatar(ctx, playlist.AvatarURL)
	}

	if err := s.repo.UpdatePlaylistAvatar(ctx, req.PlaylistID, ""); err != nil {
		return mapRepositoryError(err)
	}
	return nil
}

func (s *Service) GetPlaylistWithTracks(ctx context.Context, id uuid.UUID) (*dto.Playlist, error) {
	p, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, mapRepositoryError(err)
	}

	trackIDs, err := s.repo.GetTracksByPlaylist(ctx, id)
	if err != nil {
		return nil, mapRepositoryError(err)
	}

	return &dto.Playlist{
		ID:          p.ID.String(),
		Title:       p.Title,
		Description: p.Description,
		IsFavorite:  p.IsFavorite,
		AvatarURL:   p.AvatarURL,
		Tracks:      trackIDs,
	}, nil
}

func (s *Service) AddTrackToPlaylist(ctx context.Context, req dto.AddTrackToPlaylistRequest) error {
	if err := s.validateTrack(ctx, req.TrackID); err != nil {
		return err
	}
	return mapRepositoryError(s.repo.AddTrackToPlaylist(ctx, req.PlaylistID, req.TrackID))
}

func (s *Service) RemoveTrackFromPlaylist(ctx context.Context, req dto.RemoveTrackFromPlaylistRequest) error {
	if req.TrackID == "" {
		return errors.New("empty track id")
	}
	return mapRepositoryError(s.repo.RemoveTrackFromPlaylist(ctx, req.PlaylistID, req.TrackID))
}
