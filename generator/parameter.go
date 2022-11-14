package generator

import (
	"github.com/sasswart/gin-in-a-can/model"
	"github.com/sasswart/gin-in-a-can/openapi"
	"path/filepath"
	"strings"
)

type Parameter struct {
	model.Model
	In string
}

func NewParameterModel(openAPIFile string, openAPIParameter openapi.Parameter) Parameter {
	m := model.Model{
		Name: openAPIParameter.Name,
	}

	switch openAPIParameter.Schema.Type {
	case "boolean":
		m.Type = "*bool"
		break
	case "array":
		name := strings.ReplaceAll(openAPIParameter.Schema.Items.Ref, filepath.Dir(openAPIFile), "")
		name = strings.ReplaceAll(name, filepath.Ext(openAPIParameter.Schema.Items.Ref), "")

		m.Type = "[]" + FuncName(name)
		break
	case "integer":
		m.Type = "int"
		break
	default:
		m.Type = openAPIParameter.Schema.Type
	}

	parameter := Parameter{
		Model: m,
		In:    openAPIParameter.In,
	}

	return parameter
}
