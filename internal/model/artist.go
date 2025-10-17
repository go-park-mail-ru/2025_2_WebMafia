package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Artist struct {
	ID          uuid.UUID
	Name        string
	AvatarURL   string
	Description sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
