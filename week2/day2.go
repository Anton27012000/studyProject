package week2

import "strings"

func NormalizeLogin(login string) string {

	login = strings.TrimSpace(login)
	login = strings.ToLower(login)

	return login
}

func CountWords(text string) int {

	text = strings.TrimSpace(text)

	if text == "" {
		return 0
	}

	words := strings.Fields(text)
	return len(words)
}

func Initials(fullName string) string {

	parts := strings.Fields(fullName)

	if len(parts) < 2 {
		return ""
	}

	first := []rune(parts[0])[0]
	second := []rune(parts[1])[0]

	return string(first) + "." + string(second) + "."
}

func IsPalindrome(text string) bool {

	text = strings.ToLower(text)

	runes := []rune(text)

	left := 0
	right := len(runes) - 1

	for left < right {

		if runes[left] != runes[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func MaskPhone(phone string) string {

	if len(phone) < 7 {
		return phone
	}

	return phone[:5] + "****" + phone[len(phone)-3:]
}
