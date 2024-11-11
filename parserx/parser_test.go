package parserx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser_UseCase(t *testing.T) {
	parser := New()
	err := parser.ParseYaml("_test/usecase_onboarding.yaml")
	assert.Len(t, err, 1)
	assert.ErrorIs(t, err[0], ErrRequiredField)

}

func TestParser_Entity(t *testing.T) {
	parser := New()
	err := parser.ParseYaml("_test/entity_sensor.yaml")
	assert.Len(t, err, 1)
	assert.ErrorIs(t, err[0], ErrFieldValue)

}
