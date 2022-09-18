package openapi

type Parameter struct {
	Name            string
	In              string
	Description     string
	Required        bool
	AllowEmptyValue bool
	Schema          Schema
}

func (p *Parameter) ResolveRefs(basePath string) error {
	return nil
}
