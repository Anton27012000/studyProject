package week1

import (
	"fmt"
	"math"
	"strconv"
)

func Calculator() {

	var number1, number2, result float64
	var symbol string

	fmt.Print("Enter first number: ")

	_, err := fmt.Scanln(&number1)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print("Enter operation: ")

	_, err = fmt.Scanln(&symbol)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print("Enter second number: ")

	_, err = fmt.Scanln(&number2)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if (symbol == "/" || symbol == "%") &&
		number2 == 0 {

		fmt.Println("Error: divide by zero")

		return
	}

	if symbol == "+" {

		result = number1 + number2

	} else if symbol == "-" {

		result = number1 - number2

	} else if symbol == "*" {

		result = number1 * number2

	} else if symbol == "/" {

		result = number1 / number2

	} else if symbol == "%" {

		result = math.Mod(number1, number2)

	} else {

		fmt.Println("Unknown operation")

		return
	}

	fmt.Println("Result:", result)
}

func ConvertTemperature() {

	const (
		scale  = 9.0 / 5.0
		offset = 32.0
	)

	var celsius float64

	fmt.Print("Enter temp (Celsius): ")

	_, err := fmt.Scanln(&celsius)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fahrenheit := celsius*scale + offset

	fmt.Printf(
		"Temp: %.2f°F\n",
		fahrenheit,
	)
}

func ParseString() {

	var input string

	fmt.Print("Enter input: ")

	_, err := fmt.Scanln(&input)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	num, err := strconv.Atoi(input)

	if err != nil {

		fmt.Println(
			"Ошибка: некорректное целое число",
		)

	} else {

		fmt.Printf(
			"Result int: %d\n",
			num+10,
		)
	}

	numF, err := strconv.ParseFloat(input, 64)

	if err != nil {

		fmt.Println(
			"Ошибка: некорректное число с плавающей точкой",
		)

	} else {

		fmt.Printf(
			"Result float: %.2f\n",
			numF+10,
		)
	}
}

func MinutesToHours() {

	var minutes int

	fmt.Print("Enter minutes: ")

	_, err := fmt.Scanln(&minutes)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	hours := minutes / 60
	minutesRemain := minutes % 60

	fmt.Printf("%dh %dmin\n", hours, minutesRemain)
}

func CircleArea() {

	const pi = 3.14159

	var radius float64

	fmt.Print("Enter radius: ")

	_, err := fmt.Scanln(&radius)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	area := pi * radius * radius

	fmt.Printf("Area: %.2f\n", area)
}

func CheckAge() {

	var age int

	fmt.Print("Введите возраст: ")

	_, err := fmt.Scanln(&age)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if age < 18 {

		fmt.Println("Minor")

	} else if age <= 65 {

		fmt.Println("Adult")

	} else {

		fmt.Println("Old")
	}
}

func ConvertNumbersToString() {

	n := 123

	s := strconv.Itoa(n)

	fmt.Printf("Result: %s\n", s)

	n = 100

	s = strconv.FormatInt(int64(n), 10)

	fmt.Printf("Result: %s\n", s)

	f := 23.25

	s = strconv.FormatFloat(f, 'f', 2, 64)

	fmt.Printf("Result: %s\n", s)
}
