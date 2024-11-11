package parserx

import (
	"fmt"
	v1 "github.com/darksubmarine/torpedo/parserx/v1"
	"github.com/darksubmarine/torpedo/utils"
)

func validateEntityV1(data v1.RootEntity, filename string) []error {

	errList := make([]error, 0)

	if utils.EmptyString(data.Entity.Name) {
		errList = append(errList, fmt.Errorf("%w 'name' at %s", ErrRequiredField, filename))
	}

	if utils.EmptyString(data.Entity.Plural) {
		errList = append(errList, fmt.Errorf("%w 'plural' at %s", ErrRequiredField, filename))
	}

	if utils.EmptyString(data.Entity.Description) {
		errList = append(errList, fmt.Errorf("%w 'description' at %s", ErrRequiredField, filename))
	}

	if utils.EmptyString(data.Entity.Schema.Reserved.Id.Type) {
		errList = append(errList, fmt.Errorf("%w 'schema.reserved.id.type' at %s", ErrRequiredField, filename))
	} else if data.Entity.Schema.Reserved.Id.Type != "uuid" && data.Entity.Schema.Reserved.Id.Type != "ulid" {
		errList = append(errList, fmt.Errorf("%w 'schema.reserved.id.type' expected: uuid or ulid at file %s", ErrFieldValue, filename))
	}

	if len(data.Entity.Schema.Fields) == 0 {
		errList = append(errList, fmt.Errorf("%w 'schema.fields' at %s", ErrRequiredField, filename))
	} else {
		for i, field := range data.Entity.Schema.Fields {
			if utils.EmptyString(field.Name) {
				errList = append(errList, fmt.Errorf("%w 'schema.fields[%d].name' at %s", ErrRequiredField, i, filename))
			}

			if utils.EmptyString(field.Description) {
				errList = append(errList, fmt.Errorf("%w 'schema.fields[%d].description' at %s", ErrRequiredField, i, filename))
			}

			if utils.EmptyString(field.Type) {
				errList = append(errList, fmt.Errorf("%w 'schema.fields[%d].type' at %s", ErrRequiredField, i, filename))
			} else {
				if field.Type != v1.Integer && field.Type != v1.Float && field.Type != v1.Date && field.Type != v1.Boolean && field.Type != v1.String {
					errList = append(errList, fmt.Errorf("%w 'schema.fields[%d].type' actual: %s expected: %s, %s, %s, %s, %s at %s", ErrFieldValue, i,
						field.Type, v1.Integer, v1.Float, v1.Date, v1.Boolean, v1.String, filename))
				}
			}

			if field.Optional != nil {
				if field.Optional.Default == nil {
					errList = append(errList, fmt.Errorf("%w 'schema.fields[%d].optional.default' at %s", ErrRequiredField, i, filename))
				} else {
					if hasInvalidDataType(field.Optional.Default, field.Type) {
						errList = append(errList, fmt.Errorf("%w 'schema.fields[%d].optional.default' must be %s at %s", ErrFieldValue, i, field.Type, filename))
					}
				}
			}

			if field.Validate != nil {
				if field.Validate.List != nil {
					if len(field.Validate.List.Values) == 0 {
						errList = append(errList, fmt.Errorf("%w 'schema.fields[%d].validate.list.values' at %s", ErrRequiredField, i, filename))
					} else {
						for j, value := range field.Validate.List.Values {
							if hasInvalidDataType(value, field.Type) {
								errList = append(errList, fmt.Errorf("%w 'schema.fields[%d].validate.list.values[%d]' must be %s at %s", ErrFieldValue, i, j, field.Type, filename))
							}
						}
					}
				} else if field.Validate.Value != nil {
					if hasInvalidDataType(field.Validate.Value, field.Type) {
						errList = append(errList, fmt.Errorf("%w 'schema.fields[%d].validate.value' must be %s at %s", ErrFieldValue, i, field.Type, filename))
					}
				} else if field.Validate.Range != nil {

					if field.Type != v1.String && field.Type != v1.Integer && field.Type != v1.Float && field.Type != v1.Date {
						errList = append(errList, fmt.Errorf("%w 'schema.fields[%d].validate.range' only supported for %s, %s, %s and %s data types actual is %s at %s", ErrFieldValidator, i,
							v1.String, v1.Integer, v1.Float, v1.Date, field.Type, filename))
					} else {
						if hasInvalidDataType(field.Validate.Range.Min, field.Type) {
							errList = append(errList, fmt.Errorf("%w 'schema.fields[%d].validate.range.min' must be %s at %s", ErrFieldValue, i, field.Type, filename))
						}

						if hasInvalidDataType(field.Validate.Range.Max, field.Type) {
							errList = append(errList, fmt.Errorf("%w 'schema.fields[%d].validate.range.max' must be %s at %s", ErrFieldValue, i, field.Type, filename))
						}
					}
				} else if field.Validate.Regex != nil {

					if field.Type != v1.String {
						errList = append(errList, fmt.Errorf("%w 'schema.fields[%d].validate.regex' only supported for %s data type actual is %s at %s", ErrFieldValidator, i, v1.String, field.Type, filename))
					}
				}
			}
		}
	}

	if len(data.Entity.Adapters.Input) > 0 {
		for i, adapter := range data.Entity.Adapters.Input {
			if adapter.Type != v1.HTTP {
				errList = append(errList, fmt.Errorf("%w 'adapters.input[%d].type' actual: %s expected: %s at %s", ErrFieldValue, i, adapter.Type, v1.HTTP, filename))
			}
		}
	}

	if len(data.Entity.Adapters.Output) > 0 {
		for i, adapter := range data.Entity.Adapters.Output {
			if adapter.Type != v1.MongoDB && adapter.Type != v1.Memory && adapter.Type != v1.Redis &&
				adapter.Type != v1.SQL && adapter.Type != v1.RedisSQL && adapter.Type != v1.RedisMongoDB {
				errList = append(errList, fmt.Errorf("%w 'adapters.output[%d].type' actual: %s expected: %s,%s,%s,%s,%s,%s at %s", ErrFieldValue, i,
					adapter.Type, v1.Memory, v1.MongoDB, v1.Redis, v1.SQL, v1.RedisSQL, v1.RedisMongoDB, filename))
			}
		}
	}

	return errList
}

func hasInvalidDataType(value interface{}, kind string) bool {
	switch value.(type) {
	case int, int64, int32:
		if kind != v1.Integer && kind != v1.Date {
			return true
		}
	case float32, float64:
		if kind != v1.Float {
			return true
		}
	case bool:
		if kind != v1.Boolean {
			return true
		}
	case string:
		if kind != v1.String {
			return true
		}
	default:
		return true
	}

	return false
}
