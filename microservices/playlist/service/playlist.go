package service

import (
	"context"
	"errors"
	"fmt"
	"spotify/internal/ai"
	"spotify/internal/model"
	"spotify/microservices/playlist/dto"
	"spotify/microservices/playlist/repository/postgres"
	"strings"
	"time"

	pbCatalog "spotify/proto/catalog"

	"github.com/google/uuid"
)

const fallbackMaxTracks = 5

func (s *Service) CreatePlaylist(ctx context.Context, req dto.CreatePlaylistRequest) (*dto.Playlist, error) {
	const op = "service.CreatePlaylist"

	playlist := model.Playlist{
		ID:          uuid.New(),
		UserID:      req.UserID,
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
		CreatorID:   playlist.UserID.String(),
		Title:       playlist.Title,
		Description: playlist.Description,
		IsFavorite:  playlist.IsFavorite,
		AvatarURL:   playlist.AvatarURL,
		CreatedAt:   playlist.CreatedAt,
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
			CreatorID:   p.UserID.String(),
			Title:       p.Title,
			Description: p.Description,
			IsFavorite:  p.IsFavorite,
			AvatarURL:   p.AvatarURL,
			CreatedAt:   p.CreatedAt,
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
		CreatorID:   playlist.UserID.String(),
		Title:       playlist.Title,
		Description: playlist.Description,
		IsFavorite:  playlist.IsFavorite,
		AvatarURL:   playlist.AvatarURL,
		CreatedAt:   playlist.CreatedAt,
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
		CreatorID:   req.UserID.String(),
		Title:       playlist.Title,
		Description: playlist.Description,
		IsFavorite:  playlist.IsFavorite,
		AvatarURL:   playlist.AvatarURL,
		CreatedAt:   playlist.CreatedAt,
	}, nil
}

func (s *Service) AddTrackToFavorite(ctx context.Context, req dto.AddTrackToFavoriteRequest) error {
	const op = "service.AddTrackToFavorite"

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
			CreatorID:   p.UserID.String(),
			Title:       p.Title,
			Description: p.Description,
			IsFavorite:  p.IsFavorite,
			AvatarURL:   p.AvatarURL,
			Tracks:      []dto.Track{},
			CreatedAt:   p.CreatedAt,
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
		CreatorID:   p.UserID.String(),
		Title:       p.Title,
		Description: p.Description,
		IsFavorite:  p.IsFavorite,
		AvatarURL:   p.AvatarURL,
		Tracks:      tracks,
		CreatedAt:   p.CreatedAt,
	}, nil
}

