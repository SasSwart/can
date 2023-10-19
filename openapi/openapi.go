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
	OpenAPI    string `yaml:"openapi"`
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

// LoadFromYaml expects a path to an openapi file in yaml format. It will unmarshal the yaml file and resolve the
// references contained within.
func LoadFromYaml(openApiFilepath string) (*OpenAPI, error) {
	// Read yaml file
	// TODO: should this function have to perform file IO?
	content, err := os.ReadFile(openApiFilepath)
	if err != nil {
		return nil, fmt.Errorf("LoadFromYaml:: unable to read file \"%s\": %w", openApiFilepath, err)
	}

	api := newBaseOpenApi()
	if err := api.loadFromBytes(content); err != nil {
		return nil, err
	}
	// Resolve references
	if err := api.resolveReferences(openApiFilepath); err != nil {
		return nil, err
	}
	return &api, err
}

// resolveReferences traverses the entire openapi struct tree and loads te h
func (o *OpenAPI) resolveReferences(absoluteBasePath string) error {
	o.SetBasePath(filepath.Dir(absoluteBasePath))
	// TODO: what if JSON?
	o, err := tree.Traverse(o, tree.ResolveRefs)
	if err != nil {
		return fmt.Errorf("tree traversal error: %w", err)
	}
	return nil
}

// loadFromBytes loads the content into an OpenAPI struct. Extraneous functions that aid in this process should live here.A
func (o *OpenAPI) loadFromBytes(content []byte) error {
	err := yaml.Unmarshal(content, o)
	if err != nil {
		return fmt.Errorf("yaml unmarshalling error: %w", err)
	}
	o.SetName(o.Info.Title)
	return nil
}

func newBaseOpenApi() OpenAPI {
	return OpenAPI{
		Node: tree.Node{},
		Components: components.Components{
			Schemas: nil,
		},
		Info:  info.Info{},
		Paths: map[string]*path.Item{},
	}
}
