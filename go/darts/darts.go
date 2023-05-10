package darts

import "math"

func Score(x, y float64) int {
	dist := math.Sqrt(x*x + y*y)
	switch {
	case dist > 10:
		return 0
	case dist > 5:
		return 1
	case dist > 1:
		return 5
	default:
		return 10
	}
}
