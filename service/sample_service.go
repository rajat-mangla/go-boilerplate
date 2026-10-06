package service

import (
	"github.com/rajat-mangla/go-boilerplate/domain"
	"github.com/rajat-mangla/go-boilerplate/errors"
	"github.com/rs/zerolog/log"
)

type SampleService struct{}

func NewSampleService() *SampleService {
	return &SampleService{}
}

func (s *SampleService) SampleResponse(domain domain.SampleServiceDomain) (string, error) {
	if domain.IsErrorResponse {
		log.Info().Msg("Invalid request parameter provided")
		return "", errors.ErrInvalidRequestParam
	}

	return "This is a sample service method.", nil
}
