package data

import (
	"fmt"
	"os"
	"strings"
)

type DomainView struct {
	Package string
	Path    string

	Entities map[string]*EntityView
	UseCases map[string]*UseCaseView
}

func (v *DomainView) PackageName() string {
	parts := strings.Split(v.Path, string(os.PathSeparator))
	return strings.TrimSuffix(strings.TrimPrefix(parts[len(parts)-1], string(os.PathSeparator)), string(os.PathSeparator))
}

func (v *DomainView) Description() string {
	return "domain entry point"
}

func (v *DomainView) ImportPath() string {
	return fmt.Sprintf("%s/%s", v.Package, v.PackageName())
}

func (v *DomainView) HasAdapterHTTP() bool {
	for _, ety := range v.Entities {
		if ety.HasAdapterHTTP() {
			return true
		}
	}

	return false
}
