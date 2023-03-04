package request

import (
	"github.com/sasswart/gin-in-a-can/errors"
	"github.com/sasswart/gin-in-a-can/openapi/media"
	"github.com/sasswart/gin-in-a-can/tree"
)

const BodyKey = "Body"

var _ tree.NodeTraverser = &Body{}

// Body is a programmatic representation of the Body object defined here: https://swagger.io/specification/#request-body-object
type Body struct {
	tree.Node
	Ref         string
	Description string
	Content     map[string]*media.Type
	Required    bool
}

func (r *Body) GetRef() string {
	return r.Ref
}

func (r *Body) GetName() []string {
	if r.GetParent() == nil {
		return []string{"Request", r.Name}
	}
	return append(r.GetParent().GetName(), "Request", r.Name)
}

func (r *Body) GetChildren() map[string]tree.NodeTraverser {
	children := map[string]tree.NodeTraverser{}
	for name := range r.Content {
		mediaType := r.Content[name]
		children[name] = mediaType
	}
	return children
}

func (r *Body) SetChild(i string, t tree.NodeTraverser) {
	if r.Content == nil {
		r.Content = make(map[string]*media.Type, 4)
	}
	if content, ok := t.(*media.Type); ok {
		r.Content[i] = content
		return
	}
	errors.CastFail("(r *Body) setChild()", "NodeTraverser", "*media_type.Type")
}
