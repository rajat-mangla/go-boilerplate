package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
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
