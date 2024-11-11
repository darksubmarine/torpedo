package goengine

import (
	"github.com/darksubmarine/torpedo/console"
	"os"
	"path"
)

type ProjectGenerator struct {
	opts *Options
}

func NewProjectGenerator(opts *Options) *ProjectGenerator {
	return &ProjectGenerator{opts: opts}
}

func (e *ProjectGenerator) GenerateDirs() error {

	// Dependencies path
	console.Println("mkdir:", path.Join(e.opts.ProjectDir(), e.opts.DependencyPath()))
	if err := e.mkdir(path.Join(e.opts.ProjectDir(), e.opts.DependencyPath())); err != nil {
		return err
	}

	// Domain path
	console.Println("mkdir:", path.Join(e.opts.ProjectDir(), e.opts.DomainPath()))
	if err := e.mkdir(path.Join(e.opts.ProjectDir(), e.opts.DomainPath())); err != nil {
		return err
	}

	// Entities path
	console.Println("mkdir:", path.Join(e.opts.ProjectDir(), e.opts.EntityPath()))
	if err := e.mkdir(path.Join(e.opts.ProjectDir(), e.opts.EntityPath())); err != nil {
		return err
	}

	// Use cases path
	console.Println("mkdir:", path.Join(e.opts.ProjectDir(), e.opts.UseCasesPath()))
	if err := e.mkdir(path.Join(e.opts.ProjectDir(), e.opts.UseCasesPath())); err != nil {
		return err
	}

	// Testing path
	console.Println("mkdir:", path.Join(e.opts.ProjectDir(), e.opts.TestingPath()))
	if err := e.mkdir(path.Join(e.opts.ProjectDir(), e.opts.TestingPath())); err != nil {
		return err
	}

	// Inputs path
	console.Println("mkdir:", path.Join(e.opts.ProjectDir(), e.opts.InputsPath()))
	if err := e.mkdir(path.Join(e.opts.ProjectDir(), e.opts.InputsPath())); err != nil {
		return err
	}

	// Outputs path
	console.Println("mkdir:", path.Join(e.opts.ProjectDir(), e.opts.OutputsPath()))
	if err := e.mkdir(path.Join(e.opts.ProjectDir(), e.opts.OutputsPath())); err != nil {
		return err
	}

	return nil
}

func (e *ProjectGenerator) mkdir(p string) error {
	return os.MkdirAll(p, os.ModePerm)
}
