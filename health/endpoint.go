package health

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(Request)

		if err := svc.Check(ctx); err != nil {
			return Response{
				Status:      StatusDown,
				Description: err.Error(),
			}, nil
		}

		return Response{
			Status: StatusUp,
		}, nil
	}
}
