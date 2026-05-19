package utils

import (
	"strings"
	"unicode"
)

func NormalizeTitle(title string) string {

	title = strings.TrimSpace(title)

	if title == "" {
		return ""
	}

	title = strings.ToLower(title)

	runes := []rune(title)

	runes[0] = unicode.ToUpper(runes[0])

	return string(runes)
}
