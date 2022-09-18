package generator

import (
	"strings"

	"github.gom/sasswart/gin-in-a-can/openapi"
)

type Schema struct {
	openapi.Schema
	Name string
}

func Type(schema openapi.Schema) string {
	switch schema.Type {
	case "boolean":
		return "bool"
	case "array":
		return "[]" + Type(*schema.Items)
	case "integer":
		return "int"
	case "object":
		return "struct{}" // TODO: Support nested objects
	}
	return schema.Type
}

func Sanitize(s string) string {
	return strings.ReplaceAll(s, ":", "_")
}
