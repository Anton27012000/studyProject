package menu

import "fmt"

func PrintMenu() {

	fmt.Println("===== ТРЕНАЖЕР УСТНОГО СЧЕТА =====")
	fmt.Println("1. Выбрать/изменить уровень сложности")
	fmt.Println("2. Начать игру")
	fmt.Println("3. Показать статистику")
	fmt.Println("4. Выход")
	fmt.Print("Выберите пункт: ")
}

func PrintDifficultyMenu() {

	fmt.Println("===== СЛОЖНОСТЬ =====")
	fmt.Println("1. Легкий")
	fmt.Println("2. Средний")
	fmt.Println("3. Сложный")
	fmt.Print("Выберите уровень: ")
}
