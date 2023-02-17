package parameter

import (
	"github.com/sasswart/gin-in-a-can/openapi/errors"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/tree"
)

var _ tree.NodeTraverser = &Parameter{}

// Parameter is a programmatic representation of the Parameter object defined here: https://swagger.io/specification/#parameter-object
type Parameter struct {
	tree.Node
	Ref             string         `yaml:"$ref"`
	Name            string         `yaml:"name"`
	In              string         `yaml:"in"`
	Description     string         `yaml:"description"`
	Required        bool           `yaml:"required"`
	Deprecated      bool           `yaml:"deprecated"`
	AllowEmptyValue bool           `yaml:"allowEmptyValue"`
	Schema          *schema.Schema // Acts as alternative description of param
}

func (p *Parameter) getRef() string {
	return p.Ref
}

func (p *Parameter) getChildren() map[string]tree.NodeTraverser {
	return map[string]tree.NodeTraverser{
		"Model": p.Schema,
	}
}

func (p *Parameter) GetOutputFile() string {
	// TODO this override can be removed and this logic dealt with by the node{} composable
	return p.GetRenderer().GetOutputFile(p)
}

func (p *Parameter) GetName() string {
	// TODO can this be done through switch node.parent.(Type)?
	name := p.GetParent().GetName() + p.GetRenderer().SanitiseName(p.Name) + "Parameter"
	return name
}

func (p *Parameter) setChild(_ string, t tree.NodeTraverser) {
	s, ok := t.(*schema.Schema)
	if !ok {
		panic("(p *Parameter) setChild(): " + errors.ErrCastFail)
	}
	p.Schema = s
}
