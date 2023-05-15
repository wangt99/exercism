// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

const (
	sent1 = "For want of a %s the %s was lost."
	sent2 = "And all for the want of a %s."
)

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	ret := []string{}
	for i := range rhyme {
		if i < len(rhyme)-1 {
			ret = append(ret, fmt.Sprintf(sent1, rhyme[i], rhyme[i+1]))
		} else {
			ret = append(ret, fmt.Sprintf(sent2, rhyme[0]))
		}
	}
	return ret
}
