package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Track struct {
	ID          uuid.UUID
	Title       string
	DurationMs  int
	FileURL     string
	Description sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
