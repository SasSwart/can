// Renderer implementations shouldn't have to know about the config package as the preparatory configuration needed for
// a renderer instance is handled when setRenderStrategy() is called in main.go

package golang

import (
	"bytes"
	"fmt"
	"github.com/sasswart/gin-in-a-can/config"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/render"
	"github.com/sasswart/gin-in-a-can/tree"
	"go/format"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strconv"
	"strings"
	"text/template"
)

var _ render.Renderer = &Renderer{}

type Renderer struct {
	// TODO add config for flags to be referenced where needed once globals are removed
	funcMap *template.FuncMap
}

// RenderNode writes the relevant information contained within an openapi node into the provided template and returns
// the result.
func (g *Renderer) RenderNode(parsedTemplate *template.Template, node tree.NodeTraverser) ([]byte, error) {
	if parsedTemplate == nil {
		return nil, fmt.Errorf("parsedTemplate is nil")
	}
	buff := bytes.NewBuffer([]byte{})
	err := parsedTemplate.Execute(buff, node)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

func (g *Renderer) SetTemplateFuncMap(f *template.FuncMap) {
	g.funcMap = f
}

func (g *Renderer) GetTemplateFuncMap() *template.FuncMap {
	return g.funcMap
}

// Format applies the same formatter in use in gofmt to a rendered template. This is the standard formatter for all go
// code.
func (g *Renderer) Format(input []byte) ([]byte, error) {
	return format.Source(input)
}

// GetOutputFilename provides an appropriate filename for any rendered output based on the node provided.
func (g *Renderer) GetOutputFilename(n tree.NodeTraverser) string {
	// TODO: SanitiseName is being used to clean output while rendering templates and openapi nodes to byte output.
	// 	 Perhaps it shouldn't be used here.
	return SanitiseName(n.GetName()) + ".go"
}

// DefaultFuncMap contains a function map with sensible defaults for go template rendering
func DefaultFuncMap() *template.FuncMap {
	return &template.FuncMap{
		"ToUpper": strings.ToUpper,
		"ToTitle": ToTitle,
		"StringsReplace": func(input, from, to string) string {
			return strings.Replace(input, from, to, -1)
		},
		"SanitiseName": SanitiseName,
		"SanitiseType": SanitiseType,
	}
}

// SanitiseType sanitizes the prepares the contents of the Type field of a node for use by the renderer
func SanitiseType(n tree.NodeTraverser) string {
	if n == nil {
		return ""
	}
	if s, ok := n.(*schema.Schema); ok {
		switch s.Type {
		case "boolean":
			return "bool"
		case "array":
			return "[]" + SanitiseName(s.GetChildren()[schema.ItemsKey].(*schema.Schema).GetName())
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

// SanitiseName should consume the result of an NodeTraverser's .GetName() function.
// It creates a string array that is compliant to go function name restrictions and
// joins the result before returning a single string.
func SanitiseName(s []string) string {
	caser := cases.Title(language.English)
	var temp []string
	for _, w := range s {
		var delim string
		switch true {
		case IsHttpStatusCode(w):
			temp = append(temp, w)
			continue
		case strings.Contains(w, "/"):
			delim = "/"
		case strings.Contains(w, " "):
			delim = " "
		case strings.Contains(w, "_"):
			delim = "_"
		case strings.Contains(w, "-"):
			delim = "-"
		default:
			temp = append(temp, caser.String(CreateFunctionString(w)))
			continue
		}
		for _, split := range strings.Split(w, delim) {
			temp = append(temp, caser.String(CreateFunctionString(split)))
		}
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

func IsHttpStatusCode(s string) bool {
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

func flatten[T any](nestedList [][]T) []T {
	var flattenedList []T
	for _, subList := range nestedList {
		for _, innerMostElement := range subList {
			flattenedList = append(flattenedList, innerMostElement)
		}
	}
	return flattenedList
}

func ToTitle(s string) (ret string) {
	caser := cases.Title(language.English)
	var splitBy []rune
	for _, r := range []rune(s) {
		if !isAlphaNum(r) {
			splitBy = append(splitBy, r)
		}
	}
	buf := []string{s}
	for _, delim := range splitBy {
		splitbuf := make([][]string, 0)
		for _, word := range buf {
			splitbuf = append(splitbuf, strings.Split(word, string(delim)))
		}
		buf = flatten[string](splitbuf)
	}
	for _, word := range buf {
		ret += caser.String(word)
	}
	return ret
}

func MustLoadGinServerTestConfig(configPath, openAPIPath string) config.Data {
	config.ConfigFilePath = configPath
	config.Debug = true
	c := config.Data{
		Template: config.Template{
			Name: "go-gin",
		},
		OpenAPIFile: openAPIPath,
		OutputPath:  ".",
	}
	if err := c.Load(); err != nil {
		panic(err)
	}
	return c
}
func MustLoadGoClientTestConfig(configPath, openAPIPath string) config.Data {
	// TODO: `Must` this function and load the returned config before returning it
	config.ConfigFilePath = configPath
	config.Debug = true
	c := config.Data{
		Template: config.Template{
			Name: "go-client",
		},
		OpenAPIFile: openAPIPath,
		OutputPath:  ".",
	}
	if err := c.Load(); err != nil {
		panic(err)
	}
	return c
}
