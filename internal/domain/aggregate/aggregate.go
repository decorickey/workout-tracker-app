package aggregate

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Exercise struct {
	name string
	memo string
}

type ExerciseRecord struct {
	exerciseID uuid.UUID
	memo       string
}

type ExerciseSet struct {
	exerciseRecordID uuid.UUID
	weight           decimal.Decimal
	reps             uint
}
