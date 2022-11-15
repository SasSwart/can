package openapi

// Operation is a programmatic representation of the Operation object defined here: https://swagger.io/specification/#operation-object
type Operation struct {
	parent       refContainer
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

// communicate by sharing memory ;)
var _ refContainer = &Operation{}

func (o *Operation) getParent() Traversable {
	return nil
}

func (o *Operation) getChildren() map[string]Traversable {
	children := map[string]Traversable{}
	if o == nil {
		return children
	}
	//for _, parameter := range o.Parameters {
	//	children[parameter.Name] = parameter
	//}
	children["RequestBody"] = &o.RequestBody
	for name, response := range o.Responses {
		children[name] = &response
	}
	return children
}

func (o *Operation) setChild(i string, child Traversable) {
	// TODO
}
