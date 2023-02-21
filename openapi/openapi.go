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

var DEBUG bool
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
		return nil, fmt.Errorf("unable to read file \"%s\": %w", openAPIFile, err)
	}

	err = yaml.Unmarshal(content, &api)
	if err != nil {
		return nil, fmt.Errorf("yaml unmarshalling error: %w", err)
	}

	api.SetName(api.Info.Title)

	// Resolve references
	newApi, err := tree.Traverse(&api, ResolveRefs)

	if err != nil {
		return nil, err
	}

	return newApi, err
}

// ResolveRefs calls readRef on references with the ref path modified appropriately for it's use
func ResolveRefs(key string, parent, node tree.NodeTraverser) (tree.NodeTraverser, error) {
	node.SetParent(parent)
	if _, ok := node.(*OpenAPI); !ok {
		node.SetName(key) // Don't set the root name as that's already been done by this point
	}
	nodeRef := node.GetRef()
	if nodeRef != "" {
		openapiBasePath := node.GetBasePath()
		ref := filepath.Base(node.GetRef())
		err := readRef(filepath.Join(openapiBasePath, ref), node)
		if err != nil {
			return nil, fmt.Errorf("Unable to read reference:\n%w", err)
		}
	}
	return node, nil
}

// ReadRef takes a reference and attempts to unmarshal it's content into the struct being passed as `i`.
// As it happens, this ref is contained within the struct that is being unmarshalled into.
func readRef(absFilename string, n tree.NodeTraverser) error {
	if DEBUG { // this can be a particularly noisy Printf call
		fmt.Printf("Reading reference: %s\n", absFilename)
	}
	content, err := os.ReadFile(absFilename)
	if err != nil {
		return fmt.Errorf("unable to resolve Reference: %w", err)
	}

	err = yaml.Unmarshal(content, n)
	if err != nil {
		return fmt.Errorf("unable to unmarshal reference file:\n%w", err)
	}

	return nil
}
