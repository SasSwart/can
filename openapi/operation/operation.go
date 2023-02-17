package operation

import (
	"github.com/sasswart/gin-in-a-can/errors"
	"github.com/sasswart/gin-in-a-can/openapi/parameter"
	"github.com/sasswart/gin-in-a-can/openapi/request_body"
	"github.com/sasswart/gin-in-a-can/openapi/response"
	"github.com/sasswart/gin-in-a-can/openapi/root"
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
	RequestBody request_body.RequestBody      `yaml:"requestBody"` // can be a $ref
	Responses   map[string]*response.Response // can be a $ref
	//Callbacks 	map[string]*Callback // can be a $ref
	OperationId  string `yaml:"operationId"`
	ExternalDocs root.ExternalDocs
}

func (o *Operation) GetOutputFile() string {
	errors.Unimplemented("(o *Operation) GetOutputFile()")
	return ""
}

func (o *Operation) getRef() string {
	return ""
}

func (o *Operation) getChildren() map[string]tree.NodeTraverser {
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

	children["RequestBody"] = &o.RequestBody

	// Response
	for name := range o.Responses {
		r := o.Responses[name]
		children[name] = r
	}
	return children
}

func (o *Operation) setChild(i string, child tree.NodeTraverser) {
	switch child.(type) {
	case *parameter.Parameter:
		j, _ := strconv.Atoi(i)
		o.Parameters[j] = *child.(*parameter.Parameter)
		return
	case *request_body.RequestBody:
		o.RequestBody = *child.(*request_body.RequestBody)
		return
	case *response.Response:
		child.Responses[i] = child.(*response.Response)
		return
	default:
		errors.CastFail("(o *Operation) setChild", "NodeTraverser", "undefined (default case)")
	}
}
