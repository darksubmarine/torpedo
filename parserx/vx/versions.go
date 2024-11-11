package vx

import "strings"

const (
	V1xId = "torpedo.darksub.io/v1"
	V1ID  = "1"
)

type V int

const (
	_ V = iota
	Undefined
	V1
)

func Version(version string) V {
	if strings.HasPrefix(version, V1ID) || strings.HasPrefix(version, V1xId) {
		return V1
	}

	return Undefined
}
