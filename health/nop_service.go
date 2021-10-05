package health

import (
	"context"
)

type NopService struct {
}

func NewNopService() *NopService {
	return &NopService{}
}

func (n *NopService) Check(_ context.Context) error {
	return nil
}
