package menu

import "fmt"

func PrintMenu() {

	fmt.Println("===== ТРЕНАЖЕР УСТНОГО СЧЕТА =====")
	fmt.Println("1. Выбрать уровень сложности")
	fmt.Println("2. Тренировка сложения")
	fmt.Println("3. Тренировка вычитания")
	fmt.Println("4. Тренировка умножения")
	fmt.Println("5. Тренировка деления")
	fmt.Println("6. Случайная тренировка")
	fmt.Println("7. Показать статистику")
	fmt.Println("8. Выход")
	fmt.Print("Выберите пункт: ")
}

func PrintDifficultyMenu() {

	fmt.Println("===== СЛОЖНОСТЬ =====")
	fmt.Println("1. Легкий")
	fmt.Println("2. Средний")
	fmt.Println("3. Сложный")
	fmt.Print("Выберите уровень: ")
}
