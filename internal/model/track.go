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
	FileURL     string
	Description sql.NullString
	PlayCount   int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
