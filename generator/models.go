package generator

import (
	"github.com/sasswart/gin-in-a-can/openapi"
	"path/filepath"
	"strings"
)

func newModels(tc TemplateConfig, apiSpec openapi.OpenAPI) []Model {
	components := make([]Model, 0)
	for ref, schema := range apiSpec.Components.Schemas {
		model := newModel(tc, schema)

		name := strings.ReplaceAll(ref, filepath.Dir(tc.OpenAPIFile), "")
		name = strings.ReplaceAll(name, filepath.Ext(ref), "")
		model.Name = funcName(name)

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

func newModel(tc TemplateConfig, schema openapi.Schema) Model {
	properties := make(map[string]Model)
	for name, property := range schema.Properties {
		model := newModel(tc, property)
		for _, p := range schema.Required {
			if p == name {
				model.Required = true
				break
			}
		}

		properties[name] = model
	}

	s := Model{
		Name:       funcName(schema.Name),
		Properties: properties,
		MinLength:  schema.MinLength,
		MaxLength:  schema.MaxLength,
		Pattern:    schema.Pattern,
	}

	if schema.Items != nil {
		item := newModel(tc, *schema.Items)
		s.Items = &item
	}

	switch schema.Type {
	case "boolean":
		s.Type = "bool"
		break
	case "array":
		name := strings.ReplaceAll(schema.Items.Ref, filepath.Dir(tc.OpenAPIFile), "")
		name = strings.ReplaceAll(name, filepath.Ext(schema.Items.Ref), "")

		s.Type = "[]" + s.Items.Type
		break
	case "integer":
		s.Type = "int"
		break
	case "object":
		s.Type = s.Name
		break
	default:
		s.Type = schema.Type
	}

	return s
}
