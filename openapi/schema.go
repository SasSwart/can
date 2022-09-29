package openapi

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

type Schema struct {
	Description          string
	Type                 string
	Properties           map[string]Schema
	Items                *Schema
	Ref                  string `yaml:"$ref"`
	AdditionalProperties bool
	Name                 string
	MinLength            int `yaml:"minLength"`
	MaxLength            int `yaml:"maxLength"`
	Pattern              string
	Format               string
	Required             []string
}

func (s *Schema) ResolveRefs(basePath string, components *Components) error {
	if s.Items != nil && s.Items.Ref != "" {
		ref := filepath.Join(basePath, s.Items.Ref)
		var newSchema Schema
		err := readRef(ref, &newSchema)
		if err != nil {
			return fmt.Errorf("Unable to read schema reference:\n%w", err)
		}

		newSchema.Name = refToName(ref)
		s.Items = &newSchema

		components.Schemas[newSchema.Name] = newSchema

		err = s.Items.ResolveRefs(basePath, components)
		if err != nil {
			return err
		}
	}

	if s.Properties != nil {
		for key, schema := range s.Properties {
			ref := schema.Ref
			err := schema.ResolveRefs(basePath, components)
			if err != nil {
				return err
			}
			schema.Name = refToName(ref)
			s.Properties[key] = schema
		}
	}
	return nil
}

func refToName(ref string) string {
	docsRoot := filepath.Dir(viper.GetString("openAPIFile"))
	name := strings.ReplaceAll(ref, docsRoot, "")
	name = strings.ReplaceAll(name, filepath.Ext(name), "")
	return name
}
