package handlers

import (
	"awesomeProject/mental-math-trainer/input"
	"awesomeProject/mental-math-trainer/mathgame"
	"awesomeProject/mental-math-trainer/stats"
	"fmt"
)

func HandleGame(gameStats *stats.Stats, level int) {

	a, b := mathgame.GenerateNumbers(level)

	operation := mathgame.RandomOperation()

	correctAnswer := mathgame.Calculate(a, b, operation)

	fmt.Printf("%d %s %d = ", a, operation, b)

	userAnswer := input.ReadFloat()

	if mathgame.CheckAnswer(correctAnswer, userAnswer) {
		fmt.Println("Верно!")
		gameStats.AddCorrect()

	} else {
		fmt.Printf("Неверно! Правильный ответ: %.1f\n", correctAnswer)
		gameStats.AddWrong()
	}
}
