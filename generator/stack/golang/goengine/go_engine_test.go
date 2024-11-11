package goengine_test

import (
	"github.com/darksubmarine/torpedo/generator/stack/golang/goengine"
	"github.com/darksubmarine/torpedo/parserx"
	v1 "github.com/darksubmarine/torpedo/parserx/v1"
	"github.com/darksubmarine/torpedo/parserx/vx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchAppData(t *testing.T) {
	p := parserx.New()
	assert.Nil(t, p.ParseYaml("./.torpedo/app.yaml"))

	assert.Equal(t, p.Kind(), vx.KApp)
	assert.Equal(t, p.Version(), vx.V1)

	data, ok := p.Data().(v1.RootApp)
	assert.True(t, ok)

	assert.Equal(t, data.App.Stack.Lang, v1.Go)

}

func TestGoEngine_Fire(t *testing.T) {
	p := parserx.New()
	assert.Nil(t, p.ParseYaml("./.torpedo/app.yaml"))
	data, ok := p.Data().(v1.RootApp)
	assert.True(t, ok)

	opts := goengine.DefaultOptionsForApp(
		"./blog2",
		"bitbucket.org/darksubmarine/torpedo/blog2", data.App.Domain.Entities, nil)
	engine := goengine.New(opts)
	err := engine.Fire()
	assert.Nil(t, err)
}

func TestGoEngine_Init(t *testing.T) {
	opts := goengine.DefaultOptions(
		"./blog2")
	engine := goengine.New(opts)
	assert.Nil(t, engine.Init())
}
