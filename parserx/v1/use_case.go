package v1

import "github.com/darksubmarine/torpedo/parserx/vx"

type RootUseCase struct {
	vx.DocHeader `yaml:",inline"`
	UseCase      UseCaseSpec `yaml:"spec"`
}

type UseCaseSpec struct {
	Name        string              `yaml:"name"`
	Description string              `yaml:"description"`
	Doc         string              `yaml:"doc"`
	Domain      UseCaseDomainSpec   `yaml:"domain"`
	Actions     []UseCaseActionSpec `yaml:"actions"`
}

type UseCaseDomainSpec struct {
	Entities []string `yaml:"entities"`
}

type UseCaseActionSpec struct {
	Method string         `yaml:"method"`
	Params []UseCaseParam `yaml:"params"`
	Input  []UseCaseInput `yaml:"input"`
}

type UseCaseParam struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`
	Description string `yaml:"description"`
}

type UseCaseInput struct {
	Type string `yaml:"type"`
	Dto  bool   `yaml:"dto"`
}
