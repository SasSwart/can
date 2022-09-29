package openapi

import (
	"fmt"
	"path"
)

type MediaType struct {
	Schema Schema
}

func (m *MediaType) ResolveRefs(basePath string, components *Components) error {
	if m.Schema.Ref == "" {
		return nil
	}

	ref := m.Schema.Ref

	fullRefPath := path.Join(basePath, m.Schema.Ref)
	var newSchema Schema
	err := readRef(fullRefPath, &newSchema)
	if err != nil {
		return fmt.Errorf("Unable to read schema reference:\n%w", err)
	}
	newSchema.Name = refToName(fullRefPath)
	m.Schema = newSchema
	components.Schemas[newSchema.Name] = newSchema

	refBasePath := path.Dir(ref)
	err = newSchema.ResolveRefs(path.Join(basePath, refBasePath), components)
	if err != nil {
		return err
	}

	return nil
}
