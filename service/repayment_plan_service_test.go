package service

import (
	"testing"
	"time"

	"github.com/rajat-mangla/go-boilerplate/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateRepaymentPlanDelegatesToDomain(t *testing.T) {
	purchaseDate := time.Date(2026, time.January, 15, 0, 0, 0, 0, time.UTC)
	svc := NewRepaymentPlanService()

	result, err := svc.CreateRepaymentPlan(CreateRepaymentPlanInput{
		CustomerID:        "cust_123",
		CommodityCost:     1000.00,
		CommodityCategory: "electronics",
		Installments:      4,
		PromoCode:         "SAVE10",
		PurchaseDate:      &purchaseDate,
	})
	require.NoError(t, err)
	assert.NotEmpty(t, result.ID)

	wantPlan, err := domain.Compute(domain.ComputeInput{
		CustomerID:        "cust_123",
		CommodityCost:     1000.00,
		CommodityCategory: "electronics",
		Installments:      4,
		PromoCode:         "SAVE10",
		PurchaseDate:      purchaseDate,
	})
	require.NoError(t, err)

	assert.Equal(t, wantPlan.TotalPayable, result.Plan.TotalPayable)
	assert.Equal(t, wantPlan.ContractSummary, result.Plan.ContractSummary)
}
