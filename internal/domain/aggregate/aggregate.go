package aggregate

import (
	"fmt"

	"github.com/decorickey/workout-tracker-app/internal/constant"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Exercise struct {
	name string
	memo string
}

func NewExercise(name, memo string) (Exercise, error) {
	if name == "" {
		return Exercise{}, fmt.Errorf("empty exercise name: %w", constant.ErrInvalidParameter)
	}

	return Exercise{name: name, memo: memo}, nil
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
