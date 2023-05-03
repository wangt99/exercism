package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var interest_rate float32
	switch {
	case balance < 0:
		interest_rate = 3.213
	case balance < 1000:
		interest_rate = 0.5
	case balance < 5000:
		interest_rate = 1.621
	default:
		interest_rate = 2.475
	}
	return interest_rate
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return float64(InterestRate(balance)) * balance * 0.01
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var years int
	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		years += 1
	}
	return years
}
