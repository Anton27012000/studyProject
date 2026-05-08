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
