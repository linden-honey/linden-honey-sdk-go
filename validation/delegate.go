package validation

import (
	"github.com/linden-honey/linden-honey-sdk-go/errors"
)

// Delegate represents the validator delegation logic
type Delegate struct {
}

// NewDelegate returns a pointer to the new instance of Delegate
func NewDelegate() (*Delegate, error) {
	return &Delegate{}, nil
}

// Validate delegates validation to Validator.Validate method on input value
func (d *Delegate) Validate(in Validator) error {
	if in == nil {
		return errors.NewRequiredValueError("in")
	}

	return in.Validate()
}
