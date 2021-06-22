package pagination

import (
	"errors"
	"fmt"

	sdkerrors "github.com/linden-honey/linden-honey-sdk-go/errors"
)

var (
	ErrNegativeValue = errors.New("must be non-negative")
	ErrEmptyValue    = errors.New("must be non-empty")
)

// Validate validates a SortBy and returns an error if validation is failed
func (sb SortBy) Validate() error {
	if sb == "" {
		return ErrEmptyValue
	}

	return nil
}

// Validate validates a SortOrder and returns an error if validation is failed
func (so SortOrder) Validate() error {
	if so < Ascending || so > Descending {
		return fmt.Errorf("should have one of the following values %v, %v, %v", Ascending, Normal, Descending)
	}

	return nil
}

// Validate validates a Sort and returns an error if validation is failed
func (s Sort) Validate() error {
	for sb, so := range s {
		if err := sb.Validate(); err != nil {
			return sdkerrors.NewInvalidValueError("SortBy", err)
		}
		if err := so.Validate(); err != nil {
			return sdkerrors.NewInvalidValueError("SortOrder", err)
		}
	}

	return nil
}

// Validate validates a Pageable and returns an error if validation is failed
func (p Pageable) Validate() error {
	if p.Limit < 0 {
		return sdkerrors.NewInvalidValueError("Limit", ErrNegativeValue)
	}
	if p.Offset < 0 {
		return sdkerrors.NewInvalidValueError("Offset", ErrNegativeValue)
	}
	if err := p.Sort.Validate(); err != nil {
		return sdkerrors.NewInvalidValueError("Sort", err)
	}

	return nil
}
