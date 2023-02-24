package openapi

import (
	"reflect"
	"testing"
)

func (s *Schema) withProperties() *Schema {
	s.Properties = map[string]*Schema{
		"renderable":        {},
		"anotherRenderable": {},
	}
	return s
}

func (s *Schema) withItems() *Schema {
	s.Items = &Schema{}
	return s
}

func (s *Schema) withParent() *Schema {
	s.parent = &MediaType{
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
	return s
}

func TestOpenAPI_Schema_getChildren(t *testing.T) {
	// Sanity Check
	schema := new(Schema)
	shouldBeEmpty := schema.getChildren()
	s := &Schema{}
	if !reflect.DeepEqual(shouldBeEmpty, s.getChildren()) {
		t.Error("shouldBeEmpty is not empty")
	}
	schemaWithChildren := new(Schema).withProperties().withItems()
	shouldBePropAndItems := schemaWithChildren.getChildren()
	if shouldBePropAndItems == nil {
		t.Error("shouldBePropAndItems is nil")
	}
	schemaWithChildren = new(Schema).withItems()
	shouldBeItemChildren := schemaWithChildren.getChildren()
	if shouldBeItemChildren == nil {
		t.Error("shouldBeItemChildren is nil")
	}
	schemaWithChildren = new(Schema).withProperties()
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
