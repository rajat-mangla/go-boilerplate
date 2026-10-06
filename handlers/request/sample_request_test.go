package request_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/rajat-mangla/go-boilerplate/domain"
	"github.com/rajat-mangla/go-boilerplate/handlers/request"
	"github.com/stretchr/testify/assert"
)

func TestToSampleDomain(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want domain.SampleServiceDomain
	}{
		{
			name: "isError=true",
			args: args{
				r: &http.Request{
					URL: &url.URL{
						RawQuery: "isError=true",
					},
				},
			},
			want: domain.SampleServiceDomain{
				IsErrorResponse: true,
			},
		},
		{
			name: "isError=false",
			args: args{
				r: &http.Request{
					URL: &url.URL{
						RawQuery: "isError=false",
					},
				},
			},
			want: domain.SampleServiceDomain{
				IsErrorResponse: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := request.ToSampleDomain(tt.args.r)
			assert.Equal(t, tt.want, got)
		})
	}
}
