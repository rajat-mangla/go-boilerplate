package errors

import (
	"fmt"
)

type ServiceError interface {
	error
	GetCode() string
	GetResponseStatus() int
	UnWrap() error
}

type serviceError struct {
	err                error
	code               string
	responseStatusCode int
	metadata           any
}

func (s *serviceError) Error() string {
	if s.err != nil {
		return s.err.Error()
	}
	return fmt.Sprintf("ServiceError: %s", s.code)
}

func (s *serviceError) GetCode() string {
	return s.code
}

func (s *serviceError) GetResponseStatus() int {
	return s.responseStatusCode
}

func (s *serviceError) UnWrap() error {
	return s.err
}

func NewServiceError(code string, responseStatusCode int, err error) ServiceError {
	return &serviceError{
		err:                err,
		code:               code,
		responseStatusCode: responseStatusCode,
	}
}
