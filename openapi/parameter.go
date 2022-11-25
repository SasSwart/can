package openapi

var _ Traversable = &Parameter{}

// Parameter is a programmatic representation of the Parameter object defined here: https://swagger.io/specification/#parameter-object
type Parameter struct {
	node
	Name            string `yaml:"name"`
	In              string `yaml:"in"`
	Description     string `yaml:"description"`
	Required        bool   `yaml:"required"`
	Deprecated      bool   `yaml:"deprecated"`
	AllowEmptyValue bool   `yaml:"allowEmptyValue"`
	Schema          Schema // Acts as alternative description of param
}

func (p *Parameter) getRef() string {
	return ""
}

func (p *Parameter) getChildren() map[string]Traversable {
	return map[string]Traversable{
		"model": &p.Schema,
	}
}

func (p *Parameter) setChild(i string, t Traversable) {
	// TODO: Handle this error
	schema, _ := t.(*Schema)
	p.Schema = *schema
}

func (p *Parameter) ResolveRefs(basePath string) error {
	return nil
}
