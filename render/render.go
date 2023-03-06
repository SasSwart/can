package render

import (
	"github.com/sasswart/gin-in-a-can/tree"
	"text/template"
)

type Renderer interface {
	SanitiseName([]string) string
	SanitiseType(n tree.NodeTraverser) string

	GetOutputFilename(n tree.NodeTraverser) string
	SetTemplateFuncMap(*template.FuncMap)
	GetTemplateFuncMap() *template.FuncMap
}

// Base defines the base render object. This should be used as a compositional base for specialising it's interface
// towards a specific use case.
type Base struct {
	templateFuncMapping *template.FuncMap
}

func (b *Base) GetTemplateFuncMap() *template.FuncMap {
	return b.templateFuncMapping
}
func (b *Base) SetTemplateFuncMap(funcMap *template.FuncMap) {
	b.templateFuncMapping = funcMap
}
