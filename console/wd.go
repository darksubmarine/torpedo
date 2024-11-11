package console

import (
	"os"
)

func WorkingDir() string {
	if dir, err := os.Getwd(); err == nil {
		return dir
	}

	return ""
}
