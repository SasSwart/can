package schema_test

import (
	"github.com/sasswart/gin-in-a-can/openapi/media"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/tree"
	"reflect"
	"testing"
)

func emptySchemaWith(childProperties, childItems, parent bool) *schema.Schema {
	properties := map[string]*schema.Schema{
		"renderable":        {},
		"anotherRenderable": {},
	}
	item := &schema.Schema{}
	p := &media.Type{ // PARENT
		Node: tree.Node{},
		Schema: &schema.Schema{
			Node: tree.Node{
				Name: "parentModel",
			},
			Type: "string",
		},
	}
	p.Schema.Node.SetParent(&media.Type{})
	mainNode := tree.Node{
		Name: "mainModel",
	}
	mainNode.SetParent(p)
	switch {
	case parent && !childProperties && !childItems:
		return &schema.Schema{ // BASE SCHEMA
			Node: mainNode,
			Type: "string",
		}
	case !parent && childProperties && !childItems:
		return &schema.Schema{ // BASE SCHEMA
			Node:       mainNode,
			Type:       "string",
			Properties: properties, // CHILD POINT
		}
	case !parent && !childProperties && childItems:
		return &schema.Schema{ // BASE SCHEMA
			Node:  mainNode,
			Type:  "string",
			Items: item, // CHILD POINT
		}
	case !parent && childProperties && childItems:

		return &schema.Schema{ // BASE SCHEMA
			Node:       mainNode,
			Type:       "string",
			Properties: properties, // CHILD POINT
			Items:      item,       // CHILD POINT
		}
	}
	return &schema.Schema{}
}

func TestOpenAPI_Schema_getChildren(t *testing.T) {
	// Sanity Check
	schemaWithChildren := emptySchemaWith(false, false, false)
	shouldBeEmpty := schemaWithChildren.GetChildren()
	s := &schema.Schema{}
	if !reflect.DeepEqual(shouldBeEmpty, s.GetChildren()) {
		t.Error("shouldBeEmpty is not empty")
	}
	schemaWithChildren = emptySchemaWith(true, true, false)
	shouldBePropAndItemChildren := schemaWithChildren.GetChildren()
	if shouldBePropAndItemChildren == nil {
		t.Error("shouldBePropAndItemChildren is nil")
	}
	schemaWithChildren = emptySchemaWith(false, true, false)
	shouldBeItemChildren := schemaWithChildren.GetChildren()
	if shouldBeItemChildren == nil {
		t.Error("shouldBeItemChildren is nil")
	}
	schemaWithChildren = emptySchemaWith(true, false, false)
	shouldBePropChildren := schemaWithChildren.GetChildren()
	if shouldBePropChildren == nil {
		t.Error("shouldBePropChildren is nil")
	}
}

func TestOpenAPI_Schema_IsRequired(t *testing.T) {
	nilSchema := &schema.Schema{
		Required: nil,
	}
	schemaWithRequiredName := &schema.Schema{
		Required: []string{"name"},
	}
	if nilSchema.IsRequired("name") {
		t.Fatalf("Name does not exist in nilSchema")
	}
	if !schemaWithRequiredName.IsRequired("name") {
		t.Fatalf("Name does exist in schemaWithRequiredName and was not found")
	}
}
