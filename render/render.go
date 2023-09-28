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
}

// Base defines the base render object. This should be used as a compositional base for specialising it's interface
// towards a specific use case.
//type Base struct {
//	templateFuncMapping *template.FuncMap
//}
//
//func (b *Base) GetTemplateFuncMap() *template.FuncMap {
//	return b.templateFuncMapping
//}
//func (b *Base) SetTemplateFuncMap(funcMap *template.FuncMap) {
//	b.templateFuncMapping = funcMap
//}
//
//func (b *Base) Format(input []byte) ([]byte, error) {
//	fmt.Println("no formatter implemented")
//	return input, nil
//}
