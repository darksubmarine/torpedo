package file

import (
	"github.com/darksubmarine/torpedo/console"
	"os"
	"path"
)

// Exists checks if a given path exists or not.
func Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

// ReadFile is a wrapper function of os.ReadFile
func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// CreateIfNotExists creates a dir if not exists
func CreateIfNotExists(dir string) error {
	output := path.Join(console.WorkingDir(), dir)
	if !Exists(output) {
		if err := os.MkdirAll(output, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
