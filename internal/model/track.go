package model

import (
	"time"

	"github.com/google/uuid"
)

type Track struct {
	ID          uuid.UUID
	Title       string
	DurationMs  int
	FileURL     string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
