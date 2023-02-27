package render

import (
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/tree"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

var _ Renderer = GinRenderer{}

type GinRenderer struct{}

func (g GinRenderer) SanitiseType(n tree.NodeTraverser) string {
	// TODO this logic was moved to schema to allow for rendering to be entirely decoupled from OAS logic
	return ""
}

func (g GinRenderer) SanitiseName(s string) string {
	// TODO make this more elegant.
	// 	Move to generic interface with illegal char list and inject caser dependency
	caser := cases.Title(language.English)

	// Replace - with "" (- is not allowed in go func names)
	pathSegments := strings.Split(s, "-")
	nameSegments := make([]string, len(pathSegments))
	for i, segment := range pathSegments {
		nameSegments[i] = caser.String(segment)
	}
	s = strings.Join(nameSegments, "")

	// Replace : with "" (- is not allowed in go func names)
	pathSegments = strings.Split(s, ":")
	nameSegments = make([]string, len(pathSegments))
	for i, segment := range pathSegments {
		nameSegments[i] = caser.String(segment)
	}
	s = strings.Join(nameSegments, "")

	// Replace " " with "" (" " is not allowed in go func names)
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
func (g GinRenderer) GetOutputFilename(n tree.NodeTraverser) string {
	switch n.(type) {
	case *schema.Schema:
		return g.SanitiseName("models/") + strings.Join(n.GetName(), "")
	default:
		return g.SanitiseName(strings.Join(n.GetName(), "")) + ".go"
	}
}
