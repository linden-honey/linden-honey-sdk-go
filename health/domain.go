package health

import (
	"context"
)

// Service is a basic health check service interface.
type Service interface {
	Check(ctx context.Context) error
}
