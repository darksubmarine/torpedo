package data

import (
	"fmt"
	"github.com/darksubmarine/torpedo/utils"
)

type UseCaseView struct {
	Package string
	Path    string

	Name        string
	Description string
	Doc         string
	Entities    []EntityView
}

func (v *UseCaseView) HasEntities() bool { return len(v.Entities) > 0 }

func (e *UseCaseView) PackageName() string {
	return utils.ToSnakeCase(e.Name)
}

func (e *UseCaseView) ImportPath() string {
	return fmt.Sprintf("%s%s/%s", e.Package, e.Path, e.PackageName())
}
