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
	render(data tree.NodeTraverser, templateFile string) ([]byte, error)
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
func (e Engine) render(node tree.NodeTraverser, templateFilename string) ([]byte, error) {
	buff := bytes.NewBuffer([]byte{})
	templater := template.New(templateFilename)

	templater.Funcs(templateFuncMap)

	parsedTemplate, err := templater.ParseGlob(fmt.Sprintf("%s/*.tmpl", e.config.GetTemplateDir()))
	if err != nil {
		return nil, err
	}

	// TODO figure out why parsedTemplate.Tree is nil

	err = parsedTemplate.Execute(buff, node)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Rendering %s using %s\n", e.renderer.GetOutputFilename(node), templateFilename)

	outputDirAbs := filepath.Dir(e.config.GetOutputDir())
	if _, err := os.Stat(outputDirAbs); errors.Is(err, os.ErrNotExist) {
		_ = os.MkdirAll(outputDirAbs, 0755)
	}

	if !config.Dryrun {
		fmt.Println(string(buff.Bytes()))
	}
	outPath := filepath.Join(e.config.GetOutputDir(), e.renderer.GetOutputFilename(node))
	err = os.WriteFile(outPath, buff.Bytes(), 0644)
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

		_, err := e.render(node, templateFile)
		if err != nil {
			return node, err
		}

		return node, nil
	}
}
