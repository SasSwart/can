package root

import (
	"fmt"
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/openapi/errors"
	"github.com/sasswart/gin-in-a-can/openapi/path"
	"github.com/sasswart/gin-in-a-can/openapi/refs"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	// TODO root package should not have to know about render
	"github.com/sasswart/gin-in-a-can/render"
	"github.com/sasswart/gin-in-a-can/tree"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Config struct {
	OpenAPIFile string
}

// Root is a programmatic representation of the OpenApi Document object defined here: https://swagger.io/specification/#openapi-object
type Root struct {
	tree.Node
	OpenAPI    string `yaml:"openapi"`
	Info       openapi.Info
	Servers    []openapi.Server
	Paths      map[string]*path.Item
	Components openapi.Components
	metadata   map[string]string
}

func LoadAPISpec(openAPIFile string) (*Root, error) {
	// skeleton
	absPath, err := filepath.Abs(openAPIFile)
	if err != nil {
		return nil, err
	}
	api := Root{
		Node: tree.Node{
			basePath: filepath.Dir(absPath),
		},
		Components: openapi.Components{
			Schemas: map[string]schema.Schema{},
		},
		Paths: map[string]*path.Item{},
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

	api.SetName(api.Info.Title)

	// Resolve references
	newApi, err := tree.Traverse(&api, resolveRefs)

	if err != nil {
		return nil, err
	}

	return newApi, err
}

func SetRenderer(api *Root, renderer render.Renderer) error {
	// TODO move this to render package, asserting type where needed, to avoid import cycles
	_, err := tree.Traverse(api, func(_ string, _, child tree.NodeTraverser) (tree.NodeTraverser, error) {
		child.SetRenderer(renderer)
		parent := child.GetParent()
		if parent != nil {
			parent.SetRenderer(renderer)
		}

		return child, nil
	})
	return err
}

func (o *Root) GetRef() string {
	return ""
}

func (o *Root) GetName() string {
	name := o.GetRenderer().SanitiseName(o.name)
	return name
}

func (o *Root) GetOutputFile() string {
	// TODO passing in yourself seems like a smell
	// TODO this override could be removed and handed by the node{} composable
	fileName := o.GetRenderer().GetOutputFile(o)
	return fileName
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
	panic("(o *Root) setChild:" + errors.ErrCastFail)
}

// resolveRefs calls readRef on references with the ref path modified appropriately for it's use
func resolveRefs(key string, parent, node tree.NodeTraverser) (tree.NodeTraverser, error) {
	node.SetParent(parent)
	if _, ok := node.(*Root); !ok {
		node.SetName(key) // Don't set the root name as that's already been done by this point
	}
	nodeRef := node.GetRef()
	if nodeRef != "" {
		openapiBasePath := node.GetBasePath()
		ref := filepath.Base(node.GetRef())
		err := refs.ReadRef(filepath.Join(openapiBasePath, ref), node)
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
