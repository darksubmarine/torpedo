package goengine

import (
	"github.com/darksubmarine/torpedo/file"
	"github.com/darksubmarine/torpedo/generator/stack/golang/views/data"
	"github.com/darksubmarine/torpedo/parserx"
	v1 "github.com/darksubmarine/torpedo/parserx/v1"
	"github.com/darksubmarine/torpedo/parserx/vx"
	"os"
	"os/exec"
	"path"
)

type GoEngine struct {

	/* Config */
	options *Options

	/* Runtime */
	entityViews    map[string]*data.EntityView
	idxEntityViews map[string]*data.EntityView
	useCaseViews   map[string]*data.UseCaseView
}

func New(opts *Options) *GoEngine {
	return &GoEngine{
		options:        opts,
		entityViews:    map[string]*data.EntityView{},
		idxEntityViews: map[string]*data.EntityView{},
		useCaseViews:   map[string]*data.UseCaseView{},
	}
}

func (e *GoEngine) Init() error {
	return e.generateDirs()
}

func (e *GoEngine) Fire() []error {

	if err := e.Init(); err != nil {
		return []error{err}
	}

	if errs := e.generateEntityViews(e.options.entityYamlFiles); len(errs) > 0 {
		return errs
	}

	if errs := e.generateUseCaseViews(e.options.usecaseYamlFiles); len(errs) > 0 {
		return errs
	}

	if errs := e.writeCode(); len(errs) > 0 {
		return errs
	}

	if !file.Exists(path.Join(e.options.ProjectDir(), "go.mod")) {
		cmdGoModInit := exec.Command("go", "mod", "init", e.options.Package())
		if _, err := cmdGoModInit.Output(); err != nil {
			return []error{err}
		}

		cmdGoModTidy := exec.Command("go", "mod", "tidy")
		if _, err := cmdGoModTidy.Output(); err != nil {
			return []error{err}
		}
	}

	return nil
}

// generateUseCaseViews generates use case views
func (e *GoEngine) generateUseCaseViews(yamlFiles []string) []error {
	// 1. Generate all views.
	for _, efile := range yamlFiles {
		parser := parserx.New()
		if errs := parser.ParseYaml(efile); len(errs) > 0 {
			return errs
		}

		_, fileName := path.Split(efile)

		switch parser.Kind() {
		case vx.KUseCase:
			uc := parser.Data().(v1.RootUseCase) // TODO refactor to a viewGenerator per version.
			var entitiesView = []data.EntityView{}

			for _, etyPath := range uc.UseCase.Domain.Entities {
				_, fname := path.Split(etyPath)
				entitiesView = append(entitiesView, *e.idxEntityViews[fname])
			}

			// Add docs
			var documentation string
			if uc.UseCase.Doc != "" {
				documentationPath := path.Join(e.options.ProjectDir(), TorpedoDir, TorpedoUseCasesDir, TorpedoDocsDir, uc.UseCase.Doc)
				if _, err := os.Stat(documentationPath); err == nil {
					if docs, err := os.ReadFile(documentationPath); err == nil {
						documentation = string(docs)
					} else {
						return []error{err}
					}
				} else {
					documentation = uc.UseCase.Doc
				}
			}

			e.useCaseViews[fileName] = &data.UseCaseView{
				Package:     e.options.Package(),
				Path:        e.options.UseCasesPath(),
				Name:        uc.UseCase.Name,
				Description: uc.UseCase.Description,
				Doc:         documentation,
				Entities:    entitiesView,
			}
		default:
			return []error{ErrKindNotSupported}
		}

	}

	return nil
}

// generateEntityViews generates entity views and its relationships
func (e *GoEngine) generateEntityViews(yamlFiles []string) []error {
	var viewGenerators = map[string]IEntityViewGenerator{}
	var viewOptions = ViewOptions{AppPackage: e.options.Package(), EntityPath: e.options.EntityPath(), ProjectDir: e.options.ProjectDir()}

	// 1. Generate all views.
	for _, efile := range yamlFiles {
		parser := parserx.New()
		if errs := parser.ParseYaml(efile); len(errs) > 0 {
			return errs
		}

		_, fileName := path.Split(efile)

		switch parser.Kind() {
		case vx.KEntity:
			if viewGenerator, err := NewEntityViewGenerator(parser.Version(), parser.Data(), viewOptions); err == nil {
				viewGenerators[fileName] = viewGenerator
			} else {
				return []error{err}
			}
		default:
			return []error{ErrKindNotSupported}
		}
	}

	// 2. Generate Relationships.
	for _, vg := range viewGenerators {
		if vg.HasRelationships() {
			vg.HydrateRelationships(viewGenerators)
		}
	}

	// 3. Render all files.
	for fileName, vg := range viewGenerators {
		etyView := vg.EntityView()
		e.entityViews[etyView.Name] = etyView
		e.idxEntityViews[fileName] = etyView
	}

	return nil
}

func (e *GoEngine) generateDirs() error {
	projectGenerator := NewProjectGenerator(e.options)
	return projectGenerator.GenerateDirs()
}

func (e *GoEngine) writeCode() []error {

	codeGenerator := NewCodeGenerator(e.options)

	for _, etyv := range e.entityViews {
		errs := codeGenerator.writeEntityCode(etyv)
		if len(errs) > 0 {
			return errs
		}
	}

	for _, useCaseView := range e.useCaseViews {
		errs := codeGenerator.writeUseCaseCode(useCaseView)
		if len(errs) > 0 {
			return errs
		}
	}

	domainView := &data.DomainView{
		Package:  e.options.Package(),
		Path:     e.options.DomainPath(),
		Entities: e.entityViews,
		UseCases: e.useCaseViews,
	}

	if errs := codeGenerator.writeDomainCode(domainView); len(errs) > 0 {
		return errs
	}

	appView := &data.AppView{
		Package:        e.options.Package(),
		Path:           e.options.Package(),
		DependencyPath: e.options.DependencyPath(),
	}

	if errs := codeGenerator.writeAppCode(appView); len(errs) > 0 {
		return errs
	}

	return nil
}

// ------

// Generate inputs code
func (e *GoEngine) generateInputs() error { return nil }

// Generate outputs code
func (e *GoEngine) generateOutputs() error { return nil }

// Generate testing code
func (e *GoEngine) generateTesting() error { return nil }

func (e *GoEngine) generateDocs() error { return nil }
