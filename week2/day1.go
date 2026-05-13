package week2

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func Average(numbers []int) float64 {

	if len(numbers) == 0 {
		return 0
	}

	sum := 0

	for _, value := range numbers {
		sum += value
	}

	return float64(sum) / float64(len(numbers))
}

func Min(numbers []int) int {

	if len(numbers) == 0 {
		return 0
	}

	minNumber := numbers[0]

	for _, value := range numbers {

		if value < minNumber {
			minNumber = value
		}
	}

	return minNumber
}

func Max(numbers []int) int {

	if len(numbers) == 0 {
		return 0
	}

	maxNumber := numbers[0]

	for _, value := range numbers {

		if value > maxNumber {
			maxNumber = value
		}
	}

	return maxNumber
}

func EvenNumbers(numbers []int) []int {
	result := []int{}

	for _, number := range numbers {

		if number%2 == 0 {
			result = append(result, number)
		}
	}

	return result
}

func FindIndex(numbers []int, target int) int {

	for index, value := range numbers {

		if value == target {
			return index
		}
	}

	return -1
}

func Reverse(numbers []int) []int {

	result := []int{}

	for i := len(numbers) - 1; i >= 0; i-- {
		result = append(result, numbers[i])
	}
	return result
}
