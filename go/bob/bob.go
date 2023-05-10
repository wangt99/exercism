// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	var ans string
	remark = strings.TrimSpace(remark)

	var isShout = func() bool {
		if strings.IndexFunc(remark, unicode.IsLetter) == -1 {
			return false
		}
		return strings.ToUpper(remark) == remark
	}

	var isQuestion = func() bool {
		return strings.HasSuffix(remark, "?")
	}

	var isEmpty = func() bool {
		return len(remark) == 0
	}

	switch {
	case isQuestion() && isShout():
		ans = "Calm down, I know what I'm doing!"
	case isQuestion():
		ans = "Sure."
	case isShout():
		ans = "Whoa, chill out!"
	case isEmpty():
		ans = "Fine. Be that way!"
	default:
		ans = "Whatever."
	}
	return ans
}
