package render

import (
	"github.com/sasswart/gin-in-a-can/tree"
	"text/template"
)

var _ Renderer = Base{}

type Renderer interface {
	SanitiseName([]string) string
	SanitiseType(n tree.NodeTraverser) string

	GetOutputFilename(n tree.NodeTraverser) string
	SetTemplateFuncMapping(template.FuncMap)
	GetTemplateFuncMapping() template.FuncMap
}

// Base defines the base render object. This should be used as a compositional base for specialising it's interface
// towards a specific use case.
type Base struct {
	TemplateFuncMapping template.FuncMap
}

func (b Base) SanitiseName(_ []string) string {
	panic("This should be overridden")
}

func (b Base) SanitiseType(_ tree.NodeTraverser) string {
	panic("This should be overridden")
}

func (b Base) GetOutputFilename(_ tree.NodeTraverser) string {
	panic("This should be overridden")
}

func (b Base) GetTemplateFuncMapping() template.FuncMap {
	return b.TemplateFuncMapping
}
func (b Base) SetTemplateFuncMapping(funcMap template.FuncMap) {
	b.TemplateFuncMapping = funcMap
}
