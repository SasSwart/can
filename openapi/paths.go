package openapi

// Path is a programmatic representation of the Path Item object defined here: https://swagger.io/specification/#path-item-object
type Path struct {
	Summary     string
	Description string
	Get         *Operation
	Post        *Operation
	Patch       *Operation
	Delete      *Operation
	Parameters  []Parameter
	Ref         string `yaml:"$ref"`
}

func (p *Path) ResolveRefs(basePath string) error {
	for _, operation := range p.Operations() {
		if operation == nil {
			continue
		}
		err := operation.ResolveRefs(basePath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Path) GetSchemas(name string) (schemas map[string]Schema) {
	schemas = map[string]Schema{}
	for _, operation := range p.Operations() {
		for name, schema := range operation.GetSchemas(name) {
			schemas[name] = schema
		}
	}

	return schemas
}

func (p *Path) Operations() map[string]*Operation {
	return map[string]*Operation{
		"delete": p.Delete,
		"get":    p.Get,
		"patch":  p.Patch,
		"post":   p.Post,
	}
}
