package input

import "fmt"

func ReadInt() int {

	var n int

	for {

		_, err := fmt.Scanln(&n)

		if err != nil {

			fmt.Println("Ошибка: введите целое число")

			var skip string
			fmt.Scanln(&skip)

			continue
		}

		return n
	}
}

func ReadFloat() float64 {

	var n float64

	for {

		_, err := fmt.Scanln(&n)

		if err != nil {

			fmt.Println("Ошибка: введите число")

			var skip string
			fmt.Scanln(&skip)

			continue
		}

		return n
	}
}
