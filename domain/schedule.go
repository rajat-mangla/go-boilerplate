package domain

import "time"

func GenerateDueDates(purchaseDate time.Time, installments int) []time.Time {
	dueDates := make([]time.Time, installments)
	for i := 1; i <= installments; i++ {
		targetFirstOfMonth := time.Date(purchaseDate.Year(), purchaseDate.Month(), 1, 0, 0, 0, 0, purchaseDate.Location()).AddDate(0, i, 0)

		anchorDay := purchaseDate.Day()
		lastDay := daysInMonth(targetFirstOfMonth)
		day := min(anchorDay, lastDay)

		dueDates[i-1] = time.Date(targetFirstOfMonth.Year(), targetFirstOfMonth.Month(), day, 0, 0, 0, 0, purchaseDate.Location())
	}
	return dueDates
}

// daysInMonth returns the number of days in the month containing t.
func daysInMonth(t time.Time) int {
	firstOfNextMonth := time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location())
	return firstOfNextMonth.AddDate(0, 0, -1).Day()
}
