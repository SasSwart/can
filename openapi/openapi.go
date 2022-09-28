package openapi

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path"
)

type node interface {
	ResolveRefs(basePath string, components *Components) error
}

func LoadOpenAPI(openAPIFile string) (*OpenAPI, error) {
	api := OpenAPI{
		Components: Components{
			Schemas: map[string]Schema{},
		},
		Paths: map[string]PathItem{},
	}
	content, err := os.ReadFile(openAPIFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %w", err)
	}
	yaml.Unmarshal(content, &api)

	err = api.ResolveRefs(path.Dir(openAPIFile), &api.Components)
	return &api, err
}

type OpenAPI struct {
	OpenAPI    string
	Info       Info
	Servers    Servers
	Paths      Paths
	Components Components
}

func (o *OpenAPI) ResolveRefs(basePath string, components *Components) error {
	return o.Paths.ResolveRefs(basePath, components)
}

type ExternalDocs struct {
}
