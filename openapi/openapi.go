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
	traversable
	ResolveRefs() error
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
		Paths: map[string]pathItem{},
	}

	// Read yaml file
	content, err := os.ReadFile(openAPIFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read file \"%s\": %w", openAPIFile, err)
	}
	yaml.Unmarshal(content, &api)

	// Populate node children
	for name, item := range api.getChildren() {
		item.parent = &api
		api.children = append(api.children, &item)
	}

	var resolveRefs = func(n refContainer) refContainer {
		ref := n.getRef()
		if ref != "" {
			switch n.(type) {
			case pathItem:
				err := readRef(n.getBasePath(), n)
			case Schema:
				err := readRef(n.getBasePath(), n)
			}
			if err != nil {
				return fmt.Errorf("Unable to read reference:\n%w", err)
			}
			return pathItem
		}

		return n
	}

	var renderAll = func(n traversable) traversable {
		err := n.Render()
		if err != nil {
			return fmt.Errorf("Unable to read reference:\n%w", err)
		}

		return n
	}

	// Resolve references
	newapi := traverse(&api, resolveRefs).(*OpenAPI)

	return newapi, err
}

type openAPIMeta struct {
	basePath string
}

func (m openAPIMeta) getBasePath() string {
	return m.basePath
}

type OpenAPI struct {
	parent traversable[int]
	openAPIMeta
	OpenAPI    string
	Info       Info
	Servers    Servers
	Paths      map[string]pathItem
	Components Components
}

func (o *OpenAPI) getChildren() map[string]traversable[any] {
	return o.Paths
}

func (o *OpenAPI) setChild(i string, child traversable[any]) {
	o.Paths[i] = child
}

func (o *OpenAPI) Render() error {
	// TODO: redefine this in terms of the traverse method
	fmt.Println("Rendering API Spec")
	return nil
}

//
//func (o *OpenAPI) GetSchemas(name string) map[string]Schema {
//	return o.paths.GetSchemas(name)
//}

type ExternalDocs struct {
}
