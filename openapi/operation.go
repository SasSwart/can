package openapi

import (
	"strconv"
)

// communicate by sharing memory ;)
var _ Traversable = &Operation{}

// Operation is a programmatic representation of the Operation object defined here: https://swagger.io/specification/#operation-object
type Operation struct {
	node
	Tags        []string
	Summary     string
	Description string
	Parameters  []Parameter          // can be a $ref
	RequestBody RequestBody          `yaml:"requestBody"` // can be a $ref
	Responses   map[string]*Response // can be a $ref
	//Callbacks 	map[string]*Callback // can be a $ref
	OperationId  string `yaml:"operationId"`
	ExternalDocs ExternalDocs
}

func (o *Operation) getRef() string {
	return ""
}

func (o *Operation) getChildren() map[string]Traversable {
	children := map[string]Traversable{}
	if o == nil {
		return children
	}
	// Parameters
	for i := range o.Parameters {
		parameter := o.Parameters[i]
		paramIndex := strconv.Itoa(i)
		children[paramIndex] = &parameter
	}

	// Request Body

	children["RequestBody"] = &o.RequestBody

	// Response
	for name := range o.Responses {
		response := o.Responses[name]
		children[name] = response
	}
	return children
}

func (o *Operation) setChild(i string, child Traversable) {
	switch child.(type) {
	case *Parameter:
		j, _ := strconv.Atoi(i)
		param, _ := child.(*Parameter)
		o.Parameters[j] = *param
		return
	case *RequestBody:
		requestBody, _ := child.(*RequestBody)
		o.RequestBody = *requestBody
		return
	case *Response:
		response, _ := child.(*Response)
		o.Responses[i] = response
		return
	default:
		panic("(o *OpenAPI) setChild(): " + errCastFail)
	}
}

func (o *Operation) GetName() string {
	name := o.getRenderer().sanitiseName(o.name) + o.parent.GetName()
	return name
}
