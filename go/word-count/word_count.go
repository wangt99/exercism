package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	// case insnesitive
	phrase = strings.ToLower(phrase)
	// split
	ff := func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '\''
	}
	words := strings.FieldsFunc(phrase, ff)
	// count word frequency
	wf := Frequency{}
	for _, w := range words {
		if w = strings.Trim(w, "'"); w != "" {
			wf[w] += 1
		}
	}
	return wf
}
