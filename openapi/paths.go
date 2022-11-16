package openapi

import (
	"path/filepath"
)

var _ refContainer = PathItem{}

// PathItem is a programmatic representation of the Path Item object defined here: https://swagger.io/specification/#path-item-object
type PathItem struct {
	renderer Renderer
	node[*OpenAPI]
	Summary     string
	Description string
	Get         *Operation
	Post        *Operation
	Patch       *Operation
	Delete      *Operation
	Parameters  []Parameter
	Ref         string `yaml:"$ref"`
}

func (p *PathItem) SetRenderer(r Renderer) {
	p.renderer = r
}

func (p PathItem) GetName() string {
	return p.renderer.sanitiseName(p.parent.GetName() + p.name)
}

func (p PathItem) getChildren() map[string]Traversable {
	return p.Operations()
}

func (p PathItem) getRef() string {
	return p.Ref
}

func (p PathItem) getParent() Traversable {
	return p.parent
}

func (p PathItem) getBasePath() string {
	// TODO: Deal with absolute paths for both of these parameters
	// For now both of these params are assumed relative
	basePath := filepath.Join(p.parent.getBasePath(), filepath.Dir(p.Ref))
	return basePath
}

func (p PathItem) Operations() map[string]Traversable {
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

func (p PathItem) setChild(i string, child Traversable) {
	// TODO: handle this error
	operation, _ := child.(*Operation)
	p.Operations()[i] = operation
}
