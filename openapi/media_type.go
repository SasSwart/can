package openapi

import (
	"fmt"
	"path"
)

type MediaType struct {
	Schema Schema
}

var _ node = &MediaType{}

func (m *MediaType) ResolveRefs(basePath string) error {
	if m.Schema.Ref == "" {
		return nil
	}

	ref := m.Schema.Ref

	var newSchema Schema
	err := readRef(path.Join(basePath, m.Schema.Ref), &newSchema)
	if err != nil {
		return fmt.Errorf("Unable to read schema reference:\n%w", err)
	}
	m.Schema = newSchema

	refBasePath := path.Dir(ref)
	err = newSchema.ResolveRefs(path.Join(basePath, refBasePath))
	if err != nil {
		return err
	}

	return nil
}
