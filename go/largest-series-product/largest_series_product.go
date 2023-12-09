package lsproduct

import (
	"errors"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	// "63915"  span=3 --> "639", "391", "915"
	if len(digits) < span {
		return -1, errors.New("error")
	}
	if span < 0 {
		return -1, errors.New("error")
	}

	for _, c := range digits {
		if !unicode.IsDigit(c) {
			return -1, errors.New("error")
		}
	}
	var p int64
	var res int64
	for i := 0; i < len(digits)-span+1; i++ {
		p = 1
		for j := 0; j < span; j++ {
			p *= int64(digits[i+j] - '0')
		}
		if p > res {
			res = p
		}
	}
	return res, nil
}
