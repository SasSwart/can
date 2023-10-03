package render

import (
	"github.com/sasswart/gin-in-a-can/tree"
	"text/template"
)

// MockRenderer is a mock implementation of the Renderer interface for testing.
type MockRenderer struct{}

var _ Renderer = MockRenderer{}

func (mr MockRenderer) SetTemplateFuncMap(funcMap *template.FuncMap) {
	return
}

func (mr MockRenderer) GetTemplateFuncMap() *template.FuncMap {
	return &template.FuncMap{}
}

func (mr MockRenderer) ParseTemplate(filename, directory string) (*template.Template, error) {
	return template.New(filename), nil
}

func (mr MockRenderer) RenderNode(template *template.Template, traverser tree.NodeTraverser) ([]byte, error) {
	return []byte("rendered_output"), nil
}

func (mr MockRenderer) GetOutputFilename(node tree.NodeTraverser) string {
	return "output_filename"
}

func (mr MockRenderer) Format(output []byte) ([]byte, error) {
	return output, nil
}