func (s *Service) AddTrackToPlaylist(ctx context.Context, req dto.AddTrackToPlaylistRequest) error {
	const op = "service.AddTrackToPlaylist"

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

func (s *Service) AddAlbumToFavorite(ctx context.Context, req dto.AddAlbumToFavoriteRequest) error {
	const op = "service.AddAlbumToFavorite"

	if err := s.repo.AddAlbumToFavorite(ctx, req.UserID, req.AlbumID); err != nil {
		return fmt.Errorf("%s: repo add: %w", op, mapRepositoryError(err))
	}
	return nil
}

func (s *Service) RemoveAlbumFromFavorite(ctx context.Context, req dto.RemoveAlbumFromFavoriteRequest) error {
	const op = "service.RemoveAlbumFromFavorite"

	if err := s.repo.RemoveAlbumFromFavorite(ctx, req.UserID, req.AlbumID); err != nil {
		return fmt.Errorf("%s: repo remove: %w", op, mapRepositoryError(err))
	}
	return nil
}

func (s *Service) GetFavoriteAlbums(ctx context.Context, userID uuid.UUID) ([]dto.FavoriteAlbum, error) {
	const op = "service.GetFavoriteAlbums"

	recs, err := s.repo.GetFavoriteAlbumIDs(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: get ids: %w", op, mapRepositoryError(err))
	}
	if len(recs) == 0 {
		return []dto.FavoriteAlbum{}, nil
	}

	ids := make([]string, 0, len(recs))
	createdMap := make(map[string]time.Time, len(recs))

	for _, r := range recs {
		idStr := r.AlbumID.String()
		ids = append(ids, idStr)
		createdMap[idStr] = r.CreatedAt
	}

	resp, err := s.catalog.GetAlbumsByIDs(ctx, &pbCatalog.GetAlbumsByIDsRequest{Ids: ids})
	if err != nil {
		return nil, fmt.Errorf("%s: batch load: %w", op, err)
	}

	out := make([]dto.FavoriteAlbum, 0, len(resp.Albums))

	for _, a := range resp.Albums {
		alb := dto.FavoriteAlbum{
			ID:        a.Id,
			CreatorID: userID.String(),
			Title:     a.Title,
			AvatarURL: a.AvatarUrl,
			Type:      a.Type,
			CreatedAt: createdMap[a.Id],
		}

		for _, artist := range a.Artists {
			alb.Artists = append(alb.Artists, dto.ArtistForAlbum{
				ID:   artist.Id,
				Name: artist.Name,
			})
		}

		out = append(out, alb)
	}
	return out, nil
}

func (s *Service) AddArtistToFavorite(ctx context.Context, req dto.AddArtistToFavoriteRequest) error {
	const op = "service.AddArtistToFavorite"
	if err := s.repo.AddArtistToFavorite(ctx, req.UserID, req.ArtistID); err != nil {
		return fmt.Errorf("%s: repo add: %w", op, mapRepositoryError(err))
	}
	return nil
}

func (s *Service) RemoveArtistFromFavorite(ctx context.Context, req dto.RemoveArtistFromFavoriteRequest) error {
	const op = "service.RemoveArtistFromFavorite"

	if err := s.repo.RemoveArtistFromFavorite(ctx, req.UserID, req.ArtistID); err != nil {
		return fmt.Errorf("%s: repo remove: %w", op, mapRepositoryError(err))
	}
	return nil
}

func (s *Service) GetFavoriteArtists(ctx context.Context, userID uuid.UUID) ([]dto.FavoriteArtist, error) {
	const op = "service.GetFavoriteArtists"

	recs, err := s.repo.GetFavoriteArtistIDs(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: get ids: %w", op, mapRepositoryError(err))
	}
	if len(recs) == 0 {
		return []dto.FavoriteArtist{}, nil
	}

	ids := make([]string, 0, len(recs))
	createdMap := make(map[string]time.Time, len(recs))

	for _, r := range recs {
		idStr := r.ArtistID.String()
		ids = append(ids, idStr)
		createdMap[idStr] = r.CreatedAt
	}

	resp, err := s.catalog.GetArtistsByIDs(ctx, &pbCatalog.GetArtistsByIDsRequest{Ids: ids})
	if err != nil {
		return nil, fmt.Errorf("%s: batch load: %w", op, err)
	}

	out := make([]dto.FavoriteArtist, 0, len(resp.Artists))

	for _, a := range resp.Artists {
		out = append(out, dto.FavoriteArtist{
			ID:        a.Id,
			CreatorID: userID.String(),
			Name:      a.Name,
			AvatarURL: a.AvatarUrl,
			PlayCount: a.PlayCount,
			CreatedAt: createdMap[a.Id],
		})
	}
	return out, nil
}

func (s *Service) GeneratePlaylistMeta(ctx context.Context, playlistID uuid.UUID) (*dto.GeneratedMeta, error) {
	const op = "service.GeneratePlaylistMeta"

	p, err := s.repo.GetByID(ctx, playlistID)
	if err != nil {
		return nil, fmt.Errorf("%s: load playlist: %w", op, mapRepositoryError(err))
	}

	trackIDs, err := s.repo.GetTracksByPlaylist(ctx, playlistID)
	if err != nil {
		return nil, fmt.Errorf("%s: get tracks: %w", op, mapRepositoryError(err))
	}

	if len(trackIDs) == 0 {
		return &dto.GeneratedMeta{
			Title:       p.Title,
			Description: p.Description,
		}, nil
	}

	resp, err := s.catalog.GetTracksByIDs(
		ctx,
		&pbCatalog.GetTracksByIDsRequest{Ids: trackIDs},
	)
	if err != nil {
		return nil, fmt.Errorf("%s: load tracks: %w", op, err)
	}

	tracks := make([]dto.Track, 0, len(resp.Tracks))
	for _, t := range resp.Tracks {
		track := dto.Track{
			ID:        t.Id,
			Title:     t.Title,
			DurationS: int(t.DurationS),
			FileURL:   t.FileUrl,
		}

		if t.Album != nil {
			track.Album = dto.Album{
				ID:        t.Album.Id,
				Title:     t.Album.Title,
				AvatarURL: t.Album.AvatarUrl,
			}
		}

		for _, a := range t.Artists {
			track.Artists = append(track.Artists, dto.Artist{
				ID:   a.Id,
				Name: a.Name,
			})
		}

		tracks = append(tracks, track)
	}

	select {
	case s.aiSem <- struct{}{}:
		defer func() { <-s.aiSem }()
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	metas, err := s.ai.GeneratePlaylistMeta(ctx, tracks)
	if err != nil {
		switch {
		case errors.Is(err, ai.ErrAIRateLimit),
			errors.Is(err, ai.ErrAIUnavailable),
			errors.Is(err, ai.ErrAINoChoices):
			return &dto.GeneratedMeta{
				Title:       fallbackTitle(tracks),
				Description: fallbackDescription(tracks),
				Source:      "fallback",
				Warning:     "ai_rate_limit",
			}, nil

		case errors.Is(err, ai.ErrAIAuth):
			return nil, fmt.Errorf("%s: ai auth error: %w", op, err)

		default:
			return nil, fmt.Errorf("%s: ai error: %w", op, err)
		}
	}
	meta := metas[0]

	return &dto.GeneratedMeta{
		Title:       meta.Title,
		Description: meta.Description,
		Source:      "ai",
	}, nil
}

func fallbackTitle(tracks []dto.Track) string {
	if len(tracks) == 0 {
		return "Мой плейлист"
	}

	if len(tracks) == 1 {
		return tracks[0].Title
	}

	return fmt.Sprintf("%s и другие треки", tracks[0].Title)
}

func fallbackDescription(tracks []dto.Track) string {
	if len(tracks) == 0 {
		return ""
	}

	max := fallbackMaxTracks
	if len(tracks) < max {
		max = len(tracks)
	}
	names := make([]string, 0, max)
	for i := 0; i < max; i++ {
		names = append(names, tracks[i].Title)
	}

	if len(tracks) <= max {
		return fmt.Sprintf(
			"Плейлист на основе треков: %s.",
			strings.Join(names, ", "),
		)
	}
	return fmt.Sprintf(
		"Плейлист на основе треков: %s и других.",
		strings.Join(names, ", "),
	)
}

func (s *Service) ConfirmPlaylistMeta(ctx context.Context, playlistID uuid.UUID, title string, description string) error {
	const op = "service.ConfirmPlaylistMeta"

	upd := postgres.PlaylistUpdate{
		Title:       &title,
		Description: &description,
	}

	if err := s.repo.UpdatePlaylist(ctx, playlistID, upd); err != nil {
		return fmt.Errorf("%s: update: %w", op, mapRepositoryError(err))
	}

	return nil
}
