package aggregate

import (
	"time"

	"github.com/decorickey/workout-tracker-app/domain/entity"
	"github.com/google/uuid"
)

type exercise struct {
	id   uuid.UUID
	name string
	memo string
}

type record struct {
	id       uuid.UUID
	exercise exercise
	sets     []entity.Set
	datetime time.Time
	memo     string
}
