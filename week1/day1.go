package week1

import "fmt"

func FirstProject() {

	name := "gopher"
	number := 27

	fmt.Println("My first project")

	fmt.Printf(
		"Hello and welcome, %s%d!\n",
		name,
		number,
	)
}

func PrintHelloLoop() {

	name := "gopher"

	for i := 1; i <= 5; i++ {

		fmt.Printf(
			"Hello and welcome, %s%d!\n",
			name,
			i,
		)
	}
}
