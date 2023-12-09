package atbash

import (
	"strings"
	"unicode"
)

// type Cipher map[byte]byte

// var cipherMap Cipher

// func init() {
// 	palin := "abcdefghijklmnopqrstuvwxyz"
// 	cipher := "zyxwvutsrqponmlkjihgfedcba"
// 	for i := 0; i < len(palin); i++ {
// 		fmt.Print(cipher[i])
// 	}

// }

func Atbash(s string) string {
	// normlize
	s = strings.ToLower(s)
	cipher := []rune{}
	var count int
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			count += 1
			if unicode.IsLetter(r) {
				r = 'a' + 'z' - r
			}
			cipher = append(cipher, r)
			if count%5 == 0 {
				cipher = append(cipher, ' ')
			}
		}
	}
	if cipher[len(cipher)-1] == ' ' {
		cipher = cipher[:len(cipher)-1]
	}
	return string(cipher)
}
