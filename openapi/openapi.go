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
		Paths: map[string]Path{},
	}
	content, err := os.ReadFile(openAPIFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read file \"%s\": %w", openAPIFile, err)
	}
	yaml.Unmarshal(content, &api)

	err = api.ResolveRefs(path.Dir(openAPIFile))
	return &api, err
}

// OpenAPI is a programmatic representation of the OpenApi Document object defined here: https://swagger.io/specification/#openapi-object
type OpenAPI struct {
	OpenAPI    string `yaml:"openapi"`
	Info       Info
	Servers    []Server // TODO fix bugs after this modification
	Paths      Paths
	Components Components
}

func (o *OpenAPI) ResolveRefs(basePath string) error {
	return o.Paths.ResolveRefs(basePath)
}

// ExternalDocs is a programmatic representation of the External Docs object defined here: https://swagger.io/specification/#external-documentation-object
type ExternalDocs struct {
	Description string `yaml:"description"`
	Url         string `yaml:"url"`
}
