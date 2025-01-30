// 1
package entities

type Exercise struct {
	Name    string
	Sets    int
	Reps    string
	RestSec int
}

type WorkoutProgram struct {
	Name      string
	Exercises []Exercise
}
