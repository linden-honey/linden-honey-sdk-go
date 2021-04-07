package errors

import (
	"errors"
	"fmt"
)

var (
	RequiredValueError = errors.New("value is required")
)

// NewInvalidValueError returns error about invalid value
func NewInvalidValueError(key string, err error) error {
	return fmt.Errorf("'%s' has invalid value: %w", key, err)
}

// NewRequiredValueError returns error about missing required value
func NewRequiredValueError(key string) error {
	return NewInvalidValueError(key, RequiredValueError)
}
