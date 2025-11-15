package dto

import "spotify/internal/model"

type SearchResult struct {
	Album model.Album
	Rank  float32
}
