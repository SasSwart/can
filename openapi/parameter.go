package openapi

// Parameter is a programmatic representation of the Parameter object defined here: https://swagger.io/specification/#parameter-object
type Parameter struct {
	Name            string `yaml:"name"`
	In              string `yaml:"in"`
	Description     string `yaml:"description"`
	Required        bool   `yaml:"required"`
	Deprecated      bool   `yaml:"deprecated"`
	AllowEmptyValue bool   `yaml:"allowEmptyValue"`
	Schema          Schema // Acts as alternative description of param
}

func (p *Parameter) ResolveRefs(basePath string) error {
	return nil
}
