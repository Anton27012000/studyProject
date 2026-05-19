package handlers

import (
	"awesomeProject/task-manager/input"
	"awesomeProject/task-manager/models"
	"awesomeProject/task-manager/storage"
	"fmt"
)

func HandleAddTask(
	taskStorage *storage.TaskStorage,
) {

	fmt.Print("Введите название задачи: ")

	title := input.ReadString()

	if title == "" {
		fmt.Println("Название не может быть пустым")
		return
	}

	fmt.Print(
		"Введите приоритет (low/medium/high): ",
	)

	priority := input.ReadString()

	taskStorage.AddTask(title, priority)

	fmt.Println("Задача добавлена")
}

func HandleShowTasks(tasks []models.Task) {

	if len(tasks) == 0 {

		fmt.Println("Список задач пуст")

		return
	}

	for _, task := range tasks {

		status := "❌"

		if task.Done {
			status = "✅"
		}

		fmt.Printf(
			"ID: %d | %s | Priority: %s | %s\n",
			task.ID,
			task.Title,
			task.Priority,
			status,
		)
	}
}

func HandleCompleteTask(
	taskStorage *storage.TaskStorage,
) {

	fmt.Print("Введите ID задачи: ")

	id := input.ReadInt()

	ok := taskStorage.CompleteTask(id)

	if !ok {
		fmt.Println("Задача не найдена")
		return
	}

	fmt.Println("Задача выполнена")
}

func HandleDeleteTask(
	taskStorage *storage.TaskStorage,
) {

	fmt.Print("Введите ID задачи: ")

	id := input.ReadInt()

	ok := taskStorage.DeleteByID(id)

	if !ok {
		fmt.Println("Задача не найдена")
		return
	}

	fmt.Println("Задача удалена")
}

func HandleSearch(
	taskStorage *storage.TaskStorage,
) {

	fmt.Print("Введите текст поиска: ")

	query := input.ReadString()

	result := taskStorage.SearchByTitle(query)

	HandleShowTasks(result)
}

func HandleStats(
	taskStorage *storage.TaskStorage,
) {

	stats := taskStorage.GetStats()

	fmt.Println("===== СТАТИСТИКА =====")
	fmt.Println("Всего:", stats["all"])
	fmt.Println("Выполнено:", stats["done"])
	fmt.Println("Активных:", stats["active"])
	fmt.Println("High:", stats["high"])
	fmt.Println("Medium:", stats["medium"])
	fmt.Println("Low:", stats["low"])
}
