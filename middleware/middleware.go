package middleware

// Middleware represents a chainable behavior modifier.
type Middleware[T any] func(next T) T

// Compose is a helper function for composing middleware.
// The first middleware is treated as the outermost middleware.
// Execution will be done in the declared order.
func Compose[T any](mw ...Middleware[T]) Middleware[T] {
	return func(next T) T {
		for i := len(mw) - 1; i >= 0; i-- {
			next = mw[i](next)
		}

		return next
	}
}
