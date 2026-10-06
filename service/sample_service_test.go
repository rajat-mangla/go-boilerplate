package service_test

import (
	"testing"

	"github.com/rajat-mangla/go-boilerplate/domain"
	"github.com/rajat-mangla/go-boilerplate/errors"
	"github.com/rajat-mangla/go-boilerplate/service"
	"github.com/stretchr/testify/assert"
)

func TestSampleService_SampleResponse(t *testing.T) {
	tests := []struct {
		name     string
		param    domain.SampleServiceDomain
		wantResp string
		wantErr  error
	}{
		{
			name: "Valid response",
			param: domain.SampleServiceDomain{
				IsErrorResponse: false,
			},
			wantResp: "This is a sample service method.",
			wantErr:  nil,
		},
		{
			name: "Error response",
			param: domain.SampleServiceDomain{
				IsErrorResponse: true,
			},
			wantResp: "",
			wantErr:  errors.ErrInvalidRequestParam,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service.SampleService{}
			got, err := s.SampleResponse(tt.param)
			assert.Equal(t, tt.wantResp, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
