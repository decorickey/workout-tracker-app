package repository

import (
	"github.com/decorickey/workout-tracker-app/internal/domain/aggregate"
	"github.com/google/uuid"
)

type ExerciseRepository interface {
	Create(aggregate.Exercise) error
	Update(uuid.UUID, aggregate.Exercise) error
	Delete(uuid.UUID) error
}

type ExerciseRecordRepository interface {
	Create(aggregate.ExerciseRecord) error
	Update(uuid.UUID, aggregate.ExerciseRecord) error
	Delete(uuid.UUID) error
}

type ExerciseSetRepository interface {
	Create(aggregate.ExerciseSet) error
	Update(aggregate.ExerciseSet) error
	Delete(uuid.UUID) error
}
