package request

import (
	"net/http"

	"github.com/rajat-mangla/go-boilerplate/domain"
)

func ToSampleDomain(r *http.Request) domain.SampleServiceDomain {
	isError := r.URL.Query().Get("isError")

	return domain.SampleServiceDomain{
		IsErrorResponse: isError == "true",
	}
}
