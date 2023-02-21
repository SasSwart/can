package operation

import (
	"github.com/sasswart/gin-in-a-can/errors"
	"github.com/sasswart/gin-in-a-can/openapi/parameter"
	"github.com/sasswart/gin-in-a-can/openapi/request"
	"github.com/sasswart/gin-in-a-can/openapi/response"
	"github.com/sasswart/gin-in-a-can/tree"
	"strconv"
)

var _ tree.NodeTraverser = &Operation{}

// Operation is a programmatic representation of the Operation object defined here: https://swagger.io/specification/#operation-object
type Operation struct {
	tree.Node
	Tags        []string
	Summary     string
	Description string
	Parameters  []parameter.Parameter         // can be a $ref
	RequestBody request.Body                  `yaml:"requestBody"` // can be a $ref
	Responses   map[string]*response.Response // can be a $ref
	//Callbacks 	map[string]*Callback // can be a $ref
	OperationId string `yaml:"operationId"`

	// TODO this will cause an import cycle
	//ExternalDocs root.ExternalDocs
}

func (o *Operation) GetRef() string {
	return ""
}

func (o *Operation) GetChildren() map[string]tree.NodeTraverser {
	children := map[string]tree.NodeTraverser{}
	if o == nil {
		return children
	}
	// Parameters
	for i := range o.Parameters {
		p := o.Parameters[i]
		paramIndex := strconv.Itoa(i)
		children[paramIndex] = &p
	}

	// Request Body

	children["Body"] = &o.RequestBody

	// Response
	for name := range o.Responses {
		r := o.Responses[name]
		children[name] = r
	}
	return children
}

func (o *Operation) SetChild(i string, child tree.NodeTraverser) {
	switch child.(type) {
	case *parameter.Parameter:
		o.Parameters = append(o.Parameters, *child.(*parameter.Parameter))
		return
	case *request.Body:
		o.RequestBody = *child.(*request.Body)
		return
	case *response.Response:
		if o.Responses == nil {
			o.Responses = make(map[string]*response.Response, 4)
		}
		o.Responses[i] = child.(*response.Response)
		return
	default:
		errors.CastFail("(o *Operation) setChild", "NodeTraverser", "undefined (default case)")
	}
}
