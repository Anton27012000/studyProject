package main

import (
	"awesomeProject/mental-math-trainer/mathgame"
	"fmt"

	"awesomeProject/mental-math-trainer/handlers"
	"awesomeProject/mental-math-trainer/input"
	"awesomeProject/mental-math-trainer/menu"
	"awesomeProject/mental-math-trainer/stats"
)

func selectDifficulty() int {

	menu.PrintDifficultyMenu()

	level := input.ReadInt()

	switch level {

	case 1, 2, 3:
		fmt.Println("Сложность обновлена")
		return level

	default:
		fmt.Println("Неверная сложность")
		return 1
	}
}

func main() {

	gameStats := stats.NewStats()

	for {

		menu.PrintMenu()

		choice := input.ReadInt()

		switch choice {

		case 1:
			gameStats.SetLevel(selectDifficulty())

		case 2:
			handlers.HandleGame(
				gameStats,
				gameStats.CurrentLevel,
				mathgame.Add,
			)

		case 3:
			handlers.HandleGame(
				gameStats,
				gameStats.CurrentLevel,
				mathgame.Sub,
			)

		case 4:
			handlers.HandleGame(
				gameStats,
				gameStats.CurrentLevel,
				mathgame.Mul,
			)

		case 5:
			handlers.HandleGame(
				gameStats,
				gameStats.CurrentLevel,
				mathgame.Div,
			)

		case 6:
			handlers.HandleGame(
				gameStats,
				gameStats.CurrentLevel,
				mathgame.Random,
			)

		case 7:
			gameStats.Print()

		case 8:
			fmt.Println("Выход...")
			return

		default:
			fmt.Println("Ошибка: неизвестный пункт меню")
		}

		fmt.Println()
	}
}
