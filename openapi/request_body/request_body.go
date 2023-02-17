package request_body

import (
	"github.com/sasswart/gin-in-a-can/errors"
	"github.com/sasswart/gin-in-a-can/openapi/media_type"
	"github.com/sasswart/gin-in-a-can/tree"
)

var _ tree.NodeTraverser = &RequestBody{}

// RequestBody is a programmatic representation of the RequestBody object defined here: https://swagger.io/specification/#request-body-object
type RequestBody struct {
	tree.Node
	Ref         string
	Description string
	Content     map[string]*media_type.MediaType
	Required    bool
}

func (r *RequestBody) GetOutputFile() string {
	errors.Unimplemented("(r *RequestBody) GetOutputFile()")
	return ""
}

func (r *RequestBody) GetName() string {
	name := r.GetParent().GetName() + r.name
	return name
}

func (r *RequestBody) GetRef() string {
	return r.Ref
}

func (r *RequestBody) GetChildren() map[string]tree.NodeTraverser {
	children := map[string]tree.NodeTraverser{}
	for name := range r.Content {
		mediaType := r.Content[name]
		children[name] = mediaType
	}
	return children
}

func (r *RequestBody) SetChild(i string, t tree.NodeTraverser) {
	if content, ok := t.(*media_type.MediaType); ok {
		r.Content[i] = content
		return
	}
	errors.CastFail("(r *RequestBody) setChild()", "NodeTraverser", "*media_type.MediaType")
}
