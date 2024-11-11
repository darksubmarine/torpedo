package parserx

import (
	"fmt"
	v1 "github.com/darksubmarine/torpedo/parserx/v1"
	"github.com/darksubmarine/torpedo/utils"
)

func validateUseCaseV1(data v1.RootUseCase, filename string) []error {

	errList := make([]error, 0)

	if utils.EmptyString(data.UseCase.Name) {
		errList = append(errList, fmt.Errorf("%w name at %s", ErrRequiredField, filename))
	}

	if utils.EmptyString(data.UseCase.Description) {
		errList = append(errList, fmt.Errorf("%w description at %s", ErrRequiredField, filename))
	}

	return errList
}
