package media_type

import (
	"github.com/sasswart/gin-in-a-can/openapi/errors"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/tree"
)

var _ tree.NodeTraverser = &MediaType{}

// MediaType is a programmatic representation of the MediaType object defined here: https://swagger.io/specification/#media-type-object
type MediaType struct {
	tree.Node
	name   string
	Schema *schema.Schema
}

func (m *MediaType) GetRef() string {
	return ""
}

func (m *MediaType) GetName() string {
	return m.GetParent().GetName() + m.name
}

func (m *MediaType) GetChildren() map[string]tree.NodeTraverser {
	return map[string]tree.NodeTraverser{
		"Model": m.Schema,
	}
}

func (m *MediaType) SetChild(_ string, t tree.NodeTraverser) {
	if s, ok := t.(*schema.Schema); ok {
		m.Schema = s
		return
	}
	panic("(m *MediaType) setChild(): " + errors.ErrCastFail)
}
