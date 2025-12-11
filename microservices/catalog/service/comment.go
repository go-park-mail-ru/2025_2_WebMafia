package service

import (
	"context"
	"fmt"
	"spotify/internal/model"
	"spotify/microservices/catalog/dto"
	pbAuth "spotify/proto/auth"
	"time"

	"github.com/google/uuid"
)

func (s *Service) PostComment(ctx context.Context, userID uuid.UUID, req dto.PostCommentRequest) (*dto.Comment, error) {
	const op = "service.PostComment"

	trackID, err := uuid.Parse(req.TrackID)
	if err != nil {
		return nil, fmt.Errorf("[%s]: invalid track uuid: %w", op, err)
	}

	newComment := model.Comment{
		ID:        uuid.New(),
		TrackID:   trackID,
		UserID:    userID,
		Text:      req.Text,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.CreateComment(ctx, newComment); err != nil {
		return nil, fmt.Errorf("[%s]: repo create error: %w", op, mapError(err))
	}

	resp, err := s.authClient.GetUsers(ctx, &pbAuth.GetUsersRequest{
		UserIds: []string{userID.String()},
	})

	if err != nil {
		return nil, fmt.Errorf("%s: auth batch error: %w", op, err)
	}

	var login, avatar string
	if len(resp.Users) > 0 {
		login = resp.Users[0].Login
		avatar = resp.Users[0].AvatarUrl
	}

	res := &dto.Comment{
		ID:         newComment.ID.String(),
		Text:       newComment.Text,
		TrackID:    newComment.TrackID.String(),
		UserID:     newComment.UserID.String(),
		UserLogin:  login,
		UserAvatar: avatar,
		CreatedAt:  newComment.CreatedAt,
	}

	return res, nil
}

func (s *Service) GetCommentsByTrackID(ctx context.Context, trackID uuid.UUID, limit, offset uint64) ([]dto.Comment, error) {
	const op = "service.GetCommentsByTrackID"

	models, err := s.repo.GetCommentsByTrackID(ctx, trackID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[%s]: repo get error: %w", op, mapError(err))
	}

	if len(models) == 0 {
		return []dto.Comment{}, nil
	}

	userIDs := make([]string, 0, len(models))
	for _, c := range models {
		userIDs = append(userIDs, c.UserID.String())
	}

	userResp, err := s.authClient.GetUsers(ctx, &pbAuth.GetUsersRequest{
		UserIds: userIDs,
	})
	if err != nil {
		return nil, fmt.Errorf("%s: auth batch error: %w", op, err)
	}

	userMap := make(map[string]*pbAuth.UserInfo, len(userResp.Users))
	for _, u := range userResp.Users {
		userMap[u.UserId] = u
	}

	comments := make([]dto.Comment, 0, len(models))
	for _, m := range models {
		ui := userMap[m.UserID.String()]

		var login, avatar string
		if ui != nil {
			login = ui.Login
			avatar = ui.AvatarUrl
		}

		comments = append(comments, dto.Comment{
			ID:         m.ID.String(),
			Text:       m.Text,
			TrackID:    m.TrackID.String(),
			UserID:     m.UserID.String(),
			UserLogin:  login,
			UserAvatar: avatar,
			CreatedAt:  m.CreatedAt,
		})
	}

	return comments, nil
}
