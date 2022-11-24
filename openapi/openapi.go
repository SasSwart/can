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

func LoadOpenAPI(openAPIFile string) (*OpenAPI, error) {
	// skeleton
	api := OpenAPI{
		node: node{
			basePath: filepath.Dir(openAPIFile),
		},
		Components: Components{
			Schemas: map[string]Schema{},
		},
		Paths: map[string]*PathItem{},
	}

	// Read yaml file
	content, err := os.ReadFile(openAPIFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read file \"%s\": %w", openAPIFile, err)
	}
	yaml.Unmarshal(content, &api)

	api.setName(api.Info.Title)

	// Resolve references
	newapi, err := Traverse(&api, resolveRefs)

	if err != nil {
		return nil, err
	}

	return newapi.(*OpenAPI), err
}

func SetRenderer(api *OpenAPI, renderer Renderer) {
	Traverse(api, func(_ string, _, child Traversable) (Traversable, error) {
		child.setRenderer(renderer)
		parent := child.GetParent()
		if parent != nil {
			parent.setRenderer(renderer)
		}

		return child, nil
	})
}

// resolveRefs walks the tree and
func resolveRefs(key string, parent, child Traversable) (Traversable, error) {
	child.setParent(parent)
	child.setName(key)
	ref := child.getRef()
	if ref != "" {
		basePath := child.getBasePath()
		ref := filepath.Base(child.getRef())
		err := readRef(filepath.Join(basePath, ref), child)
		if err != nil {
			return nil, fmt.Errorf("Unable to read reference:\n%w", err)
		}
	}

	return child, nil
}

// OpenAPI is a programmatic representation of the OpenApi Document object defined here: https://swagger.io/specification/#openapi-object
type OpenAPI struct {
	node
	OpenAPI string `yaml:"openapi"`
	Info    Info
	//Servers Servers
	Servers    []Server // TODO fix bugs after this modification
	Paths      map[string]*PathItem
	Components Components
}

func (o *OpenAPI) getBasePath() string {
	return o.basePath
}

func (o *OpenAPI) GetName() string {
	return o.name
}

func (o *OpenAPI) GetOutputFile() string {
	return filepath.Join(o.getRenderer().getOutputFile(o), o.GetName())
}

func (o *OpenAPI) getChildren() map[string]Traversable {
	traversables := map[string]Traversable{}
	for s := range o.Paths {
		path := o.Paths[s]
		traversables[s] = path
	}
	return traversables
}

func (o *OpenAPI) setChild(i string, child Traversable) {
	c, _ := child.(*PathItem)
	o.Paths[i] = c
}

// ExternalDocs is a programmatic representation of the External Docs object defined here: https://swagger.io/specification/#external-documentation-object
type ExternalDocs struct {
	Description string `yaml:"description"`
	Url         string `yaml:"url"`
}
