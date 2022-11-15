package openapi

import (
	"path/filepath"
)

var _ refContainer = PathItem{}
var _ refContainer = &PathItem{}
var _ Traversable = PathItem{}

// PathItem is a programmatic representation of the Path Item object defined here: https://swagger.io/specification/#path-item-object
type PathItem struct {
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

func (p PathItem) Operations() map[string]Traversable {
	return map[string]Traversable{
		"delete": p.Delete,
		"get":    p.Get,
		"patch":  p.Patch,
		"post":   p.Post,
	}
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

func (p PathItem) setChild(i string, child Traversable) {
	// TODO
}
