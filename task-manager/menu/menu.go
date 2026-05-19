package menu

import "fmt"

func PrintMenu() {

	fmt.Println("===== TASK MANAGER =====")
	fmt.Println("1. Добавить задачу")
	fmt.Println("2. Показать все задачи")
	fmt.Println("3. Показать активные задачи")
	fmt.Println("4. Отметить выполненной")
	fmt.Println("5. Удалить задачу")
	fmt.Println("6. Поиск по названию")
	fmt.Println("7. Статистика")
	fmt.Println("8. Выход")
	fmt.Print("Выберите пункт: ")
}
