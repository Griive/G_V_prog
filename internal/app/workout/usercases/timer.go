package usecases

import (
	"context"
	"fmt"
	"time"
)

type TimerUseCase struct{}

func (uc *TimerUseCase) StartTimer(ctx context.Context, seconds int, updateFunc func(string)) error {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for i := seconds; i > 0; i-- {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			mins, secs := i/60, i%60
			progress := fmt.Sprintf("\rОтдых: %02d:%02d [%-20s]", mins, secs, progressBar(i, seconds))
			updateFunc(progress)
		}
	}
	updateFunc("\nПродолжаем тренировку!")
	return nil
}

func progressBar(current, total int) string {
	width := 20
	filled := int(float64(current) / float64(total) * float64(width))
	return repeat("█", filled) + repeat(" ", width-filled)
}

func repeat(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}
