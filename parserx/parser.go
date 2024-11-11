package parserx

import (
	v1 "github.com/darksubmarine/torpedo/parserx/v1"
	"github.com/darksubmarine/torpedo/parserx/vx"
	"gopkg.in/yaml.v3"
	"os"
)

type Parser struct {
	version vx.V
	kind    vx.K
	data    interface{}
}

func New() *Parser { return &Parser{version: vx.Undefined, kind: vx.KInvalid} }

func (p *Parser) Version() vx.V     { return p.version }
func (p *Parser) Kind() vx.K        { return p.kind }
func (p *Parser) Data() interface{} { return p.data }

func (p *Parser) ParseYaml(filename string) []error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []error{err}
	}

	out := vx.DocHeader{}
	if err := yaml.Unmarshal(data, &out); err != nil {
		return []error{err}
	}

	switch vx.Version(out.Version) {
	case vx.V1:
		p.version = vx.V1
		switch vx.Kind(out.Kind) {
		case vx.KEntity:
			p.kind = vx.KEntity
			var entityData = v1.RootEntity{}
			if err := yaml.Unmarshal(data, &entityData); err != nil {
				return []error{err}
			}
			p.data = entityData

			if errs := validateEntityV1(entityData, filename); len(errs) > 0 {
				return errs
			}

			return nil

		case vx.KUseCase:
			p.kind = vx.KUseCase
			var useCaseData = v1.RootUseCase{}
			if err := yaml.Unmarshal(data, &useCaseData); err != nil {
				return []error{err}
			}

			if errs := validateUseCaseV1(useCaseData, filename); len(errs) > 0 {
				return errs
			}

			p.data = useCaseData
			return nil

		case vx.KApp:
			p.kind = vx.KApp
			var appData = v1.RootApp{}
			if err := yaml.Unmarshal(data, &appData); err != nil {
				return []error{err}
			}

			if errs := validateAppV1(appData, filename); len(errs) > 0 {
				return errs
			}

			p.data = appData
			return nil
		default:
			return []error{ErrMissedKind}
		}
	default:
		return []error{ErrMissedVersion}
	}
}
