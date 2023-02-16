package openapi

var _ Traversable = &Parameter{}

// Parameter is a programmatic representation of the Parameter object defined here: https://swagger.io/specification/#parameter-object
type Parameter struct {
	node
	Ref             string  `yaml:"$ref"`
	Name            string  `yaml:"name"`
	In              string  `yaml:"in"`
	Description     string  `yaml:"description"`
	Required        bool    `yaml:"required"`
	Deprecated      bool    `yaml:"deprecated"`
	AllowEmptyValue bool    `yaml:"allowEmptyValue"`
	Schema          *Schema // Acts as alternative description of param
}

func (p *Parameter) getRef() string {
	return p.Ref
}

func (p *Parameter) getChildren() map[string]Traversable {
	return map[string]Traversable{
		"Model": p.Schema,
	}
}

func (p *Parameter) GetOutputFile() string {
	// TODO this override can be removed and this logic dealt with by the node{} composable
	return p.getRenderer().getOutputFile(p)
}

func (p *Parameter) GetName() string {
	// TODO can this be done through switch node.parent.(Type)?
	name := p.parent.GetName() + p.getRenderer().sanitiseName(p.Name) + "Parameter"
	return name
}

func (p *Parameter) setChild(_ string, t Traversable) {
	schema, ok := t.(*Schema)
	if !ok {
		panic("(p *Parameter) setChild(): " + errCastFail)
	}
	p.Schema = schema
}
