package util

// Ptr returns a pointer to the provided value.
func Ptr[T any](v T) *T {
	return &v
}

// Val returns the value that the provided pointer points to or zero value of type T if pointer is nil.
func Val[T any](p *T) T {
	var out T
	if p == nil {
		return out
	}

	return *p
}
