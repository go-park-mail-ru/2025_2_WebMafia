package model

import (
	"time"

	"github.com/google/uuid"
)

type Genre struct {
	ID          uuid.UUID
	Name        string
	Description string
	CreatedAt   time.Time
}
