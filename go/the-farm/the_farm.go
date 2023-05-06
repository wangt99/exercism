package thefarm

import "fmt"

var ErrNegativeFodder error = fmt.Errorf("negative fodder")
var ErrNonScale error = fmt.Errorf("non-scale error")
var ErrDivisionByZero error = fmt.Errorf("division by zero")

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (s *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", s.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	weight, err := weightFodder.FodderAmount()

	if cows == 0 {
		return 0, ErrDivisionByZero
	}

	if err != nil {
		if err == ErrScaleMalfunction {
			if weight < 0 {
				return 0, ErrNegativeFodder
			}
			return weight * 2 / float64(cows), nil
		}
		return 0, err
	}

	if weight < 0 {
		return 0, ErrNegativeFodder
	}

	if cows < 0 {
		return 0, &SillyNephewError{cows: cows}
	}

	return weight / float64(cows), nil
}
