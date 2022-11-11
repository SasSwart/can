package openapi

// Operation is a programmatic representation of the Operation object defined here: https://swagger.io/specification/#operation-object
type Operation struct {
	Tags         []string
	Summary      string
	Description  string
	Parameters   []Parameter
	RequestBody  RequestBody `yaml:"requestBody"`
	Responses    map[string]Response
	OperationId  string `yaml:"operationId"`
	ExternalDocs ExternalDocs
}

func (o *Operation) ResolveRefs(basePath string) error {
	err := o.RequestBody.ResolveRefs(basePath)
	if err != nil {
		return err
	}

	for key, response := range o.Responses {
		err := response.ResolveRefs(basePath)
		if err != nil {
			return err
		}
		o.Responses[key] = response
	}

	return nil
}

func (o *Operation) GetSchemas(name string) (schemas map[string]Schema) {
	schemas = map[string]Schema{}
	if o == nil {
		return schemas
	}
	for _, response := range o.Responses {
		for name, schema := range response.GetSchemas(name) {
			schemas[name] = schema
		}
	}

	return schemas
}
