package domain

import (
	"strings"

	apperrors "github.com/rajat-mangla/go-boilerplate/errors"
)

const (
	MinMarginHundredths     = 500  // 5.00%
	defaultMarginHundredths = 1200 // 12.00%
)

var categoryMarginHundredths = map[string]int64{
	"gold":   800,  // 8%
	"oil":    1000, // 10%
	"metals": 1200, // 12%
	"wheat":  1500, // 15%
}

func BaseMarginHundredths(category string) int64 {
	key := strings.ToLower(strings.TrimSpace(category))
	if margin, ok := categoryMarginHundredths[key]; ok {
		return margin
	}
	return defaultMarginHundredths
}

func ResolvePromoDiscount(code string) (int64, error) {
	switch code {
	case "":
		return 0, nil
	case "SAVE10":
		return 10, nil
	case "SAVE20":
		return 20, nil
	default:
		return 0, apperrors.InvalidPromoCodeError(code)
	}
}

func ApplyPromoDiscount(marginHundredths, discountPercent int64) int64 {
	reduced := marginHundredths - (marginHundredths*discountPercent)/100
	return max(reduced, MinMarginHundredths)
}
