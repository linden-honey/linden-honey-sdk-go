package errors

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyValue        = errors.New("should be non-empty")
	ErrNegativeNumber    = errors.New("should be non-negative number")
	ErrNonPositiveNumber = errors.New("should be positive number")
)

// NewRequiredValueError returns an error about missing required value
func NewRequiredValueError(key string) error {
	return fmt.Errorf("'%s' is required", key)
}

// NewInvalidValueError returns an invalid value error
func NewInvalidValueError(key string, err error) error {
	return fmt.Errorf("'%s' has invalid value: %w", key, err)
}
