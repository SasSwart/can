package schema

import (
	"github.com/sasswart/gin-in-a-can/errors"
	"github.com/sasswart/gin-in-a-can/tree"
	"path/filepath"
)

const Key = "Model"
const SubSchemaKey = "item"

var _ tree.NodeTraverser = &Schema{}

// Properties uses numerical index values converted to the string type
type Properties map[string]*Schema

// Schema is a programmatic representation of the Schema object defined here: https://swagger.io/specification/#schema-object
type Schema struct {
	tree.Node
	Ref                  string `yaml:"$ref"`
	Description          string
	Type                 string
	Properties           Properties
	Items                *Schema
	AdditionalProperties bool
	MinLength            int `yaml:"minLength"`
	MaxLength            int `yaml:"maxLength"`
	Pattern              string
	Format               string
	Required             []string
}

func (s *Schema) GetChildren() map[string]tree.NodeTraverser {
	children := map[string]tree.NodeTraverser{}
	for name := range s.Properties {
		property := s.Properties[name]
		children[name] = property
	}
	if s.Items != nil {
		children[SubSchemaKey] = s.Items
	}
	return children
}

func (s *Schema) SetChild(i string, t tree.NodeTraverser) {
	if schema, ok := t.(*Schema); ok {
		if i == SubSchemaKey {
			s.Items = schema
		} else {
			if s.Properties == nil {
				s.Properties = make(Properties, 4)
			}
			s.Properties[i] = schema
		}
		return
	}
	errors.CastFail("(s *Schema) SetChild()", "NodeTraverser", "*schema.Schema")
}

// GetBasePath should be defined on any function that could need referential resolution
func (s *Schema) GetBasePath() string {
	if s.GetParent() == nil {
		return ""
	}
	basePath := filepath.Join(s.GetParent().GetBasePath(), filepath.Dir(s.Ref))
	return basePath
}

func (s *Schema) GetRef() string {
	return s.Ref
}

func (s *Schema) IsRequired(property string) bool {
	if s.Required == nil {
		return false
	}

	for _, item := range s.Required {
		if item == property {
			return true
		}
	}
	return false
}
