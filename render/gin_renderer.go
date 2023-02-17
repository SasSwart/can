package render

import (
	"github.com/sasswart/gin-in-a-can/openapi/root"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/tree"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"path/filepath"
	"strings"
)

var _ Renderer = GinRenderer{}

type GinRenderer struct{}

func (g GinRenderer) SanitiseType(s *tree.Node) string {
	// TODO This needs to be specific to the *Schema without needing the package imported
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

func (g GinRenderer) SanitiseName(s string) string {
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

func (g GinRenderer) GetOutputFile(t tree.Node) string {
	var dir string
	switch t.(type) {
	case *root.Root:
		dir = ""
	case *schema.Schema:
		dir = "models"
	}
	name := t.GetName()
	return filepath.Join(dir, name+".go")
}
func (g GinRenderer) SetRenderer(n *tree.Node) {
	if n.GetParent() == nil {
		n.renderer = r
		return
	}
	if g.GetRenderer(n.GetParent()) == nil {
		n.GetParent().SetRenderer(r)
	}
}

func (g GinRenderer) GetRenderer(n *tree.Node) Renderer {
	if n.GetParent() == nil {
		return canRenderer
	}
	return g.GetRenderer(n.GetParent())
}

func (g GinRenderer) GetOutputFile(n *tree.Node) string {
	// TODO this function can do without it's overrides
	return n.GetRenderer().GetOutputFile(n)
}
