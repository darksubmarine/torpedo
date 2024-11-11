package views

import (
	"bytes"
	"go/format"
)

func FormatCode(code *bytes.Buffer) ([]byte, error) {
	return format.Source(code.Bytes())
}
