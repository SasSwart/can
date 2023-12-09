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
	"text/template"
)

type EngineInterface interface {
	ParseTemplate(string, string) (*template.Template, error)
	GetRenderer() Renderer
	SetRenderer(renderer Renderer)
	Render(key string, parent, node tree.NodeTraverser) (tree.NodeTraverser, error)
}
type Engine struct {
	// renderer contains the object responsible for
	renderer Renderer
	config   config.Data
}

var _ EngineInterface = &Engine{}

func NewEngine(config config.Data) Engine {
	return Engine{config: config}
}

func (e *Engine) GetRenderer() Renderer {
	return e.renderer
}

func (e *Engine) SetRenderer(r Renderer) {
	e.renderer = r
}

// ParseTemplate reads the given template and returns a template object with the appropriate function map applied.
func (e *Engine) ParseTemplate(templateFilename, templateDirectory string) (*template.Template, error) {
	renderer := e.GetRenderer()
	templater := template.New(templateFilename)
	funcMap := renderer.GetTemplateFuncMap()
	templater.Funcs(*funcMap)
	return templater.ParseGlob(fmt.Sprintf("%s/*.tmpl", templateDirectory))
}

// Render fetches the appropriate template name and renders yielded by it and the provided node to disk.
func (e *Engine) Render(key string, parent, node tree.NodeTraverser) (tree.NodeTraverser, error) {
	if s, ok := node.(*schema.Schema); ok {
		// For schemas, only render objects and arrays
		if s.Type != "object" && s.Type != "array" {
			return node, nil
		}
	}

	// find appropriate template name based on node provided
	templateFile := GetTemplateFilename(node)
	if templateFile == "" {
		return node, nil
	}

	// render node and template into output ready to be written to disk
	output, err := e.render(node, templateFile)
	if err != nil {
		return node, fmt.Errorf("could not render into %s: %w", templateFile, err)
	}
	if !config.Dryrun {
		outPath := filepath.Join(e.config.GetOutputDir(), e.GetRenderer().GetOutputFilename(node))
		if err := WriteToDisk(output, outPath); err != nil {
			return nil, err
		}
		if config.Debug {
			fmt.Printf("written %d bytes to %s\n", len(output), outPath)
		}
	}
	return node, nil
}

func GetTemplateFilename(node tree.NodeTraverser) string {
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

// Render contains the parsing and rendering steps. It parses a template, renders a node into it and applies any
// required formatting.
func (e *Engine) render(node tree.NodeTraverser, templateFilename string) ([]byte, error) {
	renderer := e.GetRenderer()
	templateDirectory := e.config.GetTemplateFilesDir()

	// fetch appropriate template object
	parsedTemplate, err := e.ParseTemplate(templateFilename, templateDirectory)
	if err != nil {
		return nil, err
	}

	// render the node and template to bytes
	renderedOutput, err := renderer.RenderNode(parsedTemplate, node)
	// TODO: this debug output could be moved into renderer.RenderNode
	if config.Debug {
		fmt.Printf("Rendering %s using %s\n", renderer.GetOutputFilename(node), templateFilename)
		fmt.Println(string(renderedOutput))
	}

	// format code based on formatter provided by interface
	formatted, err := renderer.Format(renderedOutput)
	if err != nil {
		return nil, fmt.Errorf("could not format output: %w", err)
	}
	return formatted, nil
}

func WriteToDisk(contents []byte, outPath string) error {
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

func (e *Engine) BuildTestRenderNode(outputChan chan<- []byte) tree.TraversalFunc {
	return func(key string, parent, node tree.NodeTraverser) (tree.NodeTraverser, error) {
		if s, ok := node.(*schema.Schema); ok {
			if s.Type != "object" && s.Type != "array" {
				return node, nil
			}
		}

		templateFile := GetTemplateFilename(node)
		if templateFile == "" {
			return node, nil
		}
		output, err := e.render(node, templateFile)
		if err != nil {
			return node, fmt.Errorf("could not render into %s: %w", templateFile, err)
		}
		outputChan <- output
		return node, nil
	}
}
