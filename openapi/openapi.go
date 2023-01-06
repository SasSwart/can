package openapi

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Renderer interface {
	SanitiseName(string) string
	SanitiseType(*Schema) string
	GetOutputFile(Traversable) string
}
type Config struct {
	OpenAPIFile string
}

// OpenAPI is a programmatic representation of the OpenApi Document object defined here: https://swagger.io/specification/#openapi-object
type OpenAPI struct {
	node
	OpenAPI    string `yaml:"openapi"`
	Info       Info
	Servers    []Server
	Paths      map[string]*PathItem
	Components Components
	metadata   map[string]string
}

func LoadOpenAPI(openAPIFile string) (*OpenAPI, error) {
	// skeleton
	absPath, err := filepath.Abs(openAPIFile)
	if err != nil {
		return nil, err
	}
	api := OpenAPI{
		node: node{
			basePath: filepath.Dir(absPath),
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

	err = yaml.Unmarshal(content, &api)
	if err != nil {
		return nil, fmt.Errorf("yaml unmarshalling error: %w", err)
	}

	api.setName(api.Info.Title)

	// Resolve references
	newApi, err := Traverse(&api, resolveRefs)

	if err != nil {
		return nil, err
	}

	return newApi, err
}

func SetRenderer(api *OpenAPI, renderer Renderer) error {
	_, err := Traverse(api, func(_ string, _, child Traversable) (Traversable, error) {
		child.setRenderer(renderer)
		parent := child.GetParent()
		if parent != nil {
			parent.setRenderer(renderer)
		}

		return child, nil
	})
	return err
}

func (o *OpenAPI) getRenderer() Renderer {
	return o.renderer
}

func (o *OpenAPI) SetMetadata(metadata map[string]string) {
	o.metadata = metadata
}

func (o *OpenAPI) GetMetadata() map[string]string {
	return o.metadata
}

func (o *OpenAPI) getRef() string {
	return ""
}

func (o *OpenAPI) getBasePath() string {
	return o.node.basePath
}

func (o *OpenAPI) GetName() string {
	name := o.getRenderer().SanitiseName(o.name)
	return name
}

func (o *OpenAPI) GetOutputFile() string {
	fileName := o.getRenderer().GetOutputFile(o)
	return fileName
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
	if c, ok := child.(*PathItem); ok {
		o.Paths[i] = c
		return
	}
	panic("(o *OpenAPI) setChild:" + errCastFail)
}

// resolveRefs calls readRef on references with the ref path modified appropriately for it's use
func resolveRefs(key string, parent, node Traversable) (Traversable, error) {
	node.setParent(parent)
	if _, ok := node.(*OpenAPI); !ok {
		node.setName(key) // Don't set the root name as that's already been done by this point
	}
	nodeRef := node.getRef()
	if nodeRef != "" {
		openapiBasePath := node.getBasePath()
		ref := filepath.Base(node.getRef())
		err := readRef(filepath.Join(openapiBasePath, ref), node)
		if err != nil {
			return nil, fmt.Errorf("Unable to read reference:\n%w", err)
		}
	}
	return node, nil
}

// ExternalDocs is a programmatic representation of the External Docs object defined here: https://swagger.io/specification/#external-documentation-object
type ExternalDocs struct {
	Description string `yaml:"description"`
	Url         string `yaml:"url"`
}

func Dig(node Traversable, key ...string) Traversable {
	if len(key) == 0 {
		return node
	}
	return Dig(node.getChildren()[key[0]], key[1:]...)
}
