package openapi

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type node interface {
	ResolveRefs(basePath string) error
}

func LoadOpenAPI(path string) (*OpenAPI, error) {
	api := OpenAPI{}
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %w", err)
	}
	yaml.Unmarshal(content, &api)

	return &api, nil
}

type OpenAPI struct {
	OpenAPI string
	Info    Info
	Servers Servers
	Paths   Paths
}

func (o *OpenAPI) ResolveRefs(basePath string) error {
	return o.Paths.ResolveRefs(basePath)
}

type ExternalDocs struct {
}
