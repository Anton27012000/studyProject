package tests

import (
	"awesomeProject/task-manager/storage"
	"awesomeProject/task-manager/utils"
	"testing"
)

func TestAddTask(t *testing.T) {

	taskStorage := storage.NewTaskStorage()

	taskStorage.AddTask(
		"Learn Go",
		"high",
	)

	if len(taskStorage.Tasks) != 1 {
		t.Error("task was not added")
	}
}

func TestAddTaskEmptyTitle(t *testing.T) {

	taskStorage := storage.NewTaskStorage()

	taskStorage.AddTask(
		"   ",
		"high",
	)

	if len(taskStorage.Tasks) != 0 {
		t.Error(
			"empty task should not be added",
		)
	}
}

func TestFindByID(t *testing.T) {

	taskStorage := storage.NewTaskStorage()

	taskStorage.AddTask(
		"Learn Go",
		"high",
	)

	task, ok := taskStorage.FindByID(1)

	if !ok {
		t.Fatal("task should exist")
	}

	if task.Title != "Learn go" {
		t.Error("wrong task returned")
	}
}

func TestCompleteTask(t *testing.T) {

	taskStorage := storage.NewTaskStorage()

	taskStorage.AddTask(
		"Learn Go",
		"high",
	)

	taskStorage.CompleteTask(1)

	task, _ := taskStorage.FindByID(1)

	if !task.Done {
		t.Error(
			"task should be completed",
		)
	}
}

func TestStats(t *testing.T) {

	taskStorage := storage.NewTaskStorage()

	taskStorage.AddTask(
		"Task 1",
		"high",
	)

	taskStorage.AddTask(
		"Task 2",
		"low",
	)

	taskStorage.CompleteTask(1)

	stats := taskStorage.GetStats()

	if stats["done"] != 1 {
		t.Error(
			"done count should be 1",
		)
	}

	if stats["active"] != 1 {
		t.Error(
			"active count should be 1",
		)
	}
}

func TestNormalizeTitle(t *testing.T) {

	result := utils.NormalizeTitle(
		"   пРИВЕТ МИР   ",
	)

	expected := "Привет мир"

	if result != expected {

		t.Errorf(
			"expected %s, got %s",
			expected,
			result,
		)
	}
}
