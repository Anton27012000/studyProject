package main

import (
	"awesomeProject/task-manager/handlers"
	"awesomeProject/task-manager/input"
	"awesomeProject/task-manager/menu"
	"awesomeProject/task-manager/storage"
	"fmt"
)

func main() {

	taskStorage := storage.NewTaskStorage()

	for {

		menu.PrintMenu()

		choice := input.ReadInt()

		switch choice {

		case 1:
			handlers.HandleAddTask(taskStorage)

		case 2:
			handlers.HandleShowTasks(
				taskStorage.Tasks,
			)

		case 3:

			active := taskStorage.GetActiveTasks()

			handlers.HandleShowTasks(active)

		case 4:
			handlers.HandleCompleteTask(
				taskStorage,
			)

		case 5:
			handlers.HandleDeleteTask(
				taskStorage,
			)

		case 6:
			handlers.HandleSearch(taskStorage)

		case 7:
			handlers.HandleStats(taskStorage)

		case 8:
			fmt.Println("Выход...")
			return

		default:
			fmt.Println("Неизвестный пункт меню")
		}

		fmt.Println()
	}
}
