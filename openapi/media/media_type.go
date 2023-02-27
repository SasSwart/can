package media

import (
	"github.com/sasswart/gin-in-a-can/errors"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/tree"
)

const (
	JSONKey = "application/json"
)

var _ tree.NodeTraverser = &Type{}

// Type is a programmatic representation of the Type object defined here: https://swagger.io/specification/#media-type-object
type Type struct {
	tree.Node
	//name   string // not sure if this is necessary
	Schema *schema.Schema
}

func (m *Type) GetRef() string {
	return ""
}

func (m *Type) GetChildren() map[string]tree.NodeTraverser {
	return map[string]tree.NodeTraverser{
		schema.Key: m.Schema,
	}
}

func (m *Type) SetChild(_ string, t tree.NodeTraverser) {
	if s, ok := t.(*schema.Schema); ok {
		m.Schema = s
		return
	}
	errors.CastFail("(m *Type) setChild()", "NodeTraverser", "*schema.Schema")
}
