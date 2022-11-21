package openapi

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

type Renderer interface {
	sanitiseName(string) string
	sanitiseType(*Schema) string
}

var _ Renderer = GinRenderer{}

type GinRenderer struct{}

func (g GinRenderer) sanitiseType(s *Schema) string {
	switch s.Type {
	case "boolean":
		return "bool"
	case "array":
		return "[]" + s.Items.GetName()
	case "integer":
		return "int"
	case "object":
		return "struct"
	default:
		return s.Type
	}
}

func (g GinRenderer) sanitiseName(s string) string {
	caser := cases.Title(language.English)

	// Replace - with _ (- is not allowed in go func names)
	pathSegments := strings.Split(s, "-")
	nameSegments := make([]string, len(pathSegments))
	for i, segment := range pathSegments {
		nameSegments[i] = caser.String(segment)
	}
	s = strings.Join(nameSegments, "_")

	// Replace : with _ (- is not allowed in go func names)
	pathSegments = strings.Split(s, ":")
	nameSegments = make([]string, len(pathSegments))
	for i, segment := range pathSegments {
		nameSegments[i] = caser.String(segment)
	}
	s = strings.Join(nameSegments, "_")

	// Convert from '/' delimited path to Camelcase func names
	pathSegments = strings.Split(s, "/")
	nameSegments = make([]string, len(pathSegments))
	for i, segment := range pathSegments {
		if len(segment) == 0 {
			continue
		}
		if segment[0] == '{' {
			nameSegments[i] = caser.String(segment[1 : len(segment)-1])
			continue
		}

		nameSegments[i] = caser.String(segment)
	}

	return strings.Join(nameSegments, "")
}
