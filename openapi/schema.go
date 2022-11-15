package openapi

import (
	"path/filepath"
)

// TODO see if this can be made spec-compliant while retaining original logical flow

var _ refContainer = &Schema{}

// Schema is a programmatic representation of the Schema object defined here: https://swagger.io/specification/#schema-object
type Schema struct {
	parent               refContainer
	name                 string
	Description          string
	Type                 string
	Properties           map[string]Schema
	Items                *Schema
	Ref                  string `yaml:"$ref"`
	AdditionalProperties bool
	MinLength            int `yaml:"minLength"`
	MaxLength            int `yaml:"maxLength"`
	Pattern              string
	Format               string
	Required             []string
}

func (s *Schema) GetName() string {
	return s.parent.GetName() + s.name
}

func (s *Schema) getChildren() map[string]Traversable {
	children := map[string]Traversable{}
	for name, property := range s.Properties {
		children[name] = &property
	}
	return children
}

func (s *Schema) setChild(i string, t Traversable) {
	// TODO: handle this error
	schema, _ := t.(*Schema)
	s.Properties[i] = *schema
}

func (s *Schema) getBasePath() string {
	basePath := filepath.Join(s.parent.getBasePath(), filepath.Dir(s.Ref))
	return basePath
}

func (s *Schema) getRef() string {
	return s.Ref
}
