package raindrops

import "strconv"

func Convert(number int) string {
	res := ""
	ok := false
	if number%3 == 0 {
		res += "Pling"
		ok = true
	}

	if number%5 == 0 {
		res += "Plang"
		ok = true
	}

	if number%7 == 0 {
		res += "Plong"
		ok = true
	}

	if !ok {
		res = strconv.Itoa(number)
	}
	return res
}
