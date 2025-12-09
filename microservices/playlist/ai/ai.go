package ai

import (
	"context"
	"spotify/microservices/playlist/dto"
)

type IAIGenerator interface {
	GeneratePlaylistMeta(ctx context.Context, tracks []dto.Track) (title string, description string, err error)
}
