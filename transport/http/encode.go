package http

import (
	"encoding/json"
	"net/http"
)

// EncodeJSONResponse serializes the response as a JSON to the ResponseWriter with appropriate status code
func EncodeJSONResponse(w http.ResponseWriter, statusCode int, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	if statusCode == http.StatusNoContent {
		return nil
	}

	return json.NewEncoder(w).Encode(response)
}
