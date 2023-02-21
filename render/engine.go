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
	New(renderer Renderer, config config.Data) *Engine
	GetRenderer() Renderer
	Render(data tree.NodeTraverser, templateFile string) ([]byte, error)
	BuildRenderNode() tree.TraversalFunc
}
type Engine struct {
	renderer Renderer
	config   config.Data
}

var _ EngineInterface = Engine{}

func (e Engine) New(renderer Renderer, config config.Data) *Engine {
	return &Engine{renderer: renderer, config: config}
}

func (e Engine) GetRenderer() Renderer {
	return e.renderer
}

// Render contains the parsing and rendering steps
func (e Engine) Render(node tree.NodeTraverser, templateFile string) ([]byte, error) {

	outputFileAbs = filepath.Join(outputFileAbs, e.renderer.GetOutputFile(node))

	buff := bytes.NewBuffer([]byte{})

	templater := template.New(templateFile)

	templater.Funcs(templateFuncMap)

	parsedTemplate, err := templater.ParseGlob(fmt.Sprintf("%s/*.tmpl", templateDirAbs))
	if err != nil {
		return nil, err
	}

	err = parsedTemplate.Execute(buff, node)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Rendering %s using %s\n", e.renderer.GetOutputFile(node), templateFile)

	outputDirAbs := filepath.Dir(outputFileAbs)
	if _, err := os.Stat(outputDirAbs); errors.Is(err, os.ErrNotExist) {
		_ = os.MkdirAll(outputDirAbs, 0755)
	}

	err = os.WriteFile(outputFileAbs, buff.Bytes(), 0644)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}
func (e Engine) BuildRenderNode() tree.TraversalFunc {
	return func(key string, parent, node tree.NodeTraverser) (tree.NodeTraverser, error) {
		var templateFile string
		switch node.(type) {
		case *openapi.OpenAPI:
			templateFile = "openapi.tmpl"
		case *path.Item:
			templateFile = "path_item.tmpl"
		case *schema.Schema:
			schemaType := node.(*schema.Schema).Type
			if schemaType != "object" && schemaType != "array" {
				return node, nil
			}
			templateFile = "schema.tmpl"
		case *operation.Operation:
			templateFile = "operation.tmpl"
		}

		if templateFile == "" {
			return node, nil
		}

		_, err := e.Render(node, templateFile)
		if err != nil {
			return node, err
		}

		return node, nil
	}
}
