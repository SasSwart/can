package openapi

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"path/filepath"
	"strings"
)

type Renderer interface {
	sanitiseName(string) string
	sanitiseType(*Schema) string
	getOutputFile(Traversable) string
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

	// Replace " " with _ (" " is not allowed in go func names)
	pathSegments = strings.Split(s, " ")
	nameSegments = make([]string, len(pathSegments))
	for i, segment := range pathSegments {
		nameSegments[i] = caser.String(segment)
	}
	s = strings.Join(nameSegments, "")

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

		nameSegments[i] = segment
	}

	return strings.Join(nameSegments, "")
}

func (g GinRenderer) getOutputFile(t Traversable) string {
	var dir string
	switch t.(type) {
	case *OpenAPI:
		dir = ""
	case *Schema:
		dir = "models"
	}
	return filepath.Join(dir, t.GetName()+".go")
}
