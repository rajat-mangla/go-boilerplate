package domain

import (
	"fmt"
)

// Minor represents a SAR amount in halalas (1 SAR = 100 halalas), avoiding
// float rounding issues in money math.
type Minor int64

func (m Minor) SAR() float64 {
	return float64(m) / 100
}

// String formats the amount to two decimal places, e.g. "1108.00".
func (m Minor) String() string {
	return fmt.Sprintf("%.2f", m.SAR())
}

// ApplyMargin computes the profit for cost at the given margin, where
// marginHundredths is the percentage multiplied by 100 (e.g. 10.8% -> 1080).
// Profit is rounded to the nearest halala, half-up.
func ApplyMargin(cost Minor, marginHundredths int64) Minor {
	return Minor((int64(cost)*marginHundredths + 5000) / 10000)
}

// SplitEven divides total into n installments, each floored to the same
// base amount; the last installment absorbs the rounding remainder so the
// sum always equals total exactly.
func SplitEven(total Minor, n int) []Minor {
	installments := make([]Minor, n)
	base := Minor(int64(total) / int64(n))
	for i := 0; i < n-1; i++ {
		installments[i] = base
	}
	installments[n-1] = total - base*Minor(n-1)
	return installments
}
