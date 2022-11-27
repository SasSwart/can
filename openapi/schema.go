package openapi

import (
	"path/filepath"
)

// TODO see if this can be made spec-compliant while retaining original logical flow

var _ Traversable = &Schema{}

// Schema is a programmatic representation of the Schema object defined here: https://swagger.io/specification/#schema-object
type Schema struct {
	node
	Description          string
	Type                 string
	Properties           map[string]*Schema
	Items                *Schema
	Ref                  string `yaml:"$ref"`
	AdditionalProperties bool
	MinLength            int `yaml:"minLength"`
	MaxLength            int `yaml:"maxLength"`
	Pattern              string
	Format               string
	Required             []string
}

func (s *Schema) GetType() string {
	return s.renderer.sanitiseType(s)
}

func (s *Schema) getChildren() map[string]Traversable {
	children := map[string]Traversable{}
	for name := range s.Properties {
		property := s.Properties[name]
		children[name] = property
	}
	if s.Items != nil {
		children["item"] = s.Items
	}
	return children
}

func (s *Schema) setChild(i string, t Traversable) {
	if schema, ok := t.(*Schema); ok {
		if i == "item" {
			s.Items = schema
		} else {
			s.Properties[i] = schema
		}
		return
	}
	panic("(s *Schema) setChild borked")
}

func (s *Schema) getBasePath() string {
	if s.parent == nil {
		return ""
	}
	basePath := filepath.Join(s.parent.getBasePath(), filepath.Dir(s.Ref))
	return basePath
}

//func (s *Schema) getRef() string {
//	panic("(s *Schema) getRef() This should never be called")
//	return s.Ref
//}
