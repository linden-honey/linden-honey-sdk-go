package health

import (
	"context"
)

type NoopService struct {
}

func NewNoopService() *NoopService {
	return &NoopService{}
}

func (n *NoopService) Check(_ context.Context) error {
	return nil
}
