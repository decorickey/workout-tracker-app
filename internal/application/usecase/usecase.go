package usecase

import (
	"context"

	"github.com/decorickey/workout-tracker-app/internal/domain/aggregate"
	"github.com/decorickey/workout-tracker-app/internal/domain/repository"
	"github.com/google/uuid"
)

type ExerciseUsecase interface {
	Create(ctx context.Context, name, memo string) error
	Update(ctx context.Context, id uuid.UUID, name, memo string) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type exerciseUsecase struct {
	repo repository.ExerciseRepository
}

func NewExerciseUsecase(repo repository.ExerciseRepository) ExerciseUsecase {
	return exerciseUsecase{repo: repo}
}

func (u exerciseUsecase) Create(ctx context.Context, name, memo string) error {
	exercise, err := aggregate.NewExercise(name, memo)
	if err != nil {
		return err
	}

	return u.repo.Create(ctx, exercise)
}

func (u exerciseUsecase) Update(ctx context.Context, id uuid.UUID, name, memo string) error {
	exercise, err := aggregate.NewExercise(name, memo)
	if err != nil {
		return err
	}

	return u.repo.Update(ctx, id, exercise)
}

func (u exerciseUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.repo.Delete(ctx, id)
}
