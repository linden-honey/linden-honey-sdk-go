package errors

import "fmt"

// NewFieldRequiredError returns error about missing required field value
func NewFieldRequiredError(key string) error {
	return fmt.Errorf("field '%s' is required", key)
}

// NewFieldInvalidError returns error about invalid field value
func NewFieldInvalidError(key string, value interface{}, rule string) error {
	return fmt.Errorf("fiild '%s' has invalid value '%v' - %s", key, value, rule)
}
