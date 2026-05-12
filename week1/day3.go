package week1

import (
	"fmt"
	"strings"
)

func MaxOfTwo() {

	var a, b int

	fmt.Println("Введите два числа")

	_, err := fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if a > b {

		fmt.Println("Максимум:", a)

	} else if b > a {

		fmt.Println("Максимум:", b)

	} else {

		fmt.Println("Числа равны")
	}
}

func MaxOfThree() {

	var c, d, e int

	fmt.Println("Введи три числа")

	_, err := fmt.Scanln(&c, &d, &e)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if c == d && d == e {
		fmt.Println("Числа равны")
		return
	}

	maxNumber := c

	if d > maxNumber {
		maxNumber = d
	}

	if e > maxNumber {
		maxNumber = e
	}

	fmt.Println("Максимум:", maxNumber)
}

func DiscountCalculator() {

	var sum, discount float64

	fmt.Println("Введите сумму покупок")

	_, err := fmt.Scanln(&sum)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if sum < 1000 {

		discount = 0

	} else if sum < 5000 {

		discount = 0.10

	} else {

		discount = 0.20
	}

	finalSum := sum * (1 - discount)

	fmt.Printf("Итог: %.2f\n", finalSum)
}

func ScoreGrade() {

	var score int
	fmt.Println("Введи бал")
	_, err := fmt.Scanln(&score)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if score < 0 || score > 100 {

		fmt.Println("Ошибка: некорректный балл")

		return
	}

	if score >= 90 {

		fmt.Println("A")

	} else if score >= 75 {

		fmt.Println("B")

	} else if score >= 50 {

		fmt.Println("C")

	} else {

		fmt.Println("F")
	}
}

func SeasonByMonth() {

	var month int
	fmt.Println("Введи номер месяца")
	_, err := fmt.Scanln(&month)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if month < 1 || month > 12 {

		fmt.Println("Ошибка: некорректный месяц")

		return
	}

	if month == 12 || month <= 2 {

		fmt.Println("Зима")

	} else if month <= 5 {

		fmt.Println("Весна")

	} else if month <= 8 {

		fmt.Println("Лето")

	} else {

		fmt.Println("Осень")
	}
}

func LoginValidation() {

	var login string
	fmt.Println("Введи логин")
	_, err := fmt.Scanln(&login)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(login) < 5 {

		fmt.Println("Слишком короткий логин")

		return
	}

	if strings.Contains(login, " ") {

		fmt.Println("Логин не должен содержать пробелы")

		return
	}

	fmt.Println("Логин корректный")
}
