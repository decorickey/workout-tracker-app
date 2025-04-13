package repository

import (
	"context"

	"github.com/decorickey/workout-tracker-app/internal/domain/aggregate"
	"github.com/google/uuid"
)

type ExerciseRepository interface {
	Create(context.Context, aggregate.Exercise) error
	Update(context.Context, uuid.UUID, aggregate.Exercise) error
	Delete(context.Context, uuid.UUID) error
}

type ExerciseRecordRepository interface {
	Create(context.Context, aggregate.ExerciseRecord) error
	Update(context.Context, uuid.UUID, aggregate.ExerciseRecord) error
	Delete(context.Context, uuid.UUID) error
}

type ExerciseSetRepository interface {
	Create(context.Context, aggregate.ExerciseSet) error
	Update(context.Context, aggregate.ExerciseSet) error
	Delete(context.Context, uuid.UUID) error
}
