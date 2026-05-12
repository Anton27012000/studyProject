package week1

import "fmt"

func Factorial() {

	var n int

	fmt.Print("Введите число: ")

	_, err := fmt.Scanln(&n)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result := 1

	for i := 1; i <= n; i++ {

		result *= i
	}

	fmt.Println("Factorial:", result)
}

func SumToN() {

	var number int

	_, err := fmt.Scanln(&number)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	sum := 0
	i := 1

	for i <= number {

		sum += i
		i++
	}

	fmt.Println(sum)
}

func MultiplicationTable() {

	var number int

	fmt.Print("Введите число от 1 до 10: ")

	_, err := fmt.Scanln(&number)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if number < 1 || number > 10 {

		fmt.Println("Ошибка: число должно быть от 1 до 10")

		return
	}

	for i := 1; i <= 10; i++ {

		fmt.Printf(
			"%d * %d = %d\n",
			number,
			i,
			number*i,
		)
	}
}

func EvenOddCounter() {

	even := 0
	odd := 0

	for i := 1; i <= 5; i++ {

		var num int

		_, err := fmt.Scanln(&num)

		if err != nil {
			fmt.Println("Error:", err)
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
}

func GuessNumber() {

	secret := 7

	for {

		var guess int

		fmt.Print("Введите число: ")

		_, err := fmt.Scanln(&guess)

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if guess == secret {

			fmt.Println("Угадал!")

			break
		}

		fmt.Println("Неверно")
	}
}
