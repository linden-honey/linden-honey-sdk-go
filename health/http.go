package health

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	kitendpint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(path string, endpoint kitendpint.Endpoint, logger log.Logger) http.Handler {
	opts := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	router := mux.
		NewRouter().
		StrictSlash(true)

	router.
		Path(path).
		Methods(http.MethodGet).
		Handler(httptransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
			opts...,
		))

	return router
}

func decodeRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return Request{}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(Response)
	httptransport.SetContentType("application/json")(ctx, w)
	if err := httptransport.EncodeJSONResponse(ctx, w, res); err != nil {
		return err
	}

	return nil
}
