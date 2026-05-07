package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	/*
		День 1
		Практика и ДЗ
	*/
	name := "gopher"
	number := 27
	fmt.Println("My first project")
	fmt.Printf("Hello and welcome, %s%d!\n", name, number)

	for i := 1; i <= 5; i++ {
		fmt.Printf("Hello and welcome, %s%d!\n", name, i)
	}

	/*
		День 2
		Практика и ДЗ
	*/

	//калькулятор двух чисел
	var number1, number2, result float64
	var symbol string

	fmt.Print("Enter first number:")
	_, err := fmt.Scanln(&number1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print("Enter operation:")
	_, err = fmt.Scanln(&symbol)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print("Enter second number:")
	_, err = fmt.Scanln(&number2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if (symbol == "/" || symbol == "%") && number2 == 0 {
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
		fmt.Println("Error: unknown operation")
		return
	}

	fmt.Println("Result:", result)

	//перевод градусов
	const (
		scale  = 9.0 / 5.0
		offset = 32.0
	)
	var celsius float64

	fmt.Print("Enter temp (Celsius): ")
	_, err = fmt.Scanln(&celsius)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	function := celsius*scale + offset
	fmt.Printf("Temp: %.2f°F\n", function)

	//минуты в часы и минуты
	var minutes int

	fmt.Print("Enter minutes: ")
	_, err = fmt.Scanln(&minutes)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	hours := minutes / 60
	minutesRemain := minutes % 60

	fmt.Printf("%dh %dmin\n", hours, minutesRemain)

	//площадь круга
	const pi = 3.14159

	var radius float64

	fmt.Print("Enter radius: ")
	_, err = fmt.Scanln(&radius)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	area := pi * radius * radius

	fmt.Printf("Area: %.2f\n", area)

	//проверка введенного возраста
	var age int
	fmt.Print("Введите возраст: ")
	_, err = fmt.Scanln(&age)
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

	//преобразование строки в целочисленное число, с плавающей точкой
	var input string

	fmt.Print("Enter input: ")
	_, err = fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	num, _ := strconv.Atoi(input)

	res := num + 10
	fmt.Printf("Result: %d\n", res)

	numF, _ := strconv.ParseFloat(input, 64)
	resF := numF + 10.0
	fmt.Printf("Result: %.2f\n", resF)

	//преобразование целого числа, с плавающей точкой в строку
	n := 123
	s := strconv.Itoa(n)
	fmt.Printf("Result: %s\n", s)

	n = 100
	s = strconv.FormatInt(int64(n), 10)
	fmt.Printf("Result: %s\n", s)

	f := 23.25
	s = strconv.FormatFloat(f, 'f', 2, 64)
	fmt.Printf("Result: %s\n", s)

	/*
		День 3
		Практика и ДЗ
	*/

	//Максимум из двух чисел
	var a, b int

	fmt.Println("Введи два числа")

	_, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if a > b {
		fmt.Println("Максимум:", a)
	} else if a < b {
		fmt.Println("Максимум:", b)
	} else {
		fmt.Println("Числа равны")
	}

	//Максимум из трех чисел
	var c, d, e int

	fmt.Println("Введи три числа")

	_, err = fmt.Scanln(&c, &d, &e)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if c > d {
		if c > e {
			fmt.Println("Максимум:", c)
		}
	} else if d > c {
		if d > e {
			fmt.Println("Максимум:", d)
		}
	} else if e > c {
		if e > d {
			fmt.Println("Максимум:", e)
		}
	} else {
		fmt.Println("Числа равны")
	}

	//скидка по сумме
	var sum, discount float64

	fmt.Println("Введите сумму покупок")
	_, err = fmt.Scanln(&sum)
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

	//Оценка по баллам
	var score int

	_, err = fmt.Scanln(&score)
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

	//Сезон по номеру месяца
	var month int

	_, err = fmt.Scanln(&month)
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

	//Проверка логина
	var login string

	_, err = fmt.Scanln(&login)
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

	/*
		День 4
		Практика и ДЗ
	*/

	//Сумма от 1 до n
	var numb int

	_, err = fmt.Scanln(&numb)
	if err != nil {
		fmt.Println("Error:", err)
	}

	sum1 := 0
	i := 1
	for i <= numb {
		sum1 += i
		i++
	}

	fmt.Println(sum1)

	//факториал
	var numb1 int

	_, err = fmt.Scanln(&numb1)
	if err != nil {
		fmt.Println("Error:", err)
	}

	factorial := 0

	for i := 1; i < numb1; i++ {
		factorial *= i
	}
	fmt.Println("Factorial:", factorial)

	//таблица умножения
	var numb2 int

	fmt.Print("Введите число от 1 до 10: ")
	_, err = fmt.Scanln(&numb2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if numb2 < 1 || numb2 > 10 {
		fmt.Println("Ошибка: число должно быть от 1 до 10")
		return
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", numb2, i, numb2*i)
	}

	//количество четных и нечетных чисел
	even := 0
	odd := 0

	for i := 1; i <= 5; i++ {
		var num int
		_, err = fmt.Scanln(&num)
		if err != nil {
			return
		}

		if num%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	fmt.Println("Четных:", even)
	fmt.Println("Нечетных:", odd)

	//игра "угадай число"
	secret := 7

	for {
		var guess int

		fmt.Print("Введите число: ")
		_, err = fmt.Scanln(&guess)
		if err != nil {
			fmt.Println("Error:", err)
		}

		if guess == secret {
			fmt.Println("Угадал!")
			break
		}

		fmt.Println("Неверно")
	}
}
