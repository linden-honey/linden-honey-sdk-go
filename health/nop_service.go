package health

import (
	"context"
)

// NopService is no operation health check [Service] implementation.
type NopService struct {
}

// NewNopService returns a pointer to the new instance of [NopService].
func NewNopService() *NopService {
	return &NopService{}
}

// Check checks nothing.
func (n *NopService) Check(_ context.Context) error {
	return nil
}
