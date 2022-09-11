package validation

// Validator represents an object that can be validated
type Validator interface {
	Validate() error
}
