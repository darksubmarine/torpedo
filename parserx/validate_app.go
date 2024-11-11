package parserx

import (
	"fmt"
	v1 "github.com/darksubmarine/torpedo/parserx/v1"
	"github.com/darksubmarine/torpedo/utils"
)

func validateAppV1(data v1.RootApp, filename string) []error {
	errList := make([]error, 0)

	if utils.EmptyString(data.App.Name) {
		errList = append(errList, fmt.Errorf("%w name at %s", ErrRequiredField, filename))
	}

	if utils.EmptyString(data.App.Description) {
		errList = append(errList, fmt.Errorf("%w description at %s", ErrRequiredField, filename))
	}

	if utils.EmptyString(data.App.Stack.Lang) {
		errList = append(errList, fmt.Errorf("%w stack.lang at %s", ErrRequiredField, filename))
	}

	if utils.EmptyString(data.App.Stack.Package) {
		errList = append(errList, fmt.Errorf("%w stack.package at %s", ErrRequiredField, filename))
	}

	return errList
}
