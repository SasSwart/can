package openapi

import (
	"fmt"
	"github.com/sasswart/gin-in-a-can/errors"
	"github.com/sasswart/gin-in-a-can/openapi/components"
	"github.com/sasswart/gin-in-a-can/openapi/info"
	"github.com/sasswart/gin-in-a-can/openapi/path"
	"github.com/sasswart/gin-in-a-can/openapi/servers"
	"github.com/sasswart/gin-in-a-can/tree"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var _ tree.NodeTraverser = &OpenAPI{}

// ExternalDocs is a programmatic representation of the External Docs object defined here: https://swagger.io/specification/#external-documentation-object
type ExternalDocs struct {
	Description string `yaml:"description"`
	Url         string `yaml:"url"`
}

// OpenAPI is a programmatic representation of the OpenApi Document object defined here: https://swagger.io/specification/#openapi-object
type OpenAPI struct {
	tree.Node
	OpenAPI    string `yaml:"openapi"` // TODO it's unclear what this refers to
	Info       info.Info
	Servers    []servers.Server
	Paths      map[string]*path.Item
	Components components.Components
}

func (o *OpenAPI) GetRef() string {
	return ""
}

func (o *OpenAPI) GetChildren() map[string]tree.NodeTraverser {
	traversables := map[string]tree.NodeTraverser{}
	for s := range o.Paths {
		p := o.Paths[s]
		traversables[s] = p
	}
	return traversables
}

func (o *OpenAPI) SetChild(i string, child tree.NodeTraverser) {
	if o.Paths == nil {
		o.Paths = make(map[string]*path.Item, 4)
	}
	if c, ok := child.(*path.Item); ok {
		o.Paths[i] = c
		return
	}
	errors.CastFail("(o *OpenAPI) setChild", "NodeTraverser", "*schema.Schema")
}

// LoadFromYaml expects an absolute path. It will unmarshal the yaml file and resolve the references contained within.
func LoadFromYaml(openApiFilepath string) (*OpenAPI, error) {
	// Read yaml file
	content, err := os.ReadFile(openApiFilepath)
	if err != nil {
		return nil, fmt.Errorf("LoadFromYaml:: unable to read file \"%s\": %w", openApiFilepath, err)
	}

	api := NewBaseOpenApi()
	if err := api.Load(content); err != nil {
		return nil, err
	}
	// Resolve references
	if err := api.ResolveReferences(openApiFilepath); err != nil {
		return nil, err
	}
	return &api, err
}

func (o *OpenAPI) ResolveReferences(absoluteBasePath string) error {
	o.SetBasePath(filepath.Dir(absoluteBasePath))
	o, err := tree.Traverse(o, tree.ResolveRefs)
	if err != nil {
		return fmt.Errorf("tree traversal error: %w", err)
	}
	return nil
}

// Load loads the yaml content into the OpenAPI struct. Extraneous functions that aid in this process should live here.
func (o *OpenAPI) Load(content []byte) error {
	err := yaml.Unmarshal(content, o)
	if err != nil {
		return fmt.Errorf("yaml unmarshalling error: %w", err)
	}
	o.SetName(o.Info.Title)
	return nil
}

func NewBaseOpenApi() OpenAPI {
	return OpenAPI{
		Node: tree.Node{},
		Components: components.Components{
			Schemas: nil,
		},
		Info:  info.Info{},
		Paths: map[string]*path.Item{},
	}
}
