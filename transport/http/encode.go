package http

import (
	"encoding/json"
	"net/http"
	"time"
)

// EncodeJSONResponse serializes the response object as a JSON to the ResponseWriter with appropriate status code
func EncodeJSONResponse(w http.ResponseWriter, statusCode int, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	if statusCode == http.StatusNoContent || response == nil {
		return nil
	}

	return json.NewEncoder(w).Encode(response)
}

// ErrorResponse represents the basic error response
type ErrorResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Error     string    `json:"error"`
	Message   string    `json:"message"`
	Code      string    `json:"code,omitempty"`
}

// ErrorWithCode represents an error with code
type ErrorWithCode interface {
	error
	Code() string
}

// EncodeJSONError converts an error to ErrorResponse and serializes it as JSON to the ResponseWriter
// with appropriate status code. If an error value is ErrorWithCode, the code will be used when encoding the error.
func EncodeJSONError(w http.ResponseWriter, statusCode int, err error) error {
	if err == nil {
		return nil
	}

	code := ""
	if err, ok := err.(ErrorWithCode); ok {
		code = err.Code()
	}

	return EncodeJSONResponse(
		w,
		statusCode,
		ErrorResponse{
			Timestamp: time.Now(),
			Error:     http.StatusText(statusCode),
			Message:   err.Error(),
			Code:      code,
		},
	)
}
