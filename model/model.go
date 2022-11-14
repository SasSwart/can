package model

import (
	"github.com/sasswart/gin-in-a-can/generator"
	"github.com/sasswart/gin-in-a-can/openapi"
	"path/filepath"
	"strings"
)

func newModels(openAPIFile string, apiSpec openapi.OpenAPI) []Model {
	components := make([]Model, 0)
	for ref, schema := range apiSpec.Components.Schemas {
		model := newModel(openAPIFile, schema)

		name := strings.ReplaceAll(ref, filepath.Dir(openAPIFile), "")
		name = strings.ReplaceAll(name, filepath.Ext(ref), "")
		model.Name = generator.FuncName(name)

		components = append(components, model)
	}
	return components
}

type Model struct {
	Name       string
	Properties map[string]Model
	Type       string
	Items      *Model
	MinLength  int
	MaxLength  int
	Pattern    string
	Required   bool
}

func newModel(name string, schema openapi.Schema) Model {
	properties := make(map[string]Model)
	for name, property := range schema.Properties {
		model := newModel(name, property)
		for _, p := range schema.Required {
			if p == name {
				model.Required = true
				break
			}
		}

		properties[name] = model
	}

	model := Model{
		Name:       name,
		Properties: properties,
		MinLength:  schema.MinLength,
		MaxLength:  schema.MaxLength,
		Pattern:    schema.Pattern,
	}

	if schema.Items != nil {
		item := newModel(name+"Item", *schema.Items)
		model.Items = &item
	}

	switch schema.Type {
	case "boolean":
		model.Type = "bool"
		break
	case "array":
		model.Type = "[]" + model.Items.Type
		break
	case "integer":
		model.Type = "int"
		break
	case "object":
		model.Type = model.Name
		break
	default:
		model.Type = schema.Type
	}

	return model
}
