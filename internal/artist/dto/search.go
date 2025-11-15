package dto

import "spotify/internal/model"

type SearchResult struct {
	Artist model.Artist
	Rank   float32
}
