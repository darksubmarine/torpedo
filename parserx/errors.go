package parserx

import "errors"

var (
	// ErrMissedVersion version is required
	ErrMissedVersion = errors.New("a valid version is required")

	// ErrMissedKind version is required
	ErrMissedKind = errors.New("a valid kind is required")

	// ErrVersionFormat invalid version format, must be a string
	ErrVersionFormat = errors.New("invalid version format, must be a string")

	// ErrRequiredField required field
	ErrRequiredField = errors.New("required field")

	// ErrFieldValue invalid field value
	ErrFieldValue = errors.New("invalid value at")

	// ErrFieldValidator invalid field validator
	ErrFieldValidator = errors.New("invalid validator at")
)
