package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type exercise struct {
	id   uuid.UUID
	name string
	memo string
}

type record struct {
	id        uuid.UUID
	memo      string
	createdAt time.Time

	exercise exercise
	sets     []set
}

type set struct {
	id        uuid.UUID
	recordID  uuid.UUID
	weight    decimal.Decimal
	reps      uint
	createdAt time.Time
}
