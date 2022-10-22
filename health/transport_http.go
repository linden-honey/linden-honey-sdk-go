package health

import (
	"net/http"

	sdkhttp "github.com/linden-honey/linden-honey-sdk-go/transport/http"
)

// NewHTTPHandler returns a new instance of http.Handler.
func NewHTTPHandler(svc Service) http.Handler {
	return makeHTTPHandlerFunc(svc)
}

// HTTPResponse represents an HTTP response object.
type HTTPResponse struct {
	Status      string `json:"status"`
	Description string `json:"description,omitempty"`
}

func makeHTTPHandlerFunc(svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := svc.Check(r.Context()); err != nil {
			_ = sdkhttp.EncodeJSONResponse(
				w,
				http.StatusServiceUnavailable,
				HTTPResponse{
					Status:      StatusDown,
					Description: err.Error(),
				},
			)

			return
		}

		_ = sdkhttp.EncodeJSONResponse(
			w,
			http.StatusOK,
			HTTPResponse{
				Status: StatusUp,
			},
		)
	}
}
