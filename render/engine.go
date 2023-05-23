// Render shouldn't have to know about the implementations of it's base render package. It should simply
// allow them to be pluggable through the use of the renderer interface.

package render

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/sasswart/gin-in-a-can/config"
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/openapi/operation"
	"github.com/sasswart/gin-in-a-can/openapi/path"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/tree"
	"os"
	"path/filepath"
	"text/template"
)

type EngineInterface interface {
	With(renderer Renderer, config config.Data) *Engine
	GetRenderer() *Renderer
	BuildRenderNode() tree.TraversalFunc

	render(data tree.NodeTraverser, templateFile string) ([]byte, error)
}
type Engine struct {
	renderer *Renderer
	config   config.Data
}

var _ EngineInterface = Engine{}

func (e Engine) With(renderer Renderer, config config.Data) *Engine {
	return &Engine{renderer: &renderer, config: config}
}

func (e Engine) GetRenderer() *Renderer {
	return e.renderer
}

func (e Engine) BuildRenderNode() tree.TraversalFunc {
	return func(key string, parent, node tree.NodeTraverser) (tree.NodeTraverser, error) {
		var templateFile string
		switch node.(type) {
		case *openapi.OpenAPI:
			templateFile = "openapi.tmpl"
		case *path.Item:
			templateFile = "path_item.tmpl"
		case *operation.Operation:
			templateFile = "operation.tmpl"
		case *schema.Schema:
			schemaType := node.(*schema.Schema).Type
			if schemaType != "object" && schemaType != "array" {
				return node, nil
			}
			templateFile = "schema.tmpl"
		}

		if templateFile == "" {
			return node, nil
		}

		_, err := e.render(node, templateFile)
		if err != nil {
			return node, err
		}

		return node, nil
	}
}

// Render contains the parsing and rendering steps
func (e Engine) render(node tree.NodeTraverser, templateFilename string) ([]byte, error) {
	r := *e.GetRenderer()
	if r.GetTemplateFuncMap() == nil {
		return nil, fmt.Errorf("e.render()::: template function mapping not set")
	}
	buff := bytes.NewBuffer([]byte{})
	templater := template.New(templateFilename)
	templater.Funcs(*r.GetTemplateFuncMap())

	parsedTemplate, err := templater.ParseGlob(fmt.Sprintf("%s/*.tmpl", e.config.GetTemplateFilesDir()))
	if err != nil {
		return nil, err
	}

	err = parsedTemplate.Execute(buff, node)
	if err != nil {
		return nil, err
	}

	if config.Debug {
		fmt.Printf("Rendering %s using %s\n", r.GetOutputFilename(node), templateFilename)
		fmt.Println(string(buff.Bytes()))
	}

	outPath := filepath.Join(e.config.GetOutputDir(), r.GetOutputFilename(node))
	if !config.Dryrun {
		if _, err := os.Stat(filepath.Dir(outPath)); errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll(filepath.Dir(outPath), 0755)
			if err != nil {
				return nil, err
			}
		}
		err = os.WriteFile(outPath, buff.Bytes(), 0644)
		if err != nil {
			return nil, err
		}
		if config.Debug {
			fmt.Printf("written %d bytes to %s\n", buff.Len(), outPath)
		}
	}
	return buff.Bytes(), nil
}
