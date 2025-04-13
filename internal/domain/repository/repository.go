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

type RecordRepository interface {
	Create(aggregate.Record) error
	Update(uuid.UUID, aggregate.Record) error
	Delete(uuid.UUID) error
}

type RecordSetRepository interface {
	Create(aggregate.RecordSet) error
	Update(aggregate.RecordSet) error
	Delete(uuid.UUID) error
}
