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
		openAPIMeta: openAPIMeta{
			basePath: filepath.Dir(openAPIFile),
			refContainerNode: refContainerNode{
				name: "Stratus",
			},
		},
		Components: Components{
			Schemas: map[string]Schema{},
		},
		Paths: map[string]PathItem{},
	}

	// Read yaml file
	content, err := os.ReadFile(openAPIFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read file \"%s\": %w", openAPIFile, err)
	}
	yaml.Unmarshal(content, &api)

	// Resolve references
	newapi, err := Traverse(&api, resolveRefs)
	if err != nil {
		return nil, err
	}

	return newapi.(*OpenAPI), err
}

func resolveRefs(key string, parent, child Traversable) (Traversable, error) {
	childNode, ok := child.(refContainer)
	if !ok {
		return child, nil
	}

	var err error
	switch child.(type) {
	case *PathItem:
		pathItemChild := child.(*PathItem)
		pathItemChild.node = node[*OpenAPI]{
			parent: parent.(*OpenAPI),
			name:   key,
		}
		pathItemChild.parent = parent.(*OpenAPI)
		ref := childNode.getRef()
		if ref != "" {
			basePath := pathItemChild.getBasePath()
			ref := filepath.Base(pathItemChild.Ref)
			err = readRef(filepath.Join(basePath, ref), pathItemChild)
			if err != nil {
				return nil, fmt.Errorf("Unable to read reference:\n%w", err)
			}
		}
		return pathItemChild, nil
	case *Operation:
		operationChild := child.(*Operation)
		if operationChild == nil {
			return child, nil
		}
		operationChild.parent = parent.(refContainer)
		operationChild.refContainerNode = refContainerNode{
			parent: parent.(refContainer),
			name:   key,
		}
		return operationChild, nil
	case *RequestBody:
		requestBodyChild := child.(*RequestBody)
		requestBodyChild.operationChildNode = operationChildNode{
			parent: parent.(*Operation),
			name:   key,
		}
		return requestBodyChild, nil
	case *Response:
		responseChild := child.(*Response)
		responseChild.parent = parent.(*Operation)
		responseChild.name = key
		return responseChild, nil
	case *Parameter:
		parameterChild := child.(*Parameter)
		parameterChild.operationChildNode = operationChildNode{
			parent: parent.(*Operation),
			name:   key,
		}
		return parameterChild, nil
	case *MediaType:
		mediaTypeChild := child.(*MediaType)
		mediaTypeChild.parent = parent.(refContainer)
		mediaTypeChild.name = key
		return mediaTypeChild, nil
	case *Schema:
		schemaChild := child.(*Schema)
		schemaChild.parent = parent.(refContainer)
		schemaChild.name = key
		ref := childNode.getRef()
		if ref != "" {
			basePath := schemaChild.getBasePath()
			ref := filepath.Base(schemaChild.Ref)
			err = readRef(filepath.Join(basePath, ref), schemaChild)
			if err != nil {
				return nil, fmt.Errorf("Unable to read reference:\n%w", err)
			}
		}
		return schemaChild, nil
	}

	return child, nil
}

type openAPIMeta struct {
	refContainerNode
	basePath string
}

func (m openAPIMeta) getBasePath() string {
	return m.basePath
}

// OpenAPI is a programmatic representation of the OpenApi Document object defined here: https://swagger.io/specification/#openapi-object
type OpenAPI struct {
	openAPIMeta
	OpenAPI string `yaml:"openapi"`
	Info    Info
	//Servers Servers
	Servers    []Server // TODO fix bugs after this modification
	Paths      map[string]PathItem
	Components Components
}

func (o *OpenAPI) getParent() Traversable {
	return o.parent
}

func (o *OpenAPI) getChildren() map[string]Traversable {
	traversables := map[string]Traversable{}
	for s, item := range o.Paths {
		traversables[s] = &item
	}
	return traversables
}

func (o *OpenAPI) setChild(i string, child Traversable) {
	c, _ := child.(*PathItem)
	o.Paths[i] = *c
}

func (o *OpenAPI) GetName() string {
	return o.name
}

// ExternalDocs is a programmatic representation of the External Docs object defined here: https://swagger.io/specification/#external-documentation-object
type ExternalDocs struct {
	Description string `yaml:"description"`
	Url         string `yaml:"url"`
}
