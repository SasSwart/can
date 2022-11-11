package openapi

import (
	"fmt"
	"path"
)

type Paths map[string]PathItem

func (p *Paths) ResolveRefs(basePath string) error {
	for pathName, pathItem := range *p {
		ref := pathItem.Ref
		if ref != "" {
			var newPathItem PathItem
			err := readRef(path.Join(basePath, pathItem.Ref), &newPathItem)
			if err != nil {
				return fmt.Errorf("Unable to read reference:\n%w", err)
			}
			newPathItem.Ref = ref
			pathItem = newPathItem
		}

		refBasePath := path.Dir(pathItem.Ref)
		err := pathItem.ResolveRefs(path.Join(basePath, refBasePath))
		if err != nil {
			return err
		}

		(*p)[pathName] = pathItem
	}
	return nil
}

func (p *Paths) Render() error {
	fmt.Println("Rendering API Path")
	for _, pathItem := range *p {
		err := pathItem.Render()
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Paths) GetSchemas(name string) (schemas map[string]Schema) {
	schemas = map[string]Schema{}
	for _, pathItem := range *p {
		for name, schema := range pathItem.GetSchemas(name) {
			schemas[name] = schema
		}
	}

	return schemas
}

type PathItem struct {
	Summary     string
	Description string
	Get         *Operation
	Post        *Operation
	Patch       *Operation
	Delete      *Operation
	Parameters  []Parameter
	Ref         string `yaml:"$ref"`
}

func (p *PathItem) ResolveRefs(basePath string) error {
	for _, operation := range p.Operations() {
		if operation == nil {
			continue
		}
		err := operation.ResolveRefs(basePath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *PathItem) Render() error {
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

func (p *PathItem) GetSchemas(name string) (schemas map[string]Schema) {
	schemas = map[string]Schema{}
	for _, operation := range p.Operations() {
		for name, schema := range operation.GetSchemas(name) {
			schemas[name] = schema
		}
	}

	return schemas
}

func (p *PathItem) Operations() map[string]*Operation {
	return map[string]*Operation{
		"delete": p.Delete,
		"get":    p.Get,
		"patch":  p.Patch,
		"post":   p.Post,
	}
}
