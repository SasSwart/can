package openapi

import (
	"fmt"
	"path"
)

type Paths map[string]PathItem

func (p *Paths) ResolveRefs(basePath string, components *Components) error {
	for key, pathItem := range *p {
		if pathItem.Ref == "" {
			continue
		}

		var newPathItem PathItem
		err := readRef(path.Join(basePath, pathItem.Ref), &newPathItem)
		if err != nil {
			return fmt.Errorf("Unable to read reference:\n%w", err)
		}
		(*p)[key] = newPathItem

		refBasePath := path.Dir(pathItem.Ref)
		err = newPathItem.ResolveRefs(path.Join(basePath, refBasePath), components)
		if err != nil {
			return err
		}
	}
	return nil
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

func (p *PathItem) ResolveRefs(basePath string, components *Components) error {
	for _, operation := range p.Operations() {
		if operation == nil {
			continue
		}
		err := operation.ResolveRefs(basePath, components)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *PathItem) Operations() map[string]*Operation {
	return map[string]*Operation{
		"delete": p.Delete,
		"get":    p.Get,
		"patch":  p.Patch,
		"post":   p.Post,
	}
}
