package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseMarginHundredths(t *testing.T) {
	cases := []struct {
		category string
		want     int64
	}{
		{"gold", 800},
		{"GOLD", 800},
		{"  Gold  ", 800},
		{"oil", 1000},
		{"metals", 1200},
		{"wheat", 1500},
		{"electronics", 1200},
		{"unknown-category", 1200},
	}
	for _, c := range cases {
		assert.Equalf(t, c.want, BaseMarginHundredths(c.category), "BaseMarginHundredths(%q)", c.category)
	}
}

func TestResolvePromoDiscount(t *testing.T) {
	cases := []struct {
		code      string
		want      int64
		wantError bool
	}{
		{"", 0, false},
		{"SAVE10", 10, false},
		{"SAVE20", 20, false},
		{"BOGUS", 0, true},
	}
	for _, c := range cases {
		got, err := ResolvePromoDiscount(c.code)
		if c.wantError {
			assert.Errorf(t, err, "ResolvePromoDiscount(%q)", c.code)
			continue
		}
		assert.NoErrorf(t, err, "ResolvePromoDiscount(%q)", c.code)
		assert.Equalf(t, c.want, got, "ResolvePromoDiscount(%q)", c.code)
	}
}

func TestApplyPromoDiscount(t *testing.T) {
	cases := []struct {
		margin, discount, want int64
	}{
		{1200, 10, 1080},
		{1200, 20, 960},
		{1500, 20, 1200},
		{800, 20, 640},
		// synthetic case exercising the 5% floor, which no real
		// category+code combination in the spec actually reaches
		{550, 20, 500},
	}
	for _, c := range cases {
		assert.Equal(t, c.want, ApplyPromoDiscount(c.margin, c.discount))
	}
}
