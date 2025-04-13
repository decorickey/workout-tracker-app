package dao

import "github.com/decorickey/workout-tracker-app/internal/application/dto"

type ExerciseDao interface {
	GetAll() ([]dto.Exercise, error)
}

type RecordDao interface {
	GetAll() ([]dto.Record, error)
}
