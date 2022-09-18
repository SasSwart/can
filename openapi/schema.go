package openapi

type Schema struct {
	Description          string
	Type                 string
	Properties           map[string]Schema
	Items                *Schema
	Ref                  string `yaml:"$ref""`
	AdditionalProperties bool
}

func (s *Schema) ResolveRefs(basePath string) error {
	if s.Ref == "" {
		return nil
	}

	return nil
}
