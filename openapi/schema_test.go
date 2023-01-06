package openapi

import (
	"github.com/sasswart/gin-in-a-can/test"
	"reflect"
	"testing"
)

var (
	properties = map[string]*Schema{
		"renderable":        {},
		"anotherRenderable": {},
	}
	item = &Schema{}
	p    = &MediaType{ // PARENT
		node: node{},
		name: "parentName",
		Schema: &Schema{
			node: node{
				parent: &MediaType{},
				name:   "parentModel",
			},
			Type: "string",
		},
	}
)

func emptySchemaWith(childProperties, childItems, parent bool) *Schema {
	mainNode := node{
		parent: p, // PARENT POINT
		name:   "mainModel",
	}
	switch {
	case parent && !childProperties && !childItems:
		return &Schema{ // BASE SCHEMA
			node: mainNode,
			Type: "string",
		}
	case !parent && childProperties && !childItems:
		return &Schema{ // BASE SCHEMA
			node:       mainNode,
			Type:       "string",
			Properties: properties, // CHILD POINT
		}
	case !parent && !childProperties && childItems:
		return &Schema{ // BASE SCHEMA
			node:  mainNode,
			Type:  "string",
			Items: item, // CHILD POINT
		}
	case !parent && childProperties && childItems:

		return &Schema{ // BASE SCHEMA
			node:       mainNode,
			Type:       "string",
			Properties: properties, // CHILD POINT
			Items:      item,       // CHILD POINT
		}
	}
	return &Schema{}
}

func TestGetChildren(t *testing.T) {
	// Sanity Check
	schemaWithChildren := emptySchemaWith(false, false, false)
	shouldBeEmpty := schemaWithChildren.getChildren()
	s := &Schema{}
	if !reflect.DeepEqual(shouldBeEmpty, s.getChildren()) {
		t.Error("shouldBeEmpty is not empty")
	}
	schemaWithChildren = emptySchemaWith(true, true, false)
	shouldBePropAndItemChildren := schemaWithChildren.getChildren()
	if shouldBePropAndItemChildren == nil {
		t.Error("shouldBePropAndItemChildren is nil")
	}
	schemaWithChildren = emptySchemaWith(false, true, false)
	shouldBeItemChildren := schemaWithChildren.getChildren()
	if shouldBeItemChildren == nil {
		t.Error("shouldBeItemChildren is nil")
	}
	schemaWithChildren = emptySchemaWith(true, false, false)
	shouldBePropChildren := schemaWithChildren.getChildren()
	if shouldBePropChildren == nil {
		t.Error("shouldBePropChildren is nil")
	}
}

func TestSchema_ResolveRefs(t *testing.T) {
	openapi, err := LoadOpenAPI(test.AbsOpenAPI)
	if err != nil {
		t.Errorf(err.Error())
	}
	if openapi == nil {
		t.Errorf("openapi is nil")
	}
	// TODO
	t.Error("TODO load refs through the use of the composed `node` struct and test against that")
}
