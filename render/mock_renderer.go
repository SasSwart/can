package render

import (
	"github.com/sasswart/gin-in-a-can/tree"
	"text/template"
)

// MockRenderer is a mock implementation of the Renderer interface for testing.
type MockRenderer struct {
	funcMap *template.FuncMap
}

var _ Renderer = &MockRenderer{}

func (mr *MockRenderer) SetTemplateFuncMap(funcMap *template.FuncMap) {
	mr.funcMap = funcMap
	return
}

func (mr *MockRenderer) GetTemplateFuncMap() *template.FuncMap {
	return mr.funcMap
}

func (mr *MockRenderer) RenderNode(template *template.Template, traverser tree.NodeTraverser) ([]byte, error) {
	return []byte("rendered_output"), nil
}

func (mr *MockRenderer) GetOutputFilename(node tree.NodeTraverser) string {
	return "output_filename"
}

func (mr *MockRenderer) Format(output []byte) ([]byte, error) {
	return output, nil
}
