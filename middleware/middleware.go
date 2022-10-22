package middleware

// Middleware represents a chainable behavior modifier.
type Middleware[T any] func(next T) T

// Compose is a helper function for composing [Middleware].
// The first [Middleware] is treated as the outermost.
// The execution will be done in the declared order.
func Compose[T any](mw ...Middleware[T]) Middleware[T] {
	return func(next T) T {
		for i := len(mw) - 1; i >= 0; i-- {
			next = mw[i](next)
		}

		return next
	}
}
