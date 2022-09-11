package pagination

import (
	"fmt"

	sdkerrors "github.com/linden-honey/linden-honey-sdk-go/errors"
)

// Validate validates a SortBy and returns an error if validation is failed
func (sb SortBy) Validate() error {
	if sb == "" {
		return sdkerrors.ErrEmptyValue
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
		return sdkerrors.NewInvalidValueError("Limit", sdkerrors.ErrNegativeValue)
	}
	if p.Offset < 0 {
		return sdkerrors.NewInvalidValueError("Offset", sdkerrors.ErrNegativeValue)
	}
	if err := p.Sort.Validate(); err != nil {
		return sdkerrors.NewInvalidValueError("Sort", err)
	}

	return nil
}

// Validate validates a Chunk[T] and returns an error if validation is failed
func (c Chunk[T]) Validate() error {
	if err := c.Pageable.Validate(); err != nil {
		return err
	}

	return nil
}
