package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	ret := []Record{}
	for _, r := range in {
		if predicate(r) {
			ret = append(ret, r)
		}
	}
	return ret
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	from, to := p.From, p.To
	return func(r Record) bool {
		return r.Day >= from && r.Day <= to
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	recoders := Filter(in, ByDaysPeriod(p))
	var total_amount float64
	for _, r := range recoders {
		total_amount += r.Amount
	}
	return total_amount
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	categoryExpenses := Filter(in, ByCategory(c))
	if len(categoryExpenses) == 0 {
		return 0, errors.New("unknown category")
	}
	return TotalByPeriod(categoryExpenses, p), nil
}
