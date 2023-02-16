package openapi

import (
	"reflect"
	"testing"
)

func emptySchemaWith(childProperties, childItems, parent bool) *Schema {
	properties := map[string]*Schema{
		"renderable":        {},
		"anotherRenderable": {},
	}
	item := &Schema{}
	p := &MediaType{ // PARENT
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

func TestOpenAPI_Schema_getChildren(t *testing.T) {
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

func TestOpenAPI_Schema_IsRequired(t *testing.T) {
	nilSchema := &Schema{
		Required: nil,
	}
	schemaWithRequiredName := &Schema{
		Required: []string{"name"},
	}
	if nilSchema.IsRequired("name") {
		t.Fatalf("Name does not exist in nilSchema")
	}
	if !schemaWithRequiredName.IsRequired("name") {
		t.Fatalf("Name does exist in schemaWithRequiredName and was not found")
	}
}
