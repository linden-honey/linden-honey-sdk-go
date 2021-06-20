package health

import (
	"net/http"
)

type Request struct {
}

type Response struct {
	Status      string `json:"status"`
	Description string `json:"description,omitempty"`
}

func (r Response) StatusCode() int {
	switch r.Status {
	case StatusUp:
		return http.StatusOK
	case StatusDown:
		return http.StatusServiceUnavailable
	default:
		return http.StatusInternalServerError
	}
}
