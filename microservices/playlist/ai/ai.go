package ai

import (
	"context"
	"spotify/microservices/playlist/dto"
)

//go:generate mockgen -source=ai.go -destination=../mocks/ai/ai_mock.go -package=ai_mock
type IAIGenerator interface {
	GeneratePlaylistMeta(ctx context.Context, tracks []dto.Track) (title string, description string, err error)
}
