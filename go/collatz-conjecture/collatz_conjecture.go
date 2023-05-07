package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	var steps int
	if n < 1 {
		return steps, errors.New("out of range")
	}
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		steps++
	}
	return steps, nil
}
