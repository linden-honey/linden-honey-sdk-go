package health

import (
	"context"
)

type Service interface {
	Check(ctx context.Context) error
}
