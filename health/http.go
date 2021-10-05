package health

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
)

func NewHTTPHandler(endpoint endpoint.Endpoint, logger log.Logger) http.Handler {
	opts := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	return httptransport.NewServer(
		endpoint,
		decodeRequest,
		encodeResponse,
		opts...,
	)
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
