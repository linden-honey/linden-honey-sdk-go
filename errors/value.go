package errors

import "fmt"

// NewRequiredValueError returns error about missing required value
func NewRequiredValueError(key string) error {
	return fmt.Errorf("'%s' is required", key)
}

// NewInvalidValueError returns error about invalid value
func NewInvalidValueError(key string, value interface{}, rule string) error {
	return fmt.Errorf("'%s' has invalid value '%v' - %s", key, value, rule)
}
