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
	ParseTemplate(string, string) (*template.Template, error)
	RenderToText(*template.Template, tree.NodeTraverser) ([]byte, error)
}
