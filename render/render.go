package render

import (
	"github.com/sasswart/gin-in-a-can/tree"
	"text/template"
)

type Renderer interface {
	GetOutputFilename(n tree.NodeTraverser) string
	SetTemplateFuncMap(*template.FuncMap)
	GetTemplateFuncMap() *template.FuncMap
	Format([]byte) ([]byte, error)
	RenderNode(*template.Template, tree.NodeTraverser) ([]byte, error)
}
