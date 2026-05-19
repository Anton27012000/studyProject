package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadString() string {

	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Ошибка чтения")
		return ""
	}

	return strings.TrimSpace(text)
}

func ReadInt() int {

	for {

		input := ReadString()

		number, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Введите число")
			continue
		}

		return number
	}
}
