package openapi

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Config struct {
	OpenAPIFile string
}

type refContainer interface {
	Traversable
	getBasePath() string
	getRef() string
}

func LoadOpenAPI(openAPIFile string) (*OpenAPI, error) {
	// skeleton
	api := OpenAPI{
		openAPIMeta: openAPIMeta{basePath: filepath.Dir(openAPIFile)},
		Components: Components{
			Schemas: map[string]Schema{},
		},
		Paths: map[string]PathItem{},
	}

	// Read yaml file
	content, err := os.ReadFile(openAPIFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read file \"%s\": %w", openAPIFile, err)
	}
	yaml.Unmarshal(content, &api)

	// Resolve references
	newapi, err := Traverse(&api, resolveRefs)
	if err != nil {
		return nil, err
	}

	return newapi.(*OpenAPI), err
}

func resolveRefs(parent, child Traversable) (Traversable, error) {
	node, ok := child.(refContainer)
	if !ok {
		return child, nil
	}

	var err error
	switch child.(type) {
	case *PathItem:
		pathItemChild := child.(*PathItem)
		pathItemChild.parent = parent.(*OpenAPI)
		ref := node.getRef()
		if ref != "" {
			basePath := pathItemChild.getBasePath()
			ref := filepath.Base(pathItemChild.Ref)
			err = readRef(filepath.Join(basePath, ref), pathItemChild)
			if err != nil {
				return nil, fmt.Errorf("Unable to read reference:\n%w", err)
			}
		}
		return pathItemChild, nil
	case *Operation:
		operationChild := child.(*Operation)
		if operationChild == nil {
			return child, nil
		}
		operationChild.parent = parent.(refContainer)
		return operationChild, nil
	case *RequestBody:
		requestBodyChild := child.(*RequestBody)
		requestBodyChild.parent = parent.(refContainer)
		return requestBodyChild, nil
	case *Response:
		responseChild := child.(*Response)
		responseChild.parent = parent.(refContainer)
		return responseChild, nil
	case *MediaType:
		mediaTypeChild := child.(*MediaType)
		mediaTypeChild.parent = parent.(refContainer)
		return mediaTypeChild, nil
	case *Schema:
		schemaChild := child.(*Schema)
		schemaChild.parent = parent.(refContainer)
		ref := node.getRef()
		if ref != "" {
			basePath := schemaChild.getBasePath()
			ref := filepath.Base(schemaChild.Ref)
			err = readRef(filepath.Join(basePath, ref), schemaChild)
			if err != nil {
				return nil, fmt.Errorf("Unable to read reference:\n%w", err)
			}
		}
		return schemaChild, nil
	}

	return child, nil
}

type openAPIMeta struct {
	parent   Traversable
	basePath string
}

func (m openAPIMeta) getBasePath() string {
	return m.basePath
}

// OpenAPI is a programmatic representation of the OpenApi Document object defined here: https://swagger.io/specification/#openapi-object
type OpenAPI struct {
	openAPIMeta
	OpenAPI    string `yaml:"openapi"`
	Info       Info
	Servers    []Server
	Paths      map[string]PathItem
	Components Components
}

func (o *OpenAPI) getParent() Traversable {
	return nil
}

func (o *OpenAPI) getChildren() map[string]Traversable {
	traversables := map[string]Traversable{}
	for s, item := range o.Paths {
		traversables[s] = &item
	}
	return traversables
}

func (o *OpenAPI) setChild(i string, child Traversable) {
	c, _ := child.(*PathItem)
	o.Paths[i] = *c
}

// ExternalDocs is a programmatic representation of the External Docs object defined here: https://swagger.io/specification/#external-documentation-object
type ExternalDocs struct {
	Description string `yaml:"description"`
	Url         string `yaml:"url"`
}
