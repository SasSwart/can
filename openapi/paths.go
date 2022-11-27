package openapi

import (
	"path/filepath"
)

var _ Traversable = &PathItem{}

// PathItem is a programmatic representation of the Path Item object defined here: https://swagger.io/specification/#path-item-object
type PathItem struct {
	node
	Summary     string
	Description string
	Get         *Operation
	Post        *Operation
	Patch       *Operation
	Delete      *Operation
	Parameters  []Parameter
	Ref         string `yaml:"$ref"`
}

func (p *PathItem) getChildren() map[string]Traversable {
	return p.Operations()
}

//func (p *PathItem) getRef() string {
//	panic("(p *PathItem) getRef() This should never be called")
//	return p.Ref
//}

func (p *PathItem) GetPath() string {
	name := p.name
	return name
}

func (p *PathItem) getBasePath() string {
	// TODO: Deal with absolute paths for both of these parameters
	// For now both of these params are assumed relative
	basePath := filepath.Join(p.parent.getBasePath(), filepath.Dir(p.Ref))
	return basePath
}

func (p *PathItem) Operations() map[string]Traversable {
	operations := map[string]Traversable{}
	if p.Get != nil {
		operations["get"] = p.Get
	}
	if p.Post != nil {
		operations["post"] = p.Post
	}
	if p.Patch != nil {
		operations["patch"] = p.Patch
	}
	if p.Delete != nil {
		operations["delete"] = p.Delete
	}
	return operations
}

func (p *PathItem) setChild(i string, child Traversable) {
	if operation, ok := child.(*Operation); ok {
		p.Operations()[i] = operation
		return

	}
	panic("(p *PathItem) setChild borked")
}
