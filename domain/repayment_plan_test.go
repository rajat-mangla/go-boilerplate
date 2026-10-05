package domain

import (
	"testing"
	"time"

	apperrors "github.com/rajat-mangla/go-boilerplate/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComputeSuccessPath(t *testing.T) {
	plan, err := Compute(ComputeInput{
		CustomerID:        "cust_123",
		CommodityCost:     1000.00,
		CommodityCategory: "electronics",
		Installments:      4,
		PromoCode:         "SAVE10",
		PurchaseDate:      date(2026, time.January, 15),
	})
	require.NoError(t, err)

	assert.Equal(t, int64(1080), plan.AppliedMarginHundredths)
	assert.Equal(t, Minor(10800), plan.ProfitAmount)
	assert.Equal(t, Minor(110800), plan.TotalPayable)
	assert.Equal(t, Minor(27700), plan.BaseInstallmentAmount)
	require.Len(t, plan.Schedule, 4)
	assert.True(t, plan.Schedule[0].DueDate.Equal(date(2026, time.February, 15)))
	assert.NotEmpty(t, plan.ContractSummary)
}

func TestComputeMonthEndClampingFlowsThrough(t *testing.T) {
	plan, err := Compute(ComputeInput{
		CustomerID:        "cust_123",
		CommodityCost:     100,
		CommodityCategory: "electronics",
		Installments:      2,
		PurchaseDate:      date(2026, time.January, 31),
	})
	require.NoError(t, err)
	require.Len(t, plan.Schedule, 2)
	assert.True(t, plan.Schedule[0].DueDate.Equal(date(2026, time.February, 28)))
	assert.True(t, plan.Schedule[1].DueDate.Equal(date(2026, time.March, 31)))
}

func TestComputeRejectsOutOfRangeInstallments(t *testing.T) {
	for _, n := range []int{0, 1, 13, 100} {
		_, err := Compute(ComputeInput{
			CustomerID:        "cust_123",
			CommodityCost:     100,
			CommodityCategory: "electronics",
			Installments:      n,
			PurchaseDate:      time.Now(),
		})
		assertServiceError(t, err, "installment_count_out_of_range")
	}
}

func TestComputeRejectsInvalidCost(t *testing.T) {
	for _, cost := range []float64{0, -10} {
		_, err := Compute(ComputeInput{
			CustomerID:        "cust_123",
			CommodityCost:     cost,
			CommodityCategory: "electronics",
			Installments:      4,
			PurchaseDate:      time.Now(),
		})
		assertServiceError(t, err, "invalid_commodity_cost")
	}
}

func TestComputeRejectsInvalidPromoCode(t *testing.T) {
	_, err := Compute(ComputeInput{
		CustomerID:        "cust_123",
		CommodityCost:     100,
		CommodityCategory: "electronics",
		Installments:      4,
		PromoCode:         "BOGUS",
		PurchaseDate:      time.Now(),
	})
	assertServiceError(t, err, "invalid_promo_code")
}

func assertServiceError(t *testing.T, err error, wantCode string) {
	t.Helper()
	var svcErr apperrors.ServiceError
	require.ErrorAs(t, err, &svcErr)
	assert.Equal(t, wantCode, svcErr.GetCode())
	assert.Equal(t, 400, svcErr.GetResponseStatus())
}
