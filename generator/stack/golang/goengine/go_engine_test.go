package goengine_test

import (
	"github.com/darksubmarine/torpedo/parserx"
	v1 "github.com/darksubmarine/torpedo/parserx/v1"
	"github.com/darksubmarine/torpedo/parserx/vx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchAppData(t *testing.T) {
	p := parserx.New()
	//workingDir, err := os.Getwd()
	//assert.Nil(t, err)

	assert.Nil(t, p.ParseYaml("_test/.torpedo/app.yaml"))

	assert.Equal(t, p.Kind(), vx.KApp)
	assert.Equal(t, p.Version(), vx.V1)

	data, ok := p.Data().(v1.RootApp)
	assert.True(t, ok)

	assert.Equal(t, data.App.Stack.Lang, v1.Go)

}

func TestGoEngine_Fire(t *testing.T) {
	//TODO(sarrubia) code test
}

func TestGoEngine_Init(t *testing.T) {
	//TODO(sarrubia) code test
}
