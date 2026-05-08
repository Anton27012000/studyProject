package mathgame

import (
	"math"
	"math/rand"
)

type Operation string

const (
	Add Operation = "+"
	Sub Operation = "-"
	Mul Operation = "*"
	Div Operation = "/"
)

func RandomOperation() Operation {

	operations := []Operation{
		Add,
		Sub,
		Mul,
		Div,
	}

	return operations[rand.Intn(len(operations))]
}

func GenerateNumbers(level int) (int, int) {

	switch level {

	case 1:
		return rand.Intn(10) + 1,
			rand.Intn(10) + 1

	case 2:
		return rand.Intn(100) + 1,
			rand.Intn(100) + 1

	case 3:
		return rand.Intn(1000) + 1,
			rand.Intn(1000) + 1

	default:
		return rand.Intn(10) + 1,
			rand.Intn(10) + 1
	}
}

func RoundToOne(value float64) float64 {
	return math.Round(value*10) / 10
}

func Calculate(a int, b int, op Operation) float64 {

	switch op {

	case Add:
		return float64(a + b)

	case Sub:
		return float64(a - b)

	case Mul:
		return float64(a * b)

	case Div:

		if b == 0 {
			return 0
		}

		return float64(a) / float64(b)

	default:
		return 0
	}
}

func CheckAnswer(correct float64, user float64) bool {
	return RoundToOne(correct) == RoundToOne(user)
}
