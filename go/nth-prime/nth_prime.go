package prime

import (
	"errors"
)

// 打表法
var primes []int

const MaxNum int = 1000000

func init() {
	isPrime := [MaxNum + 1]bool{}
	for i := 2; i < MaxNum; i++ {
		isPrime[i] = true
	}
	for i := 2; i <= MaxNum; i++ {
		if isPrime[i] {
			for j := 2 * i; j <= MaxNum; j += i {
				isPrime[j] = false
			}
			primes = append(primes, i)
		}
	}
}

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return -1, errors.New("n less than 1")
	}
	// 计算第n个素数
	return primes[n-1], nil
}
