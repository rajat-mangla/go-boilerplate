package handlers

import (
	"testing"

	apperrors "github.com/rajat-mangla/go-boilerplate/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateRepaymentPlanRequestValidateRejectsMissingCustomerID(t *testing.T) {
	req := CreateRepaymentPlanRequest{CustomerID: ""}
	err := req.Validate()

	var svcErr apperrors.ServiceError
	require.ErrorAs(t, err, &svcErr)
	assert.Equal(t, "request_param_missing", svcErr.GetCode())
}

func TestCreateRepaymentPlanRequestValidateNoPurchaseDate(t *testing.T) {
	req := CreateRepaymentPlanRequest{CustomerID: "cust_123"}
	assert.NoError(t, req.Validate())
	assert.Nil(t, req.ToServiceInput().PurchaseDate)
}

func TestCreateRepaymentPlanRequestValidateValidPurchaseDate(t *testing.T) {
	req := CreateRepaymentPlanRequest{CustomerID: "cust_123", PurchaseDate: "2026-01-31"}
	require.NoError(t, req.Validate())

	input := req.ToServiceInput()
	require.NotNil(t, input.PurchaseDate)
	assert.Equal(t, "2026-01-31", input.PurchaseDate.Format(purchaseDateLayout))
}

func TestCreateRepaymentPlanRequestValidateInvalidPurchaseDate(t *testing.T) {
	req := CreateRepaymentPlanRequest{CustomerID: "cust_123", PurchaseDate: "not-a-date"}
	err := req.Validate()
	require.Error(t, err)
	assert.Nil(t, req.ToServiceInput().PurchaseDate)
}

func TestCreateRepaymentPlanRequestToServiceInputMapsFields(t *testing.T) {
	req := CreateRepaymentPlanRequest{
		CustomerID:        "cust_123",
		CommodityCost:     1000,
		CommodityCategory: "electronics",
		Installments:      4,
		PromoCode:         "SAVE10",
	}
	require.NoError(t, req.Validate())

	input := req.ToServiceInput()
	assert.Equal(t, "cust_123", input.CustomerID)
	assert.Equal(t, 1000.0, input.CommodityCost)
	assert.Equal(t, "electronics", input.CommodityCategory)
	assert.Equal(t, 4, input.Installments)
	assert.Equal(t, "SAVE10", input.PromoCode)
}
