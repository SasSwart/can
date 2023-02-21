package response

import (
	"github.com/sasswart/gin-in-a-can/errors"
	"github.com/sasswart/gin-in-a-can/openapi/media"
	"github.com/sasswart/gin-in-a-can/tree"
	"strconv"
)

var _ tree.NodeTraverser = &Response{}

// Response is a programmatic representation of the Response object defined here: https://swagger.io/specification/#response-object
type Response struct {
	tree.Node
	Ref         string
	Description string `yaml:"description"`
	Content     map[string]media.Type

	// TODO these cause an import cycle
	//Headers     map[string]openapi.Header // can also be a $ref
	//Links       map[string]openapi.Link // can also be a $ref
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
	if r.Content == nil {
		r.Content = make(map[string]media.Type, 4)
	}
	if _, err := strconv.Atoi(i); err != nil || i == "default" {
		mediaType, ok := t.(*media.Type)
		if !ok {
			errors.CastFail("(r *Response) SetChild()", "NodeTraverser", "*media_type.Type")
		}
		r.Content[i] = *mediaType
		return
	}
	errors.UndefinedBehaviour("(r *Response) SetChild()")
}
