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

type Record struct {
	id        uuid.UUID
	exercise  Exercise
	memo      string
	sets      []recordSet
	createdAt time.Time
	updatedAt time.Time
}

type recordSet struct {
	id        uuid.UUID
	recordID  uuid.UUID
	weight    decimal.Decimal
	reps      uint
	createdAt time.Time
	updatedAt time.Time
}
