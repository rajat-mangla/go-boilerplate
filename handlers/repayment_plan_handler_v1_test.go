package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rajat-mangla/go-boilerplate/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	successRequestBody = `{"customer_id": "cust_123","commodity_cost": 1000.00,"commodity_category": "electronics","installments": 4,"promo_code": "SAVE10","purchase_date": "2026-01-15"}`
)

func doRequest(t *testing.T, body string) *httptest.ResponseRecorder {
	t.Helper()
	svc := service.NewRepaymentPlanService()
	req := httptest.NewRequest(http.MethodPost, "/internal/v1/repayment-plans", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()
	CreateRepaymentPlanV1(svc)(rec, req)
	return rec
}

func TestCreateRepaymentPlanV1Success(t *testing.T) {
	rec := doRequest(t, successRequestBody)

	require.Equal(t, http.StatusOK, rec.Code, rec.Body.String())

	var resp CreateRepaymentPlanResponse
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))

	assert.NotEmpty(t, resp.RepaymentPlanID)
	assert.Equal(t, 1108.00, resp.TotalPayable)
	assert.Len(t, resp.RepaymentSchedule, 4)
}

func TestCreateRepaymentPlanV1MalformedBody(t *testing.T) {
	rec := doRequest(t, `not json`)
	assertErrorResponse(t, rec, http.StatusBadRequest, "invalid_request_body")
}

func TestCreateRepaymentPlanV1MissingCustomerID(t *testing.T) {
	rec := doRequest(t, `{"commodity_cost": 100, "commodity_category": "electronics", "installments": 4}`)
	assertErrorResponse(t, rec, http.StatusBadRequest, "request_param_missing")
}

func TestCreateRepaymentPlanV1InvalidCost(t *testing.T) {
	rec := doRequest(t, `{"customer_id": "cust_123", "commodity_cost": -5, "commodity_category": "electronics", "installments": 4}`)
	assertErrorResponse(t, rec, http.StatusBadRequest, "invalid_commodity_cost")
}

func TestCreateRepaymentPlanV1InstallmentsOutOfRange(t *testing.T) {
	rec := doRequest(t, `{"customer_id": "cust_123", "commodity_cost": 100, "commodity_category": "electronics", "installments": 1}`)
	assertErrorResponse(t, rec, http.StatusBadRequest, "installment_count_out_of_range")
}

func TestCreateRepaymentPlanV1InvalidPromoCode(t *testing.T) {
	rec := doRequest(t, `{"customer_id": "cust_123", "commodity_cost": 100, "commodity_category": "electronics", "installments": 4, "promo_code": "BOGUS"}`)
	assertErrorResponse(t, rec, http.StatusBadRequest, "invalid_promo_code")
}

func TestCreateRepaymentPlanV1InvalidPurchaseDate(t *testing.T) {
	rec := doRequest(t, `{"customer_id": "cust_123", "commodity_cost": 100, "commodity_category": "electronics", "installments": 4, "purchase_date": "not-a-date"}`)
	assertErrorResponse(t, rec, http.StatusBadRequest, "invalid_request_body")
}

func assertErrorResponse(t *testing.T, rec *httptest.ResponseRecorder, wantStatus int, wantCode string) {
	t.Helper()
	require.Equal(t, wantStatus, rec.Code, rec.Body.String())

	var resp ErrorResponse
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	assert.Equal(t, wantCode, resp.Code)
}
