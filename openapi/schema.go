package openapi

import (
	"path/filepath"
	"strings"
)

// Schema is a programmatic representation of the Schema object defined here: https://swagger.io/specification/#schema-object
// TODO see if this can be made spec-compliant while retaining original logical flow
type Schema struct {
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

func (s *Schema) ResolveRefs(basePath string) error {
	if s.Ref != "" {
		splitRef := strings.Split(s.Ref, "#")
		var file string
		if len(splitRef) == 1 {
			file = splitRef[0]
		} else {
			file, _ = splitRef[0], splitRef[1]
		}

		ref, err := filepath.Abs(filepath.Join(basePath, file))
		if err != nil {
			return err
		}

		err = readRef(ref, &s)
		if err != nil {
			return err
		}
		s.Ref = ref
	}

	basePath = filepath.Dir(s.Ref)

	if s.Items != nil && s.Items.Ref != "" {
		err := s.Items.ResolveRefs(basePath)
		if err != nil {
			return err
		}
	}

	if s.Properties != nil {
		for key, schema := range s.Properties {
			err := schema.ResolveRefs(basePath)
			if err != nil {
				return err
			}
			s.Properties[key] = schema
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

	for name, schema := range s.Properties {
		for name, subSchema := range schema.GetSchemas(name) {
			schemas[name] = subSchema
		}
	}

	return schemas
}
