package openapi

import (
	"fmt"
	"path/filepath"
	"strings"
)

// TODO see if this can be made spec-compliant while retaining original logical flow

var _ refContainer = &Schema{}

// Schema is a programmatic representation of the Schema object defined here: https://swagger.io/specification/#schema-object
type Schema struct {
	parent               traversable
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

func (s *Schema) getParent() traversable {
	//TODO implement me
	panic("implement me")
}

func (s *Schema) getChildren() map[string]traversable {
	//TODO implement me
	panic("implement me")
}

func (s *Schema) setChild(i string, t traversable) {
	//TODO implement me
	panic("implement me")
}

func (s *Schema) getBasePath() string {
	//TODO implement me
	panic("implement me")
}

func (s *Schema) getRef() string {
	//TODO implement me
	panic("implement me")
}

func (s *Schema) ResolveRefs() (err error) {
	ref, err := filepath.Abs(s.getBasePath())
	if err != nil {
		return err
	}

	if s.Ref != "" {
		splitRef := strings.Split(s.Ref, "#")
		var file string
		if len(splitRef) == 1 {
			file = splitRef[0]
		} else {
			file, _ = splitRef[0], splitRef[1]
		}

		ref = filepath.Join(s.getBasePath(), file)

		err = readRef(ref, s)
		if err != nil {
			return err
		}
		s.Ref = ref
	}

	return nil
}

func (s *Schema) Render() (err error) {
	if s == nil {
		return nil
	}
	fmt.Println("Rendering API Schema")

	if s.Items != nil && s.Items.Ref != "" {
		err = s.Items.Render()
		if err != nil {
			return err
		}
	}

	if s.Properties != nil {
		for _, schema := range s.Properties {
			err := schema.Render()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Schema) GetSchemas(name string) (schemas map[string]Schema) {
	schemas = map[string]Schema{}
	if s == nil {
		return schemas
	}

	schemas[name] = *s

	s.Items.GetSchemas(name + "Item")

	for name, schema := range s.Properties {
		for name, subSchema := range schema.GetSchemas(name) {
			schemas[name] = subSchema
		}
	}

	return schemas
}
