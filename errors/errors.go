package errors

import "fmt"

var (
	ErrInvalidRequestParam = NewServiceError("invalid_request_param", 400, nil)
)

func MissingParamError(paramName string) ServiceError {
	metadata := map[string]any{
		"param_name": paramName,
	}
	return &serviceError{
		code:               "request_param_missing",
		metadata:           metadata,
		responseStatusCode: 400,
	}
}

func InvalidCommodityCostError(cost float64) ServiceError {
	metadata := map[string]any{
		"commodity_cost": cost,
	}
	return &serviceError{
		err:                fmt.Errorf("commodity cost must be a positive number, got %v", cost),
		code:               "invalid_commodity_cost",
		metadata:           metadata,
		responseStatusCode: 400,
	}
}

func InstallmentCountOutOfRangeError(count, min, max int) ServiceError {
	metadata := map[string]any{
		"installments": count,
		"min":          min,
		"max":          max,
	}
	return &serviceError{
		err:                fmt.Errorf("installments must be between %d and %d, got %d", min, max, count),
		code:               "installment_count_out_of_range",
		metadata:           metadata,
		responseStatusCode: 400,
	}
}

func InvalidPromoCodeError(code string) ServiceError {
	metadata := map[string]any{
		"promo_code": code,
	}
	return &serviceError{
		err:                fmt.Errorf("invalid or expired promo code: %q", code),
		code:               "invalid_promo_code",
		metadata:           metadata,
		responseStatusCode: 400,
	}
}

func InvalidRequestBodyError(err error) ServiceError {
	return &serviceError{
		err:                fmt.Errorf("invalid request body: %w", err),
		code:               "invalid_request_body",
		responseStatusCode: 400,
	}
}
