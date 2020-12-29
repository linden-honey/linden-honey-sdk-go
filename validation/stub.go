package validation

// ValidatorStub represents the stub implementation of Validator
type ValidatorStub struct {
	Err error
}

// Validate returns an predefined error
func (vs ValidatorStub) Validate() error {
	return vs.Err
}
