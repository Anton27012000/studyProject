package week2

import (
	"fmt"
	"sort"
)

func WordFrequency(words []string) map[string]int {
	freq := make(map[string]int)

	for _, word := range words {
		freq[word]++
	}
	return freq
}

func FindPhone(phoneBook map[string]string, name string) string {

	phone, ok := phoneBook[name]

	if !ok {
		return "Контакт не найден"
	}

	return phone
}

func PrintPhoneBook(phoneBook map[string]string) {

	keys := []string{}

	for key := range phoneBook {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, phoneBook[key])
	}
}

func RemoveProduct(products map[string]int, product string, count int) bool {

	current, ok := products[product]

	if !ok {
		return false
	}

	if current <= count {
		products[product] = 0
		return true
	}

	products[product] -= count

	return true
}

func AverageGrade(grades map[string][]int, student string) float64 {

	scores, ok := grades[student]

	if !ok || len(scores) == 0 {
		return 0
	}

	sum := 0

	for _, score := range scores {
		sum += score
	}

	return float64(sum) / float64(len(scores))
}
