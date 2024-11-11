package data

import (
	"fmt"
	"os"
	"strings"
)

type AppView struct {
	Package        string
	Path           string
	DependencyPath string
}

func (v *AppView) PackageName() string {
	parts := strings.Split(v.Path, string(os.PathSeparator))
	return strings.TrimSuffix(strings.TrimPrefix(parts[len(parts)-1], string(os.PathSeparator)), string(os.PathSeparator))
}

func (v *AppView) ImportPath() string {
	return fmt.Sprintf("%s/%s", v.Package, v.PackageName())
}
