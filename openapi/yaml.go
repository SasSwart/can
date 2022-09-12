package openapi

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type OpenAPI struct {
	OpenAPI string
	Info    Info
	Paths   Paths
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
	RequestBody  RequestBody
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

type RequestBody struct{}

type Response struct {
	Description string
	Content     map[string]MediaType
}

type MediaType struct {
	Schema Schema
}

type Schema struct {
	Type string
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
