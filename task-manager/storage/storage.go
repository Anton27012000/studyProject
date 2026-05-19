package storage

import (
	"awesomeProject/task-manager/models"
	"awesomeProject/task-manager/utils"
	"strings"
)

type TaskStorage struct {
	Tasks  []models.Task
	NextID int
}

func NewTaskStorage() *TaskStorage {

	return &TaskStorage{
		Tasks:  []models.Task{},
		NextID: 1,
	}
}

func (s *TaskStorage) AddTask(
	title string,
	priority string,
) {

	title = utils.NormalizeTitle(title)

	if title == "" {
		return
	}

	task := models.Task{
		ID:       s.NextID,
		Title:    title,
		Priority: priority,
		Done:     false,
	}

	s.Tasks = append(s.Tasks, task)

	s.NextID++
}

func (s *TaskStorage) FindByID(
	id int,
) (*models.Task, bool) {

	for i := range s.Tasks {

		if s.Tasks[i].ID == id {
			return &s.Tasks[i], true
		}
	}

	return nil, false
}

func (s *TaskStorage) DeleteByID(id int) bool {

	for i := range s.Tasks {

		if s.Tasks[i].ID == id {

			s.Tasks = append(
				s.Tasks[:i],
				s.Tasks[i+1:]...,
			)

			return true
		}
	}

	return false
}

func (s *TaskStorage) GetActiveTasks() []models.Task {

	result := []models.Task{}

	for _, task := range s.Tasks {

		if !task.Done {
			result = append(result, task)
		}
	}

	return result
}

func (s *TaskStorage) SearchByTitle(
	query string,
) []models.Task {

	query = strings.ToLower(
		strings.TrimSpace(query),
	)

	result := []models.Task{}

	for _, task := range s.Tasks {

		title := strings.ToLower(task.Title)

		if strings.Contains(title, query) {
			result = append(result, task)
		}
	}

	return result
}

func (s *TaskStorage) CompleteTask(
	id int,
) bool {

	task, ok := s.FindByID(id)

	if !ok {
		return false
	}

	task.Done = true

	return true
}

func (s *TaskStorage) GetStats() map[string]int {

	stats := map[string]int{
		"all":    len(s.Tasks),
		"done":   0,
		"active": 0,
		"high":   0,
		"medium": 0,
		"low":    0,
	}

	for _, task := range s.Tasks {

		if task.Done {
			stats["done"]++
		} else {
			stats["active"]++
		}

		switch task.Priority {

		case "high":
			stats["high"]++

		case "medium":
			stats["medium"]++

		case "low":
			stats["low"]++
		}
	}

	return stats
}
