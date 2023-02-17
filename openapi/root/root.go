package root

import (
	"fmt"
	"github.com/sasswart/gin-in-a-can/errors"
	"github.com/sasswart/gin-in-a-can/openapi/components"
	"github.com/sasswart/gin-in-a-can/openapi/info"
	"github.com/sasswart/gin-in-a-can/openapi/path"
	"github.com/sasswart/gin-in-a-can/openapi/servers"
	"github.com/sasswart/gin-in-a-can/openapi/test"

	"github.com/sasswart/gin-in-a-can/tree"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

// ExternalDocs is a programmatic representation of the External Docs object defined here: https://swagger.io/specification/#external-documentation-object
type ExternalDocs struct {
	Description string `yaml:"description"`
	Url         string `yaml:"url"`
}

// Root is a programmatic representation of the OpenApi Document object defined here: https://swagger.io/specification/#openapi-object
type Root struct {
	tree.Node
	OpenAPI    string `yaml:"openapi"`
	Info       info.Info
	Servers    []servers.Server
	Paths      map[string]*path.Item
	Components components.Components
	metadata   map[string]string
}

//func SetRenderer(api *Root, renderer render.Renderer) error {
//	// TODO move this to render package, asserting type where needed, to avoid import cycles
//	_, err := tree.Traverse(api, func(_ string, _, child tree.NodeTraverser) (tree.NodeTraverser, error) {
//		child.SetRenderer(renderer)
//		parent := child.GetParent()
//		if parent != nil {
//			parent.SetRenderer(renderer)
//		}
//
//		return child, nil
//	})
//	return err
//}

func (o *Root) GetRef() string {
	return ""
}

func (o *Root) getChildren() map[string]tree.NodeTraverser {
	traversables := map[string]tree.NodeTraverser{}
	for s := range o.Paths {
		path := o.Paths[s]
		traversables[s] = path
	}
	return traversables
}

func (o *Root) setChild(i string, child tree.NodeTraverser) {
	if c, ok := child.(*path.Item); ok {
		o.Paths[i] = c
		return
	}
	errors.CastFail("(o *Root) setChild", "NodeTraverser", "*schema.Schema")
}

func LoadAPISpec(openAPIFile string) (*Root, error) {
	// skeleton
	absPath, err := filepath.Abs(openAPIFile)
	if err != nil {
		return nil, err
	}
	api := Root{
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

	api.Name = api.Info.Title

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
	if _, ok := node.(*Root); !ok {
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
	if test.Debug { // this can be a particularly noisy Printf call
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
