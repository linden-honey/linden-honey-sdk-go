package validation

// Validator represents the validator interface
type Validator interface {
	Validate() error
}
