package dto

import "spotify/internal/model"

type SearchResult struct {
	Track model.Track
	Rank  float32
}
