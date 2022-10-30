package openapi

type MediaType struct {
	Schema *Schema
}

func (m *MediaType) ResolveRefs(basePath string) error {
	if m.Schema == nil {
		return nil
	}

	err := m.Schema.ResolveRefs(basePath)
	if err != nil {
		return err
	}

	return nil
}
