package generator

import (
	"github.com/sasswart/gin-in-a-can/openapi"
	"testing"
)

func TestNewModel(t *testing.T) {
	schema := openapi.Schema{
		Description: "",
		Type:        "",
		Properties: map[string]openapi.Schema{
			"required_field": {},
			"optional_field": {},
		},
		Items:                nil,
		Ref:                  "",
		AdditionalProperties: false,
		MinLength:            0,
		MaxLength:            0,
		Pattern:              "",
		Format:               "",
		Required: []string{
			"required_field",
		},
	}

	model := newModel("", schema)

	if !model.Properties["required_field"].Required {
		t.Log("required_field's Required property is not set to true")
		t.Fail()
	}

	if model.Properties["optional_field"].Required {
		t.Log("optional_field's Required property is set to true")
		t.Fail()
	}
}
