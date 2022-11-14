package openapi

import (
	"fmt"
	"path"
	"path/filepath"
)

//type paths map[string]pathItem

//func (p *paths) ResolveRefs() error {
//	for pathName, pathItem := range *p {
//		ref := pathItem.Ref
//		if ref != "" {
//			var newPathItem pathItem
//			err := readRef(path.Join(, pathItem.Ref), &newPathItem)
//			if err != nil {
//				return fmt.Errorf("Unable to read reference:\n%w", err)
//			}
//			newPathItem.Ref = ref
//			pathItem = newPathItem
//		}
//
//		err := pathItem.ResolveRefs()
//		if err != nil {
//			return err
//		}
//
//		(*p)[pathName] = pathItem
//	}
//	return nil
//}

// Paths is the collection of Path objects contained within the OpenAPI object
type Paths map[string]Path

func (p *Paths) ResolveRefs(basePath string) error {
	for _, pathItem := range *p {
		if pathItem.Ref != "" {
			var newPathItem Path
			err := readRef(path.Join(basePath, pathItem.Ref), &newPathItem)
			if err != nil {
				return fmt.Errorf("Unable to read reference:\n%w", err)
			}
			pathItem = newPathItem
		}
	}
	return nil
}

//
//func (p *paths) Render() error {
//	fmt.Println("Rendering API Path")
//	for _, pathItem := range *p {
//		err := pathItem.Render()
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func (p *paths) GetSchemas(name string) (schemas map[string]Schema) {
//	schemas = map[string]Schema{}
//	for _, pathItem := range *p {
//		for name, schema := range pathItem.GetSchemas(name) {
//			schemas[name] = schema
//		}
//	}
//
//	return schemas
//}

func (p *Path) ResolveRefs(basePath string) error {
	if p.Ref != "" {
		for _, operation := range p.Operations() {
			if operation == nil {
				continue
			}
			err := operation.ResolveRefs(basePath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// pathItem is a programmatic representation of the Path Item object defined here: https://swagger.io/specification/#path-item-object
type pathItem struct {
	parent      *OpenAPI
	Summary     string
	Description string
	Get         *Operation
	Post        *Operation
	Patch       *Operation
	Delete      *Operation
	Parameters  []Parameter
	Ref         string `yaml:"$ref"`
}

var _ traversable = pathItem{}

func (p pathItem) getParent() traversable {
	return p.parent
}

func (p pathItem) getBasePath() string {
	// TODO: Deal with absolute paths for both of these parameters
	// For now both of these params are assumed relative
	return filepath.Join(p.parent.getBasePath(), p.Ref)
}

func (p pathItem) ResolveRefs() error {

	return nil
}

func (p pathItem) Render() error {
	fmt.Println("Rendering API Path Item")
	for _, operation := range p.Operations() {
		if operation == nil {
			continue
		}
		err := operation.Render()
		if err != nil {
			return err
		}
	}

	return nil
}

func (p pathItem) GetSchemas(name string) (schemas map[string]Schema) {
	schemas = map[string]Schema{}
	for _, operation := range p.Operations() {
		for name, schema := range operation.GetSchemas(name) {
			schemas[name] = schema
		}
	}

	return schemas
}

func (p pathItem) Operations() map[string]*Operation {
	return map[string]*Operation{
		"delete": p.Delete,
		"get":    p.Get,
		"patch":  p.Patch,
		"post":   p.Post,
	}
}

func (p pathItem) getChildren() childContainer[int] {
	return childContainerMap[string]{
		p.Operations(),
	}
}

func (p pathItem) setChild(i int, child traversable) {
	// TODO
}
