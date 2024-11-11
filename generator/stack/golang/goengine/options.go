package goengine

import "github.com/darksubmarine/torpedo/parserx/vx"

const (
	TorpedoDir         = ".torpedo"
	TorpedoEntitiesDir = "entities"
	TorpedoUseCasesDir = "use_cases"
	TorpedoDocsDir     = "docs"

	DefaultDependencyDir = "/dependency"
	DefaultDomainPath    = "/domain"
	DefaultEntityPath    = "/domain/entities"
	DefaultUseCasesPath  = "/domain/use_cases"
	DefaultTestingPath   = "/domain/testing"
	DefaultInputsPath    = "/domain/inputs"
	DefaultOutputsPath   = "/domain/outputs"
	DefaultEntityId      = vx.ULID
)

type Options struct {
	appPakage string
	outputDir string

	domainPath     string
	entityPath     string
	usecasesPath   string
	inputsPath     string
	outputsPath    string
	testingPath    string
	dependencyPath string

	entityIdType string

	entityYamlFiles  []string
	usecaseYamlFiles []string
}

func DefaultOptions(output string) *Options {
	return &Options{
		appPakage:      "unknown",
		outputDir:      output,
		domainPath:     DefaultDomainPath,
		entityPath:     DefaultEntityPath,
		usecasesPath:   DefaultUseCasesPath,
		testingPath:    DefaultTestingPath,
		inputsPath:     DefaultInputsPath,
		outputsPath:    DefaultOutputsPath,
		dependencyPath: DefaultDependencyDir,

		entityIdType:    DefaultEntityId,
		entityYamlFiles: nil,
	}
}

func DefaultOptionsForApp(output string, pkg string, entityYamlFiles []string, usecaseYamlFiles []string) *Options {
	return &Options{
		appPakage:      pkg,
		outputDir:      output,
		domainPath:     DefaultDomainPath,
		entityPath:     DefaultEntityPath,
		usecasesPath:   DefaultUseCasesPath,
		testingPath:    DefaultTestingPath,
		inputsPath:     DefaultInputsPath,
		outputsPath:    DefaultOutputsPath,
		dependencyPath: DefaultDependencyDir,

		entityIdType:     DefaultEntityId,
		entityYamlFiles:  entityYamlFiles,
		usecaseYamlFiles: usecaseYamlFiles,
	}
}

func (o *Options) Package() string {
	return o.appPakage
}

func (o *Options) ProjectDir() string {
	return o.outputDir
}

func (o *Options) DomainPath() string {
	return o.domainPath
}

func (o *Options) EntityPath() string {
	return o.entityPath
}

func (o *Options) UseCasesPath() string {
	return o.usecasesPath
}

func (o *Options) TestingPath() string {
	return o.testingPath
}

func (o *Options) DependencyPath() string {
	return o.dependencyPath
}

func (o *Options) InputsPath() string {
	return o.inputsPath
}

func (o *Options) OutputsPath() string {
	return o.outputsPath
}

func (o *Options) EntityIdType() string {
	return o.entityIdType
}
