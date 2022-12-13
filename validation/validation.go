package validation

// Validator is any type that can be validated.
type Validator interface {
	Validate() error
}
