package response

import (
	"github.com/sasswart/gin-in-a-can/errors"
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/openapi/media_type"
	"github.com/sasswart/gin-in-a-can/tree"
	"strconv"
)

var _ tree.NodeTraverser = &Response{}

// Response is a programmatic representation of the Response object defined here: https://swagger.io/specification/#response-object
type Response struct {
	tree.Node
	Ref         string
	Description string                    `yaml:"description"`
	Headers     map[string]openapi.Header // can also be a $ref
	Content     map[string]media_type.MediaType
	Links       map[string]openapi.Link // can also be a $ref
}

func (r *Response) GetOutputFile() string {
	errors.Unimplemented("(r *Response) GetOutputFile()")
	return ""
}

func (r *Response) GetName() string {
	return r.GetParent().GetName() + r.GetRenderer().SanitiseName(r.name) + "Response"
}

func (r *Response) GetRef() string {
	return r.Ref
}

func (r *Response) GetChildren() map[string]tree.NodeTraverser {
	responses := map[string]tree.NodeTraverser{} // Where string is either `default` or an HTTP status code
	for name, mediaType := range r.Content {
		if _, err := strconv.Atoi(name); err != nil || name == "default" {
			responses[name] = &mediaType
		} else {
			errors.UndefinedBehaviour("(r *Response) GetChildren()")
		}
	}
	return responses
}

func (r *Response) SetChild(i string, t tree.NodeTraverser) {
	if _, err := strconv.Atoi(i); err != nil || i == "default" {
		mediaType, ok := t.(*media_type.MediaType)
		if !ok {
			errors.CastFail("(r *Response) SetChild()", "NodeTraverser", "*media_type.MediaType")
		}
		r.Content[i] = *mediaType
		return
	}
	errors.UndefinedBehaviour("(r *Response) SetChild()")
}
