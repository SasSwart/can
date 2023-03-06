// Renderer implementations shouldn't have to know about the config package as the preparatory configuration needed for
// a renderer instance is handled when setRenderStrategy() is called in main.go

package golang

import (
	"github.com/sasswart/gin-in-a-can/config"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/render"
	"github.com/sasswart/gin-in-a-can/tree"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
	g.Base.SetTemplateFuncMap(&template.FuncMap{
		"ToUpper": strings.ToUpper,
		"ToTitle": func(s string) string {
			caser := cases.Title(language.English)
			return caser.String(s)
		},
		// TODO this should NOT be self-referential
		"SanitiseName": g.SanitiseName,
		"SanitiseType": g.SanitiseType,
	})
}

func (g *Renderer) GetTemplateFuncMap() *template.FuncMap {
	return g.Base.GetTemplateFuncMap()
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
			return "[]" + g.SanitiseName(s.GetChildren()[schema.ItemsKey].(*schema.Schema).GetName())
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
	return g.SanitiseName(n.GetName()) + ".go"
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
				temp = append(temp, caser.String(CreateFunctionString(split)))
			}
			continue
		case strings.Contains(w, " "):
			for _, split := range strings.Split(w, " ") {
				temp = append(temp, caser.String(CreateFunctionString(split)))
			}
			continue
		case strings.Contains(w, "_"):
			for _, split := range strings.Split(w, "_") {
				temp = append(temp, caser.String(CreateFunctionString(split)))
			}
			continue
		}
		temp = append(temp, caser.String(CreateFunctionString(w)))
	}
	return strings.Join(temp, "")
}

// CreateFunctionString strips a string of any leading non-alphabetical chars, and all non-alphabetical and non-numerical
// characters that follow.
func CreateFunctionString(s string) (ret string) {
	for i, char := range []rune(s) {
		if i == 0 {
			// function names must start with alphabetical characters in go
			if isAlpha(char) {
				ret += string(char)
			}
			continue
		}
		if isAlphaNum(char) {
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

func isAlpha(r rune) bool {
	return ('A' <= r && r <= 'Z') || ('a' <= r && r <= 'z')
}
func isAlphaNum(r rune) bool {
	return isAlpha(r) || ('0' <= r && r <= '9')
}

// TODO Fix and test for robustness. Make sure this doesn't infringe on logic dealt with elsewhere
func ToTitle(s string) (ret string) {
	caser := cases.Title(language.English)
	var splitBy []rune
	for _, r := range []rune(s) {
		if !isAlphaNum(r) {
			splitBy = append(splitBy, r)
		}
	}
	buf := []string{s}
	for 0 < len(splitBy) {
		for _, word := range buf {
			buf = strings.Split(word, string(splitBy[0]))
			splitBy = splitBy[1:]
		}
	}
	return caser.String(ret)
}

func NewGinServerTestConfig() config.Data {
	config.ConfigFilePath = "../render/go/config_goginserver_test.yml"
	config.Debug = true
	return config.Data{
		Template: config.Template{
			Name: "go-gin",
		},
		OpenAPIFile: "../openapi/test/fixtures/validation_no_refs.yaml",
		OutputPath:  ".",
	}
}
func NewGoClientTestConfig() config.Data {
	config.ConfigFilePath = "../render/go/config_goclient_test.yml"
	config.Debug = true
	return config.Data{
		Template: config.Template{
			Name: "go-client",
		},
		OpenAPIFile: "../openapi/test/fixtures/validation_no_refs.yaml",
		OutputPath:  ".",
	}
}
