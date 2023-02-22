package parameter

import (
	"github.com/sasswart/gin-in-a-can/errors"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/tree"
)

var _ tree.NodeTraverser = &Parameter{}

// Parameter is a programmatic representation of the Parameter object defined here: https://swagger.io/specification/#parameter-object
type Parameter struct {
	tree.Node
	Ref             string         `yaml:"$ref"`
	ParamName       string         `yaml:"name"` // TODO not sure if this is a duplicate of the name var in the accompanying Node
	In              string         `yaml:"in"`
	Description     string         `yaml:"description"`
	Required        bool           `yaml:"required"`
	Deprecated      bool           `yaml:"deprecated"`
	AllowEmptyValue bool           `yaml:"allowEmptyValue"`
	Schema          *schema.Schema // Acts as alternative description of param
}

func (p *Parameter) GetRef() string {
	return p.Ref
}

func (p *Parameter) GetChildren() map[string]tree.NodeTraverser {
	return map[string]tree.NodeTraverser{
		schema.Key: p.Schema,
	}
}
func (p *Parameter) GetName() string {
	if p.GetParent() == nil {
		return p.Name + "Parameter"
	}
	return p.GetParent().GetName() + p.Name + "Parameter"
}

func (p *Parameter) SetChild(_ string, t tree.NodeTraverser) {
	if s, ok := t.(*schema.Schema); ok {
		p.Schema = s
		return
	}
	errors.CastFail("(r *Body) setChild()", "NodeTraverser", "*media_type.Type")
}
