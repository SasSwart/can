package openapi

import "strconv"

// communicate by sharing memory ;)
var _ refContainer = &Operation{}

// Operation is a programmatic representation of the Operation object defined here: https://swagger.io/specification/#operation-object
type Operation struct {
	refContainerNode
	Tags         []string
	Summary      string
	Description  string
	Parameters   []Parameter
	RequestBody  RequestBody `yaml:"requestBody"`
	Responses    map[string]Response
	OperationId  string `yaml:"operationId"`
	ExternalDocs ExternalDocs
}

func (o *Operation) getBasePath() string {
	return o.parent.getBasePath()
}

func (o *Operation) getRef() string {
	return ""
}

func (o *Operation) getParent() Traversable {
	return nil
}

func (o *Operation) getChildren() map[string]Traversable {
	children := map[string]Traversable{}
	if o == nil {
		return children
	}
	for i := range o.Parameters {
		parameter := o.Parameters[i]
		children[string(i)] = &parameter
	}
	children["RequestBody"] = &o.RequestBody
	for name := range o.Responses {
		response := o.Responses[name]
		children[name] = &response
	}
	return children
}

func (o *Operation) setChild(i string, child Traversable) {
	switch child.(type) {
	case *Parameter:
		// TODO: Handle this error
		j, _ := strconv.Atoi(i)
		param, _ := child.(*Parameter)
		o.Parameters[j] = *param
	case *RequestBody:
		requestBody, _ := child.(*RequestBody)
		o.RequestBody = *requestBody
	case *Response:
		response, _ := child.(*Response)
		o.Responses[i] = *response
	}
}

type operationChildNode node[*Operation]

func (o operationChildNode) getBasePath() string {
	return o.parent.getBasePath()
}

func (o operationChildNode) GetName() string {
	return o.parent.GetName() + o.name
}
