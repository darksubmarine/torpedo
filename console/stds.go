package console

import (
	"fmt"
	"os"
)

func ExitWithError(err error) {
	ExitWithErrorCode(err, 1)
}

func ExitIfError(err error) {
	if err != nil {
		ExitWithErrorCode(err, 1)
	}
}

func ExitIfErrors(errs []error) {
	if len(errs) > 0 {
		fmt.Println("[ERROR] Torpedo has terminated with errors:")
		for _, err := range errs {
			fmt.Fprintln(os.Stderr, "  -", err)
		}
		os.Exit(1)
	}
}

func ExitWithErrorCode(err error, code int) {
	fmt.Fprintln(os.Stderr, "ERROR:", err)
	os.Exit(code)
}

func ExitOk(message string) {
	fmt.Println(message)
	os.Exit(0)
}

func Println(a ...any) {
	fmt.Println(a...)
}
