package data

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

type EntityView struct {
	Meta EntityViewMeta

	Name        string
	PluralName  string
	Description string

	Docs   string
	Schema EntitySchema

	Relationships map[string]EntityRelationship

	Adapters EntityAdapters
}

func (e *EntityView) ImportPath() string {
	return fmt.Sprintf("%s%s/%s", e.Meta.Package, e.Meta.EntityPath, e.PackageName())
}

func (e *EntityView) PackageName() string {
	return strings.ToLower(e.Name)
}

func (e *EntityView) EntityName() string {
	return fmt.Sprintf("%sEntity", cases.Title(language.English).String(e.Name))
}

func (e *EntityView) FetchDocs() string {
	return e.Docs
}

func (e *EntityView) HasOptionalFields() bool {
	for _, field := range e.Schema.Fields {
		if field.Optional != nil {
			return true
		}
	}
	return false
}

func (e *EntityView) HasRelationships() bool {
	return len(e.Relationships) > 0
}

func (e *EntityView) HasRelationshipsBelongsTo() bool {
	for _, r := range e.Relationships {
		if r.Cardinality == BelongsTo {
			return true
		}
	}
	return false
}

func (e *EntityView) FetchRelationshipsBelongsTo() []EntityRelationship {
	ret := []EntityRelationship{}
	for _, r := range e.Relationships {
		if r.Cardinality == BelongsTo {
			ret = append(ret, r)
		}
	}

	return ret
}

func (e *EntityView) HasNestedLoading() bool {
	for _, v := range e.Relationships {
		if v.HasNestedLoading() {
			return true // at least one relationship has nested loading
		}
	}
	return false
}

func (e *EntityView) FetchNestedLoading() []EntityRelationship {
	ret := []EntityRelationship{}
	for _, r := range e.Relationships {
		if r.HasNestedLoading() {
			ret = append(ret, r)
		}
	}

	return ret
}

func (e *EntityView) HasAdapterHTTP() bool {
	return e.Adapters.Input.Http != nil
}
