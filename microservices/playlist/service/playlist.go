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
		return nil, fmt.Errorf("%s: create playlist: %w", op, mapRepositoryError(err))
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
		return nil, fmt.Errorf("%s: get by id: %w", op, mapRepositoryError(err))
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
		return nil, fmt.Errorf("%s: get all by user: %w", op, mapRepositoryError(err))
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
		return nil, fmt.Errorf("%s: load before update: %w", op, mapRepositoryError(err))
	}

	repoUpd := postgres.PlaylistUpdate{
		Title:       req.Title,
		Description: req.Description,
		IsFavorite:  req.IsFavorite,
	}

	if err := s.repo.UpdatePlaylist(ctx, req.ID, repoUpd); err != nil {
		return nil, fmt.Errorf("%s: update playlist: %w", op, mapRepositoryError(err))
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

	if err := s.repo.DeletePlaylist(ctx, req.ID); err != nil {
		return fmt.Errorf("%s: delete: %w", op, mapRepositoryError(err))
	}
	return nil
}

func (s *Service) getOrCreateFavorite(ctx context.Context, user uuid.UUID) (*model.Playlist, error) {
	const op = "service.getOrCreateFavorite"

	p, err := s.repo.GetFavoritePlaylist(ctx, user)
	if err != nil {
		err = mapRepositoryError(err)
		if !errors.Is(err, ErrNotFound) {
			return nil, fmt.Errorf("%s: get favorite: %w", op, err)
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
		return nil, fmt.Errorf("%s: create new favorite: %w", op, mapRepositoryError(err))
	}

	return &playlist, nil
}

func (s *Service) GetFavoritePlaylist(ctx context.Context, req dto.GetFavoritePlaylistRequest) (*dto.Playlist, error) {
	const op = "service.GetFavoritePlaylist"

	playlist, err := s.getOrCreateFavorite(ctx, req.UserID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
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
	const op = "service.validateTrack"

	if trackID == "" {
		return fmt.Errorf("%s: empty track id", op)
	}

	_, err := s.catalog.GetTrackByID(ctx, &pbCatalog.GetTrackByIDRequest{Id: trackID})
	if err != nil {
		return fmt.Errorf("%s: track not found: %w", op, ErrNotFound)
	}

	return nil
}

func (s *Service) AddTrackToFavorite(ctx context.Context, req dto.AddTrackToFavoriteRequest) error {
	const op = "service.AddTrackToFavorite"

	if err := s.validateTrack(ctx, req.TrackID); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	fav, err := s.getOrCreateFavorite(ctx, req.UserID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := s.repo.AddTrackToPlaylist(ctx, fav.ID, req.TrackID); err != nil {
		return fmt.Errorf("%s: repo add track: %w", op, mapRepositoryError(err))
	}

	return nil
}

func (s *Service) UploadPlaylistAvatar(ctx context.Context, req dto.UploadPlaylistAvatarRequest) (*dto.UploadPlaylistAvatarResponse, error) {
	const op = "service.UploadPlaylistAvatar"

	playlist, err := s.repo.GetByID(ctx, req.PlaylistID)
	if err != nil {
		return nil, fmt.Errorf("%s: get playlist: %w", op, mapRepositoryError(err))
	}

	objectName, err := s.storage.UploadAvatar(ctx, req.File, req.Size, req.ContentType)
	if err != nil {
		return nil, fmt.Errorf("%s: upload avatar: %w", op, err)
	}

	if playlist.AvatarURL != "" {
		_ = s.storage.DeleteAvatar(ctx, playlist.AvatarURL)
	}

	if err := s.repo.UpdatePlaylistAvatar(ctx, req.PlaylistID, objectName); err != nil {
		_ = s.storage.DeleteAvatar(ctx, objectName)
		return nil, fmt.Errorf("%s: update avatar in repo: %w", op, mapRepositoryError(err))
	}

	return &dto.UploadPlaylistAvatarResponse{URL: objectName}, nil
}

func (s *Service) DeletePlaylistAvatar(ctx context.Context, req dto.DeletePlaylistAvatarRequest) error {
	const op = "service.DeletePlaylistAvatar"

	playlist, err := s.repo.GetByID(ctx, req.PlaylistID)
	if err != nil {
		return fmt.Errorf("%s: get playlist: %w", op, mapRepositoryError(err))
	}

	if playlist.AvatarURL != "" {
		_ = s.storage.DeleteAvatar(ctx, playlist.AvatarURL)
	}

	if err := s.repo.UpdatePlaylistAvatar(ctx, req.PlaylistID, ""); err != nil {
		return fmt.Errorf("%s: clear avatar: %w", op, mapRepositoryError(err))
	}

	return nil
}

func (s *Service) GetPlaylistWithTracks(ctx context.Context, id uuid.UUID) (*dto.Playlist, error) {
	const op = "service.GetPlaylistWithTracks"

	p, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%s: load playlist: %w", op, mapRepositoryError(err))
	}

	trackIDs, err := s.repo.GetTracksByPlaylist(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%s: get tracks by playlist: %w", op, mapRepositoryError(err))
	}

	if len(trackIDs) == 0 {
		return &dto.Playlist{
			ID:          p.ID.String(),
			Title:       p.Title,
			Description: p.Description,
			IsFavorite:  p.IsFavorite,
			AvatarURL:   p.AvatarURL,
			Tracks:      []dto.Track{},
		}, nil
	}

	resp, err := s.catalog.GetTracksByIDs(ctx, &pbCatalog.GetTracksByIDsRequest{Ids: trackIDs})
	if err != nil {
		return nil, fmt.Errorf("%s: batch load tracks: %w", op, err)
	}

	tracks := make([]dto.Track, 0, len(resp.Tracks))
	for _, t := range resp.Tracks {
		trackDTO := dto.Track{
			ID:        t.Id,
			Title:     t.Title,
			DurationS: int(t.DurationS),
			FileURL:   t.FileUrl,
			Album: dto.Album{
				ID:        t.Album.Id,
				Title:     t.Album.Title,
				AvatarURL: t.Album.AvatarUrl,
			},
		}

		for _, a := range t.Artists {
			trackDTO.Artists = append(trackDTO.Artists, dto.Artist{
				ID:   a.Id,
				Name: a.Name,
			})
		}

		tracks = append(tracks, trackDTO)
	}

	return &dto.Playlist{
		ID:          p.ID.String(),
		Title:       p.Title,
		Description: p.Description,
		IsFavorite:  p.IsFavorite,
		AvatarURL:   p.AvatarURL,
		Tracks:      tracks,
	}, nil
}

func (s *Service) AddTrackToPlaylist(ctx context.Context, req dto.AddTrackToPlaylistRequest) error {
	const op = "service.AddTrackToPlaylist"

	if err := s.validateTrack(ctx, req.TrackID); err != nil {
		return fmt.Errorf("%s: validate: %w", op, err)
	}

	if err := s.repo.AddTrackToPlaylist(ctx, req.PlaylistID, req.TrackID); err != nil {
		return fmt.Errorf("%s: repo add: %w", op, mapRepositoryError(err))
	}
	return nil
}

func (s *Service) RemoveTrackFromPlaylist(ctx context.Context, req dto.RemoveTrackFromPlaylistRequest) error {
	const op = "service.RemoveTrackFromPlaylist"

	if req.TrackID == "" {
		return fmt.Errorf("%s: empty track id", op)
	}
	if err := s.repo.RemoveTrackFromPlaylist(ctx, req.PlaylistID, req.TrackID); err != nil {
		return fmt.Errorf("%s: repo remove: %w", op, mapRepositoryError(err))
	}
	return nil
}
