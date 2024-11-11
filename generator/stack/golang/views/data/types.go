package data

import (
	"github.com/darksubmarine/torpedo-lib-go/enum"
	"strings"
)

const undefined = "undefined"

type DataTypeEnum enum.Type

const (
	Undefined DataTypeEnum = iota
	_
	String
	Integer
	Float
	Date
	Boolean
	UUID
	ULID
)

func NewDataTypeEnumFromString(s string) DataTypeEnum {
	switch strings.ToLower(s) {
	case "string":
		return String
	case "integer", "int":
		return Integer
	case "float":
		return Float
	case "date":
		return Date
	case "boolean":
		return Boolean
	case "uuid":
		return UUID
	case "ulid":
		return ULID
	default:
		return Undefined
	}
}

func (c DataTypeEnum) ToInt() int {
	return int(c)
}

func (c DataTypeEnum) Value() enum.Type { return enum.Type(c) }

func (c DataTypeEnum) String() string {
	switch c {
	case Undefined:
		return "undefined"
	case String:
		return "string"
	case Integer:
		return "integer"
	case Float:
		return "float"
	case Date:
		return "date"
	case Boolean:
		return "boolean"
	case UUID:
		return "uuid"
	case ULID:
		return "ulid"
	}

	return undefined
}

// GoTypeFromEnum returns the string representation of Go syntax
func GoTypeFromEnum(dataType DataTypeEnum) string {
	switch dataType {
	case String:
		return "string"
	case Integer:
		return "int64"
	case Float:
		return "float64"
	case Date:
		return "int64"
	case Boolean:
		return "bool"
		// TODO defines UUID and ULID
		//case parser.UUID:
		//	return "uuid"
		//case parser.ULID:
		//	return "ulid"
	}

	return ""
}

/**
Relationship Types
*/

// RelationshipTypeEnum supported relationship types
type RelationshipTypeEnum enum.Type

const (
	_ RelationshipTypeEnum = iota
	Rel
	Urn
)

func NewRelationshipEnumFromString(s string) RelationshipTypeEnum {
	switch strings.ToLower(s) {
	case "$rel":
		return Rel
	case "$urn":
		return Urn
	default:
		return 0
	}
}

func (c RelationshipTypeEnum) ToInt() int {
	return int(c)
}

func (c RelationshipTypeEnum) Value() enum.Type { return enum.Type(c) }

func (c RelationshipTypeEnum) String() string {
	switch c {
	case Rel:
		return "$rel"
	case Urn:
		return "$urn"
	default:
		return undefined
	}
}

// CardinalityTypeEnum supported relationship cardinalities
type CardinalityTypeEnum enum.Type

const (
	_ CardinalityTypeEnum = iota
	HasOne
	HasMany
	BelongsTo
)

func NewCardinalityTypeEnumFromString(s string) CardinalityTypeEnum {
	switch strings.ToLower(s) {
	case "hasMany":
		return HasMany
	case "hasOne":
		return HasOne
	default:
		return 0
	}
}

func (c CardinalityTypeEnum) ToInt() int {
	return int(c)
}

func (c CardinalityTypeEnum) Value() enum.Type { return enum.Type(c) }

func (c CardinalityTypeEnum) String() string {
	switch c {
	case HasOne:
		return "hasOne"
	case HasMany:
		return "hasMany"
	default:
		return undefined
	}
}
