package openapi

import "fmt"

type MediaType struct {
	Schema *Schema
}

func (m *MediaType) Render() error {
	fmt.Println("Rendering API Media Type")

	if m.Schema == nil {
		return nil
	}

	err := m.Schema.Render()
	if err != nil {
		return err
	}

	return nil
}
