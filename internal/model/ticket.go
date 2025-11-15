package model

import (
	"database/sql"

	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Status      string
	Category    string
	Title       string
	Description string
	Rating      sql.NullInt32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
