package health

import (
	"context"
)

// Service represents basic health check service
type Service interface {
	Check(ctx context.Context) error
}
