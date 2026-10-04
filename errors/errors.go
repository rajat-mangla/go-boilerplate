package errors

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
