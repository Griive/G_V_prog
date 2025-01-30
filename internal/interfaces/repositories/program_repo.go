package repositories

import (
	"main.go/internal/app/workout/entities"
)

type ProgramRepository interface {
	GetAll() ([]entities.WorkoutProgram, error)
	GetByName(name string) (*entities.WorkoutProgram, error)
}
