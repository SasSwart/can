package path

import (
	"github.com/sasswart/gin-in-a-can/errors"
	"github.com/sasswart/gin-in-a-can/openapi/operation"
	"github.com/sasswart/gin-in-a-can/openapi/parameter"
	"github.com/sasswart/gin-in-a-can/tree"
	"net/http"
	"path/filepath"
	"strings"
)

var _ tree.NodeTraverser = &Item{}

// Item is a programmatic representation of the Path Item object defined here: https://swagger.io/specification/#path-item-object
type Item struct {
	tree.Node
	Ref         string `yaml:"$ref"` // must be defined in the format of a PathItem object
	Summary     string
	Description string
	Get         *operation.Operation
	Post        *operation.Operation
	Patch       *operation.Operation
	Delete      *operation.Operation
	Parameters  []parameter.Parameter
}

func (p *Item) GetRef() string {
	return p.Ref
}

func (p *Item) GetPath() string {
	return strings.Join(p.GetName(), "")
}

func (p *Item) SetChild(i string, child tree.NodeTraverser) {
	if n, ok := child.(*operation.Operation); ok {
		switch strings.ToUpper(i) {
		case http.MethodGet:
			p.Get = n
		case http.MethodPost:
			p.Post = n
		case http.MethodPatch:
			p.Patch = n
		case http.MethodDelete:
			p.Delete = n
		}
		return
	}
	errors.CastFail("(o *Root) setChild", "NodeTraverser", "*schema.Schema")
}
func (p *Item) GetChildren() map[string]tree.NodeTraverser {
	return p.Operations()
}

// Operations is public as it's called by the templater before rendering output
func (p *Item) Operations() map[string]tree.NodeTraverser {
	operations := map[string]tree.NodeTraverser{}
	if p.Get != nil {
		operations[http.MethodGet] = p.Get
	}
	if p.Post != nil {
		operations[http.MethodPost] = p.Post
	}
	if p.Patch != nil {
		operations[http.MethodPatch] = p.Patch
	}
	if p.Delete != nil {
		operations[http.MethodDelete] = p.Delete
	}
	return operations
}

func (p *Item) GetBasePath() string {
	if p.GetParent() == nil {
		return p.Node.GetBasePath()
	}
	return filepath.Join(p.GetParent().GetBasePath(), filepath.Dir(p.Ref))
}
