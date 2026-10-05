package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	apperrors "github.com/rajat-mangla/go-boilerplate/errors"
)

// Interface guard.
var _ http.ResponseWriter = &HTTPResponse{}

// HTTPResponse records the response status and body.
type HTTPResponse struct {
	http.ResponseWriter
	Status        int
	ResponseBytes int64
}

// NewHTTPResponse creates a new response writer that keeps track
// of response body size and status code.
func NewHTTPResponse(w http.ResponseWriter) *HTTPResponse {
	return &HTTPResponse{ResponseWriter: w}
}

// Write transparently proxies to ResponseWriter.Write
// and also records the response bytes length.
func (r *HTTPResponse) Write(p []byte) (int, error) {
	written, err := r.ResponseWriter.Write(p)
	r.ResponseBytes += int64(written)

	return written, err
}

// WriteHeader transparently proxies to ResponseWriter.WriteHeader
// and also records the status code.
func (r *HTTPResponse) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

// WriteJSON writes a struct as a JSON object and sets the json content type header.
func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
	body, err := json.Marshal(v)

	if err != nil {
		msg := fmt.Sprintf("Could not convert the given object to JSON: %v", err)
		WriteText(w, http.StatusInternalServerError, msg)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	_, _ = w.Write(body)
}

// WriteText writes the string to the response.
func WriteText(w http.ResponseWriter, status int, body string) {
	w.WriteHeader(status)
	_, _ = w.Write([]byte(body))
}

// ErrorResponse is the JSON shape written for every error response.
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// WriteError writes a generic error response, for failures that are not a
// apperrors.ServiceError (e.g. unexpected internal errors).
func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, ErrorResponse{Message: message})
}

// WriteServiceError writes an error response derived from a
// apperrors.ServiceError, using its status, code, and message.
func WriteServiceError(w http.ResponseWriter, svcErr apperrors.ServiceError) {
	WriteJSON(w, svcErr.GetResponseStatus(), ErrorResponse{
		Code:    svcErr.GetCode(),
		Message: svcErr.Error(),
	})
}
