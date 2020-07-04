package configoration

import "errors"

var (
	// ErrNil is returned when the selected
	// section or value is nil.
	ErrNil = errors.New("section or value is nil")

	// ErrInvalidType is returned when the
	// selected value is not the requested
	// value type
	ErrInvalidType = errors.New("invalid value type")
)
