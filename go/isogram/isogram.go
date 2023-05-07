package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	phrase := strings.ToLower(word)
	for _, letter := range phrase {
		if !unicode.IsLetter(letter) {
			continue
		}
		if strings.Count(phrase, string(letter)) > 1 {
			return false
		}
	}
	return true
}
