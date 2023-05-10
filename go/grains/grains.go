package grains

import (
	"errors"
	"math"
	"math/big"
)

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("out of bound")
	}
	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() uint64 {
	num, exp := big.NewInt(2), big.NewInt(64)
	return num.Exp(num, exp, nil).Sub(num, big.NewInt(1)).Uint64()
}
