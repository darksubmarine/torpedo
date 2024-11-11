package data

import "fmt"

type IItemValidator interface {
	ConstructorCode() string
}

type ItemValidatorValue struct {
	Type  DataTypeEnum
	Value interface{}
}

func (e *ItemValidatorValue) ConstructorCode() string {
	switch e.Type {
	case String:
		return fmt.Sprintf("NewValue(\"%s\")", e.Value)
	case Integer, Date:
		return fmt.Sprintf("NewValue(%d)", e.Value)
	case Float:
		return fmt.Sprintf("NewValue(%v)", e.Value)
	}

	return ""
}

type ItemValidatorList struct {
	Type DataTypeEnum
	List []interface{}
}

func (e *ItemValidatorList) ConstructorCode() string {

	var toRet string
	for i, elem := range e.List {
		switch e.Type {
		case String:
			if i == 0 {
				toRet = fmt.Sprintf("\"%v\"", elem)
			} else {
				toRet = fmt.Sprintf("%s,\"%v\"", toRet, elem)
			}
		case Integer, Float, Date:
			if i == 0 {
				toRet = fmt.Sprintf("%v", elem)
			} else {
				toRet = fmt.Sprintf("%s,%v", toRet, elem)
			}
		}
	}

	return fmt.Sprintf("NewList([]%s{%s})", GoTypeFromEnum(e.Type), toRet)
}

type ItemValidatorRegex struct {
	Pattern   string
	GoPattern string
}

func (e *ItemValidatorRegex) ConstructorCode() string {
	if e.GoPattern != "" {
		return fmt.Sprintf("NewRegex(`%s`)", e.GoPattern)
	} else if e.Pattern != "" {
		return fmt.Sprintf("NewRegex(`%s`)", e.Pattern)
	}

	return ""
}

type ItemValidatorRange struct {
	Type DataTypeEnum
	Min  interface{}
	Max  interface{}
}

func (e *ItemValidatorRange) ConstructorCode() string {
	var toRet string
	switch e.Type {
	case String:
		toRet = fmt.Sprintf("NewRange(\"%v\", \"%v\")", e.Min, e.Max)
	case Integer, Float, Date:
		toRet = fmt.Sprintf("NewRange(%v, %v)", e.Min, e.Max)
	}
	return toRet
}
