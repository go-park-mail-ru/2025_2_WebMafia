package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type SupportTicket struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Title       string
	Description string
	Category    string
	Status      string
	Rating      sql.NullInt32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
