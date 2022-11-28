package openapi

var _ Traversable = &MediaType{}

// MediaType is a programmatic representation of the MediaType object defined here: https://swagger.io/specification/#media-type-object
type MediaType struct {
	node
	name   string
	Schema *Schema
}

func (m *MediaType) getRef() string {
	return ""
}

// GetName overrides the default node implementation to avoid
func (m *MediaType) GetName() string {
	return m.parent.GetName() + m.name
}

func (m *MediaType) getChildren() map[string]Traversable {
	return map[string]Traversable{
		"Model": m.Schema,
	}
}

func (m *MediaType) setChild(_ string, t Traversable) {
	if s, ok := t.(*Schema); ok {
		m.Schema = s
		return
	}
	panic("(m *MediaType) setChild borked")
}
