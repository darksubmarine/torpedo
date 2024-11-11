package goengine

import "errors"

var (

	// ErrInvalidEntityProvidedVersion the provided entity mismatch with the version
	ErrInvalidEntityProvidedVersion = errors.New("the provided entity mismatch with the version")

	// ErrEntityProvidedVersionNotSupported the provided version is not supported
	ErrEntityProvidedVersionNotSupported = errors.New("the provided version is not supported")

	// ErrKindNotSupported the provided kind is not supported
	ErrKindNotSupported = errors.New("the provided kind is not supported")
)
