package schema

import (
	"github.com/sasswart/gin-in-a-can/openapi/errors"
	"github.com/sasswart/gin-in-a-can/render"
	"github.com/sasswart/gin-in-a-can/tree"
	"path/filepath"
)

var _ tree.NodeTraverser = &Schema{}

// Schema is a programmatic representation of the Schema object defined here: https://swagger.io/specification/#schema-object
type Schema struct {
	tree.Node
	Ref                  string `yaml:"$ref"`
	Description          string
	Type                 string
	Properties           map[string]*Schema
	Items                *Schema
	AdditionalProperties bool
	MinLength            int `yaml:"minLength"`
	MaxLength            int `yaml:"maxLength"`
	Pattern              string
	Format               string
	Required             []string
}

func (s *Schema) GetRenderer() render.Renderer {
	//TODO implement me
	panic("implement me")
}

func (s *Schema) GetType() string {
	renderer := s.GetRenderer()
	sanitisedType := renderer.SanitiseType(s)
	return sanitisedType
}

func (s *Schema) GetChildren() map[string]tree.NodeTraverser {
	children := map[string]tree.NodeTraverser{}
	for name := range s.Properties {
		property := s.Properties[name]
		children[name] = property
	}
	if s.Items != nil {
		children["item"] = s.Items
	}
	return children
}

func (s *Schema) SetChild(i string, t tree.NodeTraverser) {
	if schema, ok := t.(*Schema); ok {
		if i == "item" {
			s.Items = schema
		} else {
			s.Properties[i] = schema
		}
		return
	}
	panic("(s *Schema) setChild(): " + errors.ErrCastFail)
}

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
