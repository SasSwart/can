// Render shouldn't have to know about the implementations of it's base render package. It should simply
// allow them to be pluggable through the use of the renderer interface.

package render

import (
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
)

type EngineInterface interface {
	With(renderer Renderer, config config.Data) *Engine
	GetRenderer() Renderer
	BuildRenderNode() tree.TraversalFunc

	render(data tree.NodeTraverser, templateFile string) ([]byte, error)
}
type Engine struct {
	renderer Renderer
	config   config.Data
}

var _ EngineInterface = Engine{}

func (e Engine) With(renderer Renderer, config config.Data) *Engine {
	return &Engine{renderer: renderer, config: config}
}

func (e Engine) GetRenderer() Renderer {
	return e.renderer
}

func (e Engine) BuildRenderNode() tree.TraversalFunc {
	return func(key string, parent, node tree.NodeTraverser) (tree.NodeTraverser, error) {
		if s, ok := node.(*schema.Schema); ok {
			if s.Type != "object" && s.Type != "array" {
				return node, nil
			}
		}

		templateFile := getTemplateFilename(node)
		if templateFile == "" {
			return node, nil
		}
		output, err := e.render(node, templateFile)
		if err != nil {
			return node, fmt.Errorf("could not render into %s - possible syntax error in output after templating: %w", templateFile, err)
		}
		if !config.Dryrun {
			outPath := filepath.Join(e.config.GetOutputDir(), e.GetRenderer().GetOutputFilename(node))
			if err := writeToDisk(output, outPath); err != nil {
				return nil, err
			}
			if config.Debug {
				fmt.Printf("written %d bytes to %s\n", len(output), outPath)
			}
		}
		return node, nil
	}
}

func getTemplateFilename(node tree.NodeTraverser) string {
	switch node.(type) {
	case *openapi.OpenAPI:
		return "openapi.tmpl"
	case *path.Item:
		return "path_item.tmpl"
	case *operation.Operation:
		return "operation.tmpl"
	case *schema.Schema:
		return "schema.tmpl"
	}
	return ""
}

// Render contains the parsing and rendering steps
func (e Engine) render(node tree.NodeTraverser, templateFilename string) ([]byte, error) {
	r := e.GetRenderer()
	templateDirectory := e.config.GetTemplateFilesDir()
	parsedTemplate, err := r.ParseTemplate(templateFilename, templateDirectory)
	if err != nil {
		return nil, err
	}
	renderedOutput, err := r.RenderToText(parsedTemplate, node)
	if config.Debug {
		fmt.Printf("Rendering %s using %s\n", r.GetOutputFilename(node), templateFilename)
		fmt.Println(string(renderedOutput))
	}
	// format code based on formatter provided by interface
	formatted, err := r.Format(renderedOutput)
	if err != nil {
		return nil, err
	}
	return formatted, nil
}

func writeToDisk(contents []byte, outPath string) error {
	if _, err := os.Stat(filepath.Dir(outPath)); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(filepath.Dir(outPath), 0755)
		if err != nil {
			return err
		}
	}
	if err := os.WriteFile(outPath, contents, 0644); err != nil {
		return err
	}
	return nil
}
