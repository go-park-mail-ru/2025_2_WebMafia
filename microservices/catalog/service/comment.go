package service

import (
	"context"
	"fmt"
	"spotify/internal/model"
	"spotify/microservices/catalog/dto"
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

	// TODO: здесь нужно сделать gRPC вызов в Auth, чтобы получить Login и AvatarURL по UserID.
	res := &dto.Comment{
		ID:         newComment.ID.String(),
		Text:       newComment.Text,
		TrackID:    newComment.TrackID.String(),
		UserID:     newComment.UserID.String(),
		UserLogin:  "",
		UserAvatar: "",
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

	comments := make([]dto.Comment, 0, len(models))
	for _, m := range models {
		comments = append(comments, dto.Comment{
			ID:         m.ID.String(),
			Text:       m.Text,
			TrackID:    m.TrackID.String(),
			UserID:     m.UserID.String(),
			UserLogin:  "",
			UserAvatar: "",
			CreatedAt:  m.CreatedAt,
		})
	}

	// TODO: Здесь нужно собрать все UserID и сделать Batch-запрос в Auth сервис

	return comments, nil
}
