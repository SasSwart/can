package openapi

import (
	"fmt"
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

type pathItem struct {
	node[refContainer, refContainer]
	Summary     string
	Description string
	Get         *Operation
	Post        *Operation
	Patch       *Operation
	Delete      *Operation
	Parameters  []Parameter
	Ref         string `yaml:"$ref"`
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

func (p pathItem) getChildren() []traversable {
	return []traversable{
		p.Delete,
		p.Get,
		p.Patch,
		p.Post,
	}
}

func (p pathItem) setChild(i int, child traversable) {
	// TODO
}
