package main

import (
	"fmt"
)

func main() {
	name := "gopher"
	number := 27
	fmt.Printf("Hello and welcome, %s%d!\n", name, number)

	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 100/i)
	}
}
