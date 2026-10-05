package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func date(y int, m time.Month, d int) time.Time {
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

func TestGenerateDueDatesNormalMonths(t *testing.T) {
	purchase := date(2026, time.March, 15)
	got := GenerateDueDates(purchase, 3)
	want := []time.Time{
		date(2026, time.April, 15),
		date(2026, time.May, 15),
		date(2026, time.June, 15),
	}
	for i := range want {
		assert.Truef(t, got[i].Equal(want[i]), "due date[%d] = %v, want %v", i, got[i], want[i])
	}
}

func TestGenerateDueDatesClampsAtMonthEndWithoutDrift(t *testing.T) {
	purchase := date(2026, time.January, 31)
	got := GenerateDueDates(purchase, 4)
	want := []time.Time{
		date(2026, time.February, 28), // not a leap year
		date(2026, time.March, 31),    // does not drift down from Feb's clamp
		date(2026, time.April, 30),
		date(2026, time.May, 31),
	}
	for i := range want {
		assert.Truef(t, got[i].Equal(want[i]), "due date[%d] = %v, want %v", i, got[i], want[i])
	}
}

func TestGenerateDueDatesLeapYearFebruary(t *testing.T) {
	purchase := date(2028, time.January, 31) // 2028 is a leap year
	got := GenerateDueDates(purchase, 1)
	want := date(2028, time.February, 29)
	assert.True(t, got[0].Equal(want))
}
