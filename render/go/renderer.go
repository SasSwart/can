// Renderer implementations shouldn't have to know about the config package. They should simply plug into the
// pre-configured engine instance created in main.go

package golang

import (
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/render"
	"github.com/sasswart/gin-in-a-can/tree"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

var _ render.Renderer = &Renderer{}

type Renderer struct {
	*render.Base
}

func (g *Renderer) SetTemplateFuncMap(f *template.FuncMap) {
	if f != nil {
		g.Base.SetTemplateFuncMap(f)
		return
	}
	g.Base.TemplateFuncMapping = &template.FuncMap{
		"ToUpper": strings.ToUpper,
		"ToTitle": func(s string) string {
			caser := cases.Title(language.English)
			return caser.String(s)
		},
		"SanitiseName": g.SanitiseName,
		"SanitiseType": g.SanitiseType,
	}
}

func (g *Renderer) GetTemplateFuncMap() template.FuncMap {
	return *g.TemplateFuncMapping
}

// SanitiseType sanitizes the prepares the contents of the Type field of a node for use by the renderer
func (g *Renderer) SanitiseType(n tree.NodeTraverser) string {
	if n == nil {
		return ""
	}
	if s, ok := n.(*schema.Schema); ok {
		switch s.Type {
		case "boolean":
			return "bool"
		case "array":
			return "[]" + g.SanitiseName(s.GetChildren()[schema.SubSchemaKey].(*schema.Schema).GetName())
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

func (g *Renderer) GetOutputFilename(n tree.NodeTraverser) string {
	switch n.(type) {
	case *schema.Schema:
		return filepath.Join("models", g.SanitiseName(n.GetName())+".go")
	default:
		return g.SanitiseName(n.GetName()) + ".go"
	}
}

// SanitiseName should consume the result of an NodeTraverser's .GetName() function.
// It creates a string array that is compliant to go function name restrictions and
// joins the result before returning a single string.
func (g *Renderer) SanitiseName(s []string) string {
	caser := cases.Title(language.English)
	var temp []string
	for _, w := range s {
		switch true {
		case isHttpStatusCode(w):
			temp = append(temp, w)
			continue
		case strings.Contains(w, "/"):
			for _, split := range strings.Split(w, "/") {
				temp = append(temp, caser.String(CreateGoFunctionString(split)))
			}
			continue
		case strings.Contains(w, " "):
			for _, split := range strings.Split(w, " ") {
				temp = append(temp, caser.String(CreateGoFunctionString(split)))
			}
			continue
		case strings.Contains(w, "_"):
			for _, split := range strings.Split(w, "_") {
				temp = append(temp, caser.String(CreateGoFunctionString(split)))
			}
			continue
		}
		temp = append(temp, caser.String(CreateGoFunctionString(w)))
	}
	return strings.Join(temp, "")
}

// CreateGoFunctionString strips a string of any leading non-alphabetical chars, and all non-alphabetical and non-numerical
// characters that follow.
func CreateGoFunctionString(s string) (ret string) {
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

func isHttpStatusCode(s string) bool {
	if code, err := strconv.Atoi(s); err == nil {
		if 100 <= code && code <= 599 {
			return true
		}
	}
	return false
}
