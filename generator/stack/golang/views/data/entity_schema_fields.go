package data

import "fmt"

type EntitySchemaItem struct {
	Name            string
	Type            DataTypeEnum // Torpedo data type
	Description     string
	LongDescription string

	//Unique    bool // not supported yet
	Encrypted bool

	ReadOnly bool

	Optional *OptionalField

	// Validation
	Validator IItemValidator
}

type OptionalField struct {
	Default interface{}
}

func (e *EntitySchemaItem) IsOptional() bool {
	return e.Optional != nil
}

func (e *EntitySchemaItem) OptionalValue() string {
	switch e.Type {
	case String:
		return fmt.Sprintf("\"%s\"", e.Optional.Default)
	case Integer, Date:
		return fmt.Sprintf("%d", e.Optional.Default)
	case Float:
		return fmt.Sprintf("%v", e.Optional.Default)
	}

	return "nil"
}

func (e *EntitySchemaItem) HasDescription() bool {
	return e.Description != ""
}

func (e *EntitySchemaItem) FieldType() string {
	//if Relationship == e.Type {
	//	//ret, _ := data.WalkAsString("meta", "cardinality")
	//	//return ret
	//	return "<meta-cardinality>" //TODO return cardinality ?
	//} else {
	return GoTypeFromEnum(e.Type)
	//}
}

func (e *EntitySchemaItem) HasValidation() bool {
	return e.Validator != nil
}
