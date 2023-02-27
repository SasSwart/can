package golang

import (
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/render"
	"github.com/sasswart/gin-in-a-can/tree"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

var _ render.Renderer = Renderer{}

type Renderer struct{}

// SanitiseType sanitizes the prepares the contents of the Type field of a schema for use by the renderer
func (r Renderer) SanitiseType(n tree.NodeTraverser) string {
	if s, ok := n.(*schema.Schema); ok {
		switch s.Type {
		case "boolean":
			return "bool"
		case "array":
			return "[]" + strings.Join(s.GetChildren()["0"].(*schema.Schema).GetName(), "")
		case "integer":
			return "int"
		case "object":
			return "struct"
		default:
			return s.Type
		}
	}
	return ""
}

func (r Renderer) GetOutputFilename(n tree.NodeTraverser) string {
	switch n.(type) {
	case *schema.Schema:
		return r.SanitiseName([]string{"models/"}) + strings.Join(n.GetName(), "")
	default:
		return r.SanitiseName(n.GetName()) + ".go"
	}
}

// SanitiseName should consume the result of an NodeTraverser's .GetName() function.
// It creates a string array that is compliant to go function name restrictions and
// joins the result before returning a single string.
func (r Renderer) SanitiseName(s []string) string {
	caser := cases.Title(language.English)
	var temp []string
	for _, w := range s {
		temp = append(temp, caser.String(cleanFunctionString(w)))
	}
	return strings.Join(temp, "")
}

// cleanFunctionString strips a string of any leading non-alphabetical chars, and all non-alphabetical and non-numerical
// characters that follow.
func cleanFunctionString(s string) (ret string) {
	for i, char := range []rune(s) {
		if i == 0 {
			if ('A' <= char && char <= 'Z') || ('a' <= char && char <= 'z') {
				ret += string(char)
			}
			continue
		}
		if ('A' <= char && char <= 'Z') || ('a' <= char && char <= 'z') || ('0' <= char && char <= '9') {
			ret += string(char)
		}
	}
	return ret
}
