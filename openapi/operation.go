package openapi

type Operation struct {
	Tags         []string
	Description  string
	Parameters   []Parameter
	RequestBody  RequestBody `yaml:"requestBody"`
	Responses    map[string]Response
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
