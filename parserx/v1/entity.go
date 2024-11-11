package v1

import (
	"github.com/darksubmarine/torpedo/parserx/vx"
)

type RootEntity struct {
	vx.DocHeader `yaml:",inline"`
	Entity       EntitySpec `yaml:"spec"`
}

type EntitySpec struct {
	Name          string         `yaml:"name"`
	Plural        string         `yaml:"plural"`
	Description   string         `yaml:"description"`
	Doc           string         `yaml:"doc"`
	Schema        EntitySchema   `yaml:"schema"`
	Relationships []EntityRel    `yaml:"relationships"`
	Adapters      EntityAdapters `yaml:"adapters"`
}

type EntitySchema struct {
	Reserved EntitySchemaReserved `yaml:"reserved"`
	Fields   []EntitySchemaItem   `yaml:"fields"`
}

type EntitySchemaReserved struct {
	Id EntitySchemaReservedId `yaml:"id"`
}

type EntitySchemaReservedId struct {
	Type string `yaml:"type"`
}

type EntitySchemaItem struct {
	Name        string        `yaml:"name"`
	Type        string        `yaml:"type"`
	Description string        `yaml:"description"`
	Doc         string        `yaml:"doc"`
	Encrypted   bool          `yaml:"encrypted"`
	ReadOnly    bool          `yaml:"readonly"`
	Optional    *OptionalItem `yaml:"optional,omitempty"`
	Validate    *ValidateItem `yaml:"validate,omitempty"`
}

type OptionalItem struct {
	Default interface{} `yaml:"default"`
}

type ValidateItem struct {
	List  *ValidateItemList  `yaml:"list,omitempty"`
	Range *ValidateItemRange `yaml:"range,omitempty"`
	Regex *ValidateItemRegex `yaml:"regex,omitempty"`
	Value interface{}        `yaml:"value,omitempty"`
}

type ValidateItemList struct {
	Values []interface{} `yaml:"values"`
}

type ValidateItemRange struct {
	Min interface{} `yaml:"min"`
	Max interface{} `yaml:"max"`
}

type ValidateItemRegex struct {
	Default   string `yaml:"default"`
	GoPattern string `yaml:"go"`
}

type EntityAdapters struct {
	Input  []EntityAdapterItem `yaml:"input"`
	Output []EntityAdapterItem `yaml:"output"`
}

type EntityAdapterItem struct {
	Type     string `yaml:"type"`
	Metadata map[string]interface{}
}

type EntityRel struct {
	Name        string            `yaml:"name"`
	Type        string            `yaml:"type"` // Type is reserved for future implementations ($rel, $urn)
	Ref         string            `yaml:"ref"`
	Cardinality string            `yaml:"cardinality"`
	Load        *EntityRelLoading `yaml:"load"`
}

type EntityRelLoading struct {
	Type     string                 `yaml:"type"`
	Metadata map[string]interface{} `yaml:"metadata"`
}
