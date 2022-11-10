package openapi

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path"
)

type Config struct {
	OpenAPIFile string
}

type node interface {
	ResolveRefs(basePath string) error
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
		return nil, fmt.Errorf("unable to read file \"%s\": %w", openAPIFile, err)
	}
	yaml.Unmarshal(content, &api)

	err = api.ResolveRefs(path.Dir(openAPIFile))
	return &api, err
}

type OpenAPI struct {
	OpenAPI    string
	Info       Info
	Servers    Servers
	Paths      Paths
	Components Components
}

func (o *OpenAPI) ResolveRefs(basePath string) error {
	return o.Paths.ResolveRefs(basePath)
}

func (o *OpenAPI) GetSchemas(name string) map[string]Schema {
	return o.Paths.GetSchemas(name)
}

type ExternalDocs struct {
}
