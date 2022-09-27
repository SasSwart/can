package openapi

import (
	"fmt"
	"path/filepath"
)

type Schema struct {
	Description          string
	Type                 string
	Properties           map[string]Schema
	Items                *Schema
	Ref                  string `yaml:"$ref"`
	AdditionalProperties bool
}

func (s *Schema) ResolveRefs(basePath string, components *Components) error {
	if s.Ref != "" {
		s.Ref = filepath.Join(basePath, s.Ref)

		var newSchema Schema
		err := readRef(s.Ref, &newSchema)
		if err != nil {
			return fmt.Errorf("Unable to read schema reference:\n%w", err)
		}
		components.Schemas[s.Ref] = newSchema

		err = newSchema.ResolveRefs(basePath, components)
		if err != nil {
			return err
		}
	}

	if s.Items != nil {
		err := s.Items.ResolveRefs(basePath, components)
		if err != nil {
			return err
		}
	}

	if s.Properties != nil {
		for _, schema := range s.Properties {
			err := schema.ResolveRefs(basePath, components)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
