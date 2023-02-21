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
	if c, ok := child.(*path.Item); ok {
		o.Paths[i] = c
		return
	}
	errors.CastFail("(o *OpenAPI) setChild", "NodeTraverser", "*schema.Schema")
}

func LoadAPISpec(openAPIFile string) (*OpenAPI, error) {
	// skeleton
	absPath, err := filepath.Abs(openAPIFile)
	if err != nil {
		return nil, err
	}
	api := OpenAPI{
		Node: tree.Node{},
		Components: components.Components{
			Schemas: nil,
		},
		Info:  info.Info{},
		Paths: map[string]*path.Item{},
	}
	api.SetBasePath(filepath.Dir(absPath))

	// Read yaml file
	content, err := os.ReadFile(openAPIFile)
	if err != nil {
		return nil, fmt.Errorf("LoadAPISpec:: unable to read file \"%s\": %w", openAPIFile, err)
	}

	err = yaml.Unmarshal(content, &api)
	if err != nil {
		return nil, fmt.Errorf("LoadAPISpec:: yaml unmarshalling error: %w", err)
	}

	api.SetName(api.Info.Title)

	// Resolve references
	newApi, err := tree.Traverse(&api, tree.ResolveRefs)

	if err != nil {
		return nil, fmt.Errorf("LoadAPISpec:: tree traversal error: %w", err)
	}

	return newApi, err
}

type Config struct {
	ModuleName        string
	BasePackageName   string
	TemplateDirectory string
	TemplateName      string
	// Used to live in renderer config
	// TODO structure this in a more accessible way
	OpenAPIFile      string
	OutputPath       string
	WorkingDirectory string
	ConfigFilePath   string
}
