package openapi

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path"
)

type OpenAPI struct {
	OpenAPI string
	Info    Info
	Paths   Paths
}

func (o *OpenAPI) ResolveRefs(basePath string) error {
	for key, pathItem := range o.Paths {
		if pathItem.Ref == "" {
			continue
		}
		content, err := os.ReadFile(path.Join(basePath, pathItem.Ref))
		if err != nil {
			return fmt.Errorf("unable to resolve PathItem Reference: %w", err)
		}
		var newPathItem PathItem
		yaml.Unmarshal(content, &newPathItem)
		o.Paths[key] = newPathItem

		refBasePath := path.Dir(pathItem.Ref)
		err = newPathItem.ResolveRefs(path.Join(basePath, refBasePath))
		if err != nil {
			return err
		}
	}
	return nil
}

type Info struct {
	Title   string
	Version string
}

type Paths map[string]PathItem

type PathItem struct {
	Summary     string
	Description string
	Get         Operation
	Post        Operation
	Patch       Operation
	Delete      Operation
	Parameters  []Parameter
	Ref         string `yaml:"$ref""`
}

func (p *PathItem) ResolveRefs(basePath string) error {
	return nil
}

func (p PathItem) Operations() map[string]Operation {
	return map[string]Operation{
		"delete": p.Delete,
		"get":    p.Get,
		"patch":  p.Patch,
		"post":   p.Post,
	}
}

type Operation struct {
	Tags         []string
	Description  string
	Parameters   []Parameter
	RequestBody  RequestBody `yaml:"requestBody"`
	Responses    map[string]Response
	ExternalDocs ExternalDocs
}

type Parameter struct {
	Name            string
	In              string
	Description     string
	Required        bool
	AllowEmptyValue bool
	Schema          Schema
}

type RequestBody struct {
	Description string
	Content     map[string]MediaType
	Required    bool
}

type Response struct {
	Description string
	Content     map[string]MediaType
}

type MediaType struct {
	Schema Schema
}

type Schema struct {
	Description          string
	Type                 string
	Properties           map[string]Schema
	Items                *Schema
	Ref                  string
	AdditionalProperties bool
}

type ExternalDocs struct {
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
