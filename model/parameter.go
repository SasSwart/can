package model

import (
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/sanitizer"
	"path/filepath"
	"strings"
)

type Parameter struct {
	In string
	Model
}

func NewParameterModel(openAPIFile string, openAPIParameter openapi.Parameter) Parameter {
	m := Model{
		Name: openAPIParameter.Name,
	}

	switch openAPIParameter.Schema.Type {
	case "boolean":
		m.Type = "*bool"
		break
	case "array":
		name := strings.ReplaceAll(openAPIParameter.Schema.Items.Ref, filepath.Dir(openAPIFile), "")
		name = strings.ReplaceAll(name, filepath.Ext(openAPIParameter.Schema.Items.Ref), "")

		m.Type = "[]" + sanitizer.GoFuncName(name)
		break
	case "integer":
		m.Type = "int"
		break
	default:
		m.Type = openAPIParameter.Schema.Type
	}

	parameterModel := Parameter{
		Model: m,
		In:    openAPIParameter.In,
	}

	return parameterModel
}
