package week2

import "fmt"

type Student struct {
	ID    int
	Name  string
	Score float64
}

func FindStudentByID(students []Student, id int) *Student {

	for i := range students {
		if students[i].ID == id {
			return &students[i]
		}
	}
	return nil
}

func AverageScore(students []Student) float64 {

	if len(students) == 0 {
		return 0
	}

	sum := 0.0

	for _, student := range students {
		sum += student.Score
	}

	return sum / float64(len(students))
}

type Product struct {
	ID      int
	Name    string
	Price   float64
	InStock bool
}

func FilterAvailable(products []Product) []Product {
	result := []Product{}

	for _, product := range products {

		if product.InStock {
			result = append(result, product)
		}
	}
	return result

}

func AveragePrice(products []Product) float64 {

	if len(products) == 0 {
		return 0
	}

	sum := 0.0

	for _, product := range products {
		sum += product.Price
	}

	return sum / float64(len(products))
}

type Task struct {
	ID       int
	Title    string
	Done     bool
	Priority int
}

func FilterDoneTasks(tasks []Task) []Task {

	result := []Task{}

	for _, task := range tasks {

		if task.Done {
			result = append(result, task)
		}
	}

	return result
}

func PrintTasks(tasks []Task) {

	for _, task := range tasks {
		fmt.Printf("ID: %d | %s | Done: %t\n", task.ID, task.Title, task.Done)
	}
}
