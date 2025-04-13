package dao

import "github.com/decorickey/workout-tracker-app/internal/application/dto"

type ExerciseDao interface {
	GetAll() ([]dto.Exercise, error)
}

type ExerciseRecordDao interface {
	GetAll() ([]dto.ExerciseRecord, error)
}
