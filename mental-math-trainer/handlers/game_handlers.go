package handlers

import (
	"awesomeProject/mental-math-trainer/input"
	"awesomeProject/mental-math-trainer/mathgame"
	"awesomeProject/mental-math-trainer/stats"
	"fmt"
)

func HandleGame(gameStats *stats.Stats, level int, operation mathgame.Operation) {

	a, b := mathgame.GenerateNumbers(level)

	op := mathgame.ResolveOperation(operation)

	correctAnswer := mathgame.Calculate(a, b, op)

	fmt.Printf("%d %s %d = ", a, op, b)

	userAnswer := input.ReadFloat()

	if mathgame.CheckAnswer(correctAnswer, userAnswer) {
		fmt.Println("Верно!")
		gameStats.AddCorrect()

	} else {
		fmt.Printf("Неверно! Правильный ответ: %.1f\n", correctAnswer)
		gameStats.AddWrong()
	}
}
