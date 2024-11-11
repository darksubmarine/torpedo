package v1

import "github.com/darksubmarine/torpedo/parserx/vx"

type RootApp struct {
	vx.DocHeader `yaml:",inline"`
	App          AppSpec `yaml:"spec"`
}

type AppSpec struct {
	Name        string     `yaml:"name"`
	Description string     `yaml:"description"`
	Stack       StackSpec  `yaml:"stack"`
	Domain      DomainSpec `yaml:"domain"`
}

type StackSpec struct {
	Lang    string `yaml:"lang"`
	Package string `yaml:"package"`
}

type DomainSpec struct {
	Entities []string `yaml:"entities"`
	UseCases []string `yaml:"useCases"`
}
