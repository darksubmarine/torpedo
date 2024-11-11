package views

import (
	"github.com/darksubmarine/torpedo/generator/stack/golang/views/data"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"reflect"
	"text/template"
)

var FuncMap = template.FuncMap{
	"ToTitle": func(s string) string {
		return cases.Title(language.English, cases.NoLower).String(s)
	},

	"backQuote": func() string { return "`" },
	"hashtag":   func() string { return "#" },

	"incr": func(i int, step int) int { return i + step },

	"toPointerFn": func(kind data.DataTypeEnum) string {
		return ptrFnNameFromEnum(kind)
	},

	"fromInputFieldsMap": func(fieldName string, fmap map[string]string) string {
		if val, ok := fmap[fieldName]; ok {
			return val
		}

		return fieldName
	},

	"isNotNil": func(obj interface{}) bool {
		iv := reflect.ValueOf(obj)
		//if iv.IsValid() {
		//	return true
		//}
		switch iv.Kind() {
		case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface:
			return !iv.IsNil()
		default:
			return false
		}
	},

	"isNotEmpty": func(obj interface{}) bool {
		switch o := obj.(type) {
		case string:
			return o != ""
		case int:
			return true
		}

		return false
	},

	"isHasOne": func(cardinality data.CardinalityTypeEnum) bool {
		if cardinality == data.HasOne {
			return true
		}

		return false
	},
}

// ptrFnNameFromEnum returns the function name to get the pointer to a type value
func ptrFnNameFromEnum(dataType data.DataTypeEnum) string {
	switch dataType {
	case data.String:
		return "String"
	case data.Integer:
		return "Int64"
	case data.Float:
		return "Float64"
	case data.Date:
		return "Int64"
	case data.Boolean:
		return "Bool"
	case data.UUID:
		return "String"
	case data.ULID:
		return "String"
	}

	return ""
}
