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

	var resolveRefs = func(n refContainer) refContainer {
		ref := n.getRef()
		if ref != "" {
			var err error
			switch n.(type) {
			case pathItem:
				err = readRef(n.getBasePath(), n)
			case Schema:
				err = readRef(n.getBasePath(), n)
			}
			if err != nil {
				return fmt.Errorf("Unable to read reference:\n%w", err)
			}
			return pathItem
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

// OpenAPI is a programmatic representation of the OpenApi Document object defined here: https://swagger.io/specification/#openapi-object
type OpenAPI struct {
	parent traversable
	openAPIMeta
	OpenAPI string `yaml:"openapi"`
	Info    Info
	//Servers Servers
	Servers    []Server // TODO fix bugs after this modification
	Paths      Paths
	Components Components
}

func (o *OpenAPI) getParent() traversable {
	return nil
}

func (o *OpenAPI) getChildren() childContainer[string] {
	return childContainerMap[string]{
		o.Paths,
	}
}

func (o *OpenAPI) setChild(i string, child traversable) {
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

// ExternalDocs is a programmatic representation of the External Docs object defined here: https://swagger.io/specification/#external-documentation-object
type ExternalDocs struct {
	Description string `yaml:"description"`
	Url         string `yaml:"url"`
}
