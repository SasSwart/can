package generator

import (
	"strings"
)

func Type(schema Model) string {
	switch schema.Type {
	case "boolean":
		return "*bool"
	case "array":
		//if schema.Items == nil {
		//	return "[]interface{}"
		//}
		return "[]" //+ Type(*schema.Items)
	case "integer":
		return "int"
	case "object":
		return schema.Name
	}
	return schema.Type
}

func Sanitize(s string) string {
	noColons := strings.ReplaceAll(s, ":", "_")
	noDashes := strings.ReplaceAll(noColons, "-", "_")
	return noDashes
}
