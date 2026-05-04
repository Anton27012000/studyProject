package main

import (
	"fmt"
)

func main() {
	name := "gopher"
	number := 27
	fmt.Println("My first project")
	fmt.Printf("Hello and welcome, %s%d!\n", name, number)

	for i := 1; i <= 5; i++ {
		fmt.Printf("Hello and welcome, %s%d!\n", name, i)
	}

}
