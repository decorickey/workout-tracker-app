package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Exercise struct {
	id   uuid.UUID
	name string
	memo string
}

type ExerciseRecord struct {
	id        uuid.UUID
	exercise  Exercise
	memo      string
	sets      []exerciseSet
	createdAt time.Time
	updatedAt time.Time
}

type exerciseSet struct {
	id               uuid.UUID
	exerciseRecordID uuid.UUID
	weight           decimal.Decimal
	reps             uint
	createdAt        time.Time
	updatedAt        time.Time
}
