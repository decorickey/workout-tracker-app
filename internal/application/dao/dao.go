package dao

import (
	"context"

	"github.com/decorickey/workout-tracker-app/internal/application/dto"
)

type ExerciseDao interface {
	GetAll(context.Context) ([]dto.Exercise, error)
}

type ExerciseRecordDao interface {
	GetAll(context.Context) ([]dto.ExerciseRecord, error)
}
