package main

import (
	"main.go/internal/app/workout/usecases"
	"main.go/internal/infra/repositories"
	"main.go/internal/interfaces/cli"
)

func main() {
	// Инициализация зависимостей
	repo := repositories.NewMemoryProgramRepo()
	workoutUC := usecases.NewWorkoutUseCase(repo)
	console := cli.NewConsole()

	handler := cli.NewCLIHandler(workoutUC, console)

	// Запуск приложения
	if err := handler.Start(); err != nil {
		panic(err)
	}
}
