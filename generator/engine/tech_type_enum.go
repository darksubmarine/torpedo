package engine

import (
	"github.com/darksubmarine/torpedo-lib-go/enum"
)

type TechType enum.Type

const (
	_ TechType = iota
	Invalid
	Go
	Java
	NodeJs
	Python
	PHP

	StrInvalid = "invalid"
	StrGo      = "go"
	StrJava    = "java"
	StrNodeJs  = "nodejs"
	StrPython  = "python"
	StrPHP     = "php"
)

func (e TechType) Value() enum.Type {
	return enum.Type(e)
}

func (e TechType) String() string {
	switch e {
	case Go:
		return StrGo
	case Java:
		return StrJava
	case NodeJs:
		return StrNodeJs
	case Python:
		return StrPython
	case PHP:
		return StrPHP
	case Invalid:
		return StrInvalid
	default:
		return StrInvalid
	}
}
