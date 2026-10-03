package handlers

import (
	"net/http"
)

func SampleHandlerV1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation for sample handler v1
		w.WriteHeader(http.StatusOK)
		WriteJSON(w, http.StatusOK, SampleResponse{Message: "Sample Handler V1 Response"})
	}
}

type SampleResponse struct {
	Message string `json:"message"`
}
