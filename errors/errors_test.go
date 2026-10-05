package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingParamError(t *testing.T) {
	err := MissingParamError("param_name")
	assert.Equal(t, "ServiceError: request_param_missing", err.Error())
	assert.Equal(t, "request_param_missing", err.GetCode())
	assert.Equal(t, 400, err.GetResponseStatus())
	assert.Nil(t, err.UnWrap())
}

func TestInvalidCommodityCostError(t *testing.T) {
	err := InvalidCommodityCostError(-5)
	assert.Equal(t, "invalid_commodity_cost", err.GetCode())
	assert.Equal(t, 400, err.GetResponseStatus())
	assert.NotEmpty(t, err.Error())
}

func TestInstallmentCountOutOfRangeError(t *testing.T) {
	err := InstallmentCountOutOfRangeError(1, 2, 12)
	assert.Equal(t, "installment_count_out_of_range", err.GetCode())
	assert.Equal(t, 400, err.GetResponseStatus())
}

func TestInvalidPromoCodeError(t *testing.T) {
	err := InvalidPromoCodeError("BOGUS")
	assert.Equal(t, "invalid_promo_code", err.GetCode())
	assert.Equal(t, 400, err.GetResponseStatus())
}

func TestInvalidRequestBodyError(t *testing.T) {
	cause := errors.New("boom")
	err := InvalidRequestBodyError(cause)
	assert.Equal(t, "invalid_request_body", err.GetCode())
	assert.Equal(t, 400, err.GetResponseStatus())
	assert.NotNil(t, err.UnWrap())
}
