package main

import (
	"awesomeProject/week2"
	"fmt"
)

func main() {

	students := []week2.Student{
		{0, "Anton", 5.0},
		{1, "Bob", 4.2},
		{2, "Tom", 3.7},
	}
	student := week2.FindStudentByID(students, 0)

	if student != nil {
		fmt.Println(student)
	}
}
