package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinorString(t *testing.T) {
	m := Minor(110800)
	assert.Equal(t, "1108.00", m.String())
}

func TestApplyMargin(t *testing.T) {
	cost := Minor(100000)             // SAR 1000.00
	profit := ApplyMargin(cost, 1080) // 10.8%
	assert.Equal(t, Minor(10800), profit)
}

func TestSplitEvenAbsorbsRemainderInLastInstallment(t *testing.T) {
	total := Minor(115000) // SAR 1150.00
	installments := SplitEven(total, 3)

	assert.Len(t, installments, 3)

	var sum Minor
	for _, amt := range installments {
		sum += amt
	}
	assert.Equal(t, total, sum)

	assert.Equal(t, installments[0], installments[1])
	assert.NotEqual(t, installments[0], installments[2], "last installment should absorb the rounding remainder")
}

func TestSplitEvenExactDivision(t *testing.T) {
	total := Minor(110800) // SAR 1108.00
	installments := SplitEven(total, 4)
	for _, amt := range installments {
		assert.Equal(t, Minor(27700), amt)
	}
}
