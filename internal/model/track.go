package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Track struct {
	ID          uuid.UUID
	Title       string
	DurationS   int
	Plays       int
	FileURL     string
	Description sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
