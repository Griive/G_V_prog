package usecases

import (
	"context"
	"fmt"
)

type WorkoutUseCase struct {
	repo repositories.ProgramRepository
}

func NewWorkoutUseCase(repo repositories.ProgramRepository) *WorkoutUseCase {
	return &WorkoutUseCase{repo: repo}
}

func (uc *WorkoutUseCase) GetPrograms(ctx context.Context) ([]entities.WorkoutProgram, error) {
	return uc.repo.GetAll()
}

func (uc *WorkoutUseCase) StartWorkout(ctx context.Context, programName string, progressFunc func(string)) error {
	program, err := uc.repo.GetByName(programName)
	if err != nil {
		return err
	}

	for _, exercise := range program.Exercises {
		progressFunc("\nУпражнение: " + exercise.Name)
		for set := 1; set <= exercise.Sets; set++ {
			progressFunc(fmt.Sprintf("Подход %d/%d. Нажмите Enter когда готовы...", set, exercise.Sets))
			Timer(exercise.RestSec, progressFunc)
		}
	}
	return nil
}
