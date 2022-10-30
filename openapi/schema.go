package openapi

import (
	"path/filepath"
)

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
	ref, err := filepath.Abs(filepath.Join(basePath, s.Ref))
	if err != nil {
		return err
	}

	err = readRef(ref, &s)
	if err != nil {
		return err
	}
	s.Ref = ref

	basePath = filepath.Dir(s.Ref)

	if s.Items != nil && s.Items.Ref != "" {
		err = s.Items.ResolveRefs(basePath)
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
