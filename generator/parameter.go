package generator

import (
	"github.com/sasswart/gin-in-a-can/openapi"
	"path/filepath"
	"strings"
)

type Parameter struct {
	Model
	In string
}

func newParameterModel(tc TemplateConfig, openAPIParameter openapi.Parameter) Parameter {
	model := Model{
		Name: funcName(openAPIParameter.Name),
	}

	switch openAPIParameter.Schema.Type {
	case "boolean":
		model.Type = "bool"
		break
	case "array":
		name := strings.ReplaceAll(openAPIParameter.Schema.Items.Ref, filepath.Dir(tc.OpenAPIFile), "")
		name = strings.ReplaceAll(name, filepath.Ext(openAPIParameter.Schema.Items.Ref), "")

		model.Type = "[]" + funcName(name)
		break
	case "integer":
		model.Type = "int"
		break
	default:
		model.Type = openAPIParameter.Schema.Type
	}

	parameter := Parameter{
		Model: model,
		In:    openAPIParameter.In,
	}

	return parameter
}
