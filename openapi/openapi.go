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

	var resolveRefs = func(n traversable) (traversable, error) {
		node, ok := n.(refContainer)
		if !ok {
			return nil, fmt.Errorf("not a valid refContainer")
		}
		ref := node.getRef()
		if ref != "" {
			var err error
			switch n.(type) {
			case pathItem:
				err = readRef(node.getBasePath(), n)
			case *Schema:
				err = readRef(node.getBasePath(), n)
			}
			if err != nil {
				return nil, fmt.Errorf("Unable to read reference:\n%w", err)
			}
			return n, nil
		}

		return n, nil
	}

	// Resolve references
	newapi, err := traverse(&api, resolveRefs)
	if err != nil {
		return nil, err
	}

	return newapi.(*OpenAPI), err
}

type openAPIMeta struct {
	basePath string
}

func (m openAPIMeta) getBasePath() string {
	return m.basePath
}

// OpenAPI is a programmatic representation of the OpenApi Document object defined here: https://swagger.io/specification/#openapi-object
type OpenAPI struct {
	parent traversable
	openAPIMeta
	OpenAPI string `yaml:"openapi"`
	Info    Info
	//Servers Servers
	Servers    []Server // TODO fix bugs after this modification
	Paths      map[string]pathItem
	Components Components
}

func (o *OpenAPI) getParent() traversable {
	return nil
}

func (o *OpenAPI) getChildren() map[string]traversable {
	traversables := map[string]traversable{}
	for s, item := range o.Paths {
		traversables[s] = item
	}
	return traversables
}

func (o *OpenAPI) setChild(i string, child traversable) {
	o.Paths[i] = child.(pathItem)
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

// ExternalDocs is a programmatic representation of the External Docs object defined here: https://swagger.io/specification/#external-documentation-object
type ExternalDocs struct {
	Description string `yaml:"description"`
	Url         string `yaml:"url"`
}
