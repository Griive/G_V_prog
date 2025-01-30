package repositories

import (
	"errors"
)

type MemoryProgramRepo struct {
	programs []entities.WorkoutProgram
}

func NewMemoryProgramRepo() *MemoryProgramRepo {
	return &MemoryProgramRepo{
		programs: []entities.WorkoutProgram{
			{
				Name: "Грудь + Пресс",
				Exercises: []entities.Exercise{
					{Name: "Жим штанги лежа", Sets: 4, Reps: "8-12", RestSec: 90},
					{Name: "Жим гантелей на наклонной", Sets: 3, Reps: "10-12", RestSec: 90},
				},
			},
			// остальные программы
		},
	}
}

func (r *MemoryProgramRepo) GetAll() ([]entities.WorkoutProgram, error) {
	return r.programs, nil
}

func (r *MemoryProgramRepo) GetByName(name string) (*entities.WorkoutProgram, error) {
	for _, p := range r.programs {
		if p.Name == name {
			return &p, nil
		}
	}
	return nil, errors.New("program not found")
}
