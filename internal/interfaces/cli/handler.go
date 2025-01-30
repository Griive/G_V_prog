package cli

import (
	"context"
)

type CLIHandler struct {
	workoutUC *usercases.WorkoutUseCase
	timerUC   *usercases.TimerUseCase
	writer    Writer
}

type Writer interface {
	Print(string)
	Printf(string, ...interface{})
	ReadLine() (string, error)
}

func NewCLIHandler(uc *usercases.WorkoutUseCase, writer Writer) *CLIHandler {
	return &CLIHandler{workoutUC: uc, writer: writer}
}

func (h *CLIHandler) Start() error {
	for {
		h.writer.Print(menu)
		choice, _ := h.writer.ReadLine()

		switch choice {
		case "1":
			return h.selectProgram()
		case "0":
			return nil
		default:
			h.writer.Print("Неверный выбор")
		}
	}
}

func (h *CLIHandler) selectProgram() error {
	programs, _ := h.workoutUC.GetPrograms(context.Background())

	for i, p := range programs {
		h.writer.Printf("%d. %s\n", i+1, p.Name)
	}

	choice, _ := h.writer.ReadLine()
	// Обработка выбора программы
	return h.workoutUC.StartWorkout(context.Background(), selectedProgram, h.writer.Print)
}
