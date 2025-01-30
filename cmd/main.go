package main

func main() {
	// Инициализация зависимостей
	repo := repositories.NewMemoryProgramRepo()
	workoutUC := usercases.NewWorkoutUseCase(repo)
	console := cli.NewConsole()

	handler := cli.NewCLIHandler(workoutUC, console)

	// Запуск приложения
	if err := handler.Start(); err != nil {
		panic(err)
	}
}
