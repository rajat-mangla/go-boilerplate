package handlers

import (
	"net/http"

	"github.com/rajat-mangla/go-boilerplate/handlers/request"
	"github.com/rajat-mangla/go-boilerplate/handlers/response"
	"github.com/rajat-mangla/go-boilerplate/service"
)

func SampleHandlerV1(svc *service.SampleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation for sample handler v1
		resp, err := svc.SampleResponse(request.ToSampleDomain(r))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		WriteJSON(w, http.StatusOK, response.SampleResponse{Message: resp})
	}
}
