package http

import (
	"encoding/json"
	"net/http"
	"time"
)

// EncodeJSONResponse serializes the provided response as a JSON to the ResponseWriter
// with appropriate status code.
func EncodeJSONResponse(w http.ResponseWriter, statusCode int, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	if statusCode == http.StatusNoContent || response == nil {
		return nil
	}

	return json.NewEncoder(w).Encode(response)
}

// ErrorResponse is the basic error response.
type ErrorResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Error     string    `json:"error"`
	Message   string    `json:"message"`
}

// EncodeJSONError converts an error to the [ErrorResponse] and serializes it as a JSON to the [http.ResponseWriter]
// with appropriate status code.
func EncodeJSONError(w http.ResponseWriter, statusCode int, err error) error {
	if err == nil {
		return nil
	}

	return EncodeJSONResponse(
		w,
		statusCode,
		ErrorResponse{
			Timestamp: time.Now(),
			Error:     http.StatusText(statusCode),
			Message:   err.Error(),
		},
	)
}
