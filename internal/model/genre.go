package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Genre struct {
	ID          uuid.UUID
	Name        string
	Description sql.NullString
	CreatedAt   time.Time
}
