package repositories

import (
	"entities"
)

type ProgramRepository interface {
	GetAll() ([]entities.WorkoutProgram, error)
	GetByName(name string) (*entities.WorkoutProgram, error)
}
