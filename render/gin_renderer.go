package render

import (
	"github.com/sasswart/gin-in-a-can/errors"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/tree"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

var _ Renderer = GinRenderer{}

type GinRenderer struct{}

func (g GinRenderer) sanitiseType(n tree.NodeTraverser) string {
	// TODO This needs to be specific to the *Schema without needing the package imported
	s, ok := n.(*schema.Schema)
	if !ok {
		errors.CastFail("(g GinRenderer) SanitiseType(s *tree.Node)", "*Node", "*schema.Schema")
	}
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
func (g GinRenderer) getOutputFile(n tree.NodeTraverser) string {
	switch n.(type) {
	case *schema.Schema:
		return g.sanitiseName("models/") + n.GetName()
	//case *root.Root:
	//	return n.GetName() + ".go"
	//case *media.Type:
	//	return n.GetName() + ".go"
	//case *operation.Operation:
	//	return n.GetName() + ".go"
	//case *parameter.Parameter:
	//	return n.GetName() + ".go"
	//case *path.Item:
	//	return n.GetName() + ".go"
	//case *request.Body:
	//	return n.GetName() + ".go"
	//case *response.Response:
	//	return n.GetName() + ".go"
	default:
		return g.sanitiseName(n.GetName()) + ".go"
	}
}
