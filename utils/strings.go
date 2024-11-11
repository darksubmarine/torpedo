package utils

import "strings"

func EmptyString(s string) bool {

	if strings.Trim(s, " ") == "" {
		return true
	}

	return false
}
