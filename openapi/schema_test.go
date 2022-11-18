package openapi

import (
	"path/filepath"
	"reflect"
	"testing"
)

var (
	properties = map[string]*Schema{
		"renderable":        &Schema{},
		"anotherRenderable": &Schema{},
	}
	item = &Schema{}
	p    = &MediaType{ // PARENT
		parent: nil, // RESPONSE USED
		name:   "parentName",
		Schema: &Schema{
			parent:      &MediaType{},
			name:        "parentModel",
			Description: "",
			Type:        "string",
		},
	}
)

func emptySchemaWith(childProperties, childItems, parent bool) *Schema {
	switch {
	case parent && !childProperties && !childItems:
		return &Schema{ // BASE SCHEMA
			renderer:   nil,
			parent:     p, // PARENT POINT
			name:       "mainModel",
			Type:       "string",
			Properties: nil, // CHILD POINT
			Items:      nil, // CHILD POINT
		}
	case !parent && childProperties && !childItems:
		return &Schema{ // BASE SCHEMA
			renderer:   nil,
			parent:     p, // PARENT POINT
			name:       "mainModel",
			Type:       "string",
			Properties: properties, // CHILD POINT
			Items:      nil,        // CHILD POINT
		}
	case !parent && !childProperties && childItems:
		return &Schema{ // BASE SCHEMA
			renderer:   nil,
			parent:     p, // PARENT POINT
			name:       "mainModel",
			Type:       "string",
			Properties: nil,  // CHILD POINT
			Items:      item, // CHILD POINT
		}
	case !parent && childProperties && childItems:

		return &Schema{ // BASE SCHEMA
			renderer:   nil,
			parent:     p, // PARENT POINT
			name:       "mainModel",
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
		t.Fail()
	}
	schemaWithChildren = emptySchemaWith(true, true, false)
	shouldBePropAndItemChildren := schemaWithChildren.getChildren()
	if shouldBePropAndItemChildren == nil {
		t.Error("shouldBePropAndItemChildren is nil")
		t.Fail()
	}
	schemaWithChildren = emptySchemaWith(false, true, false)
	shouldBeItemChildren := schemaWithChildren.getChildren()
	if shouldBeItemChildren == nil {
		t.Error("shouldBeItemChildren is nil")
		t.Fail()
	}
	schemaWithChildren = emptySchemaWith(true, false, false)
	shouldBePropChildren := schemaWithChildren.getChildren()
	if shouldBePropChildren == nil {
		t.Error("shouldBePropChildren is nil")
		t.Fail()
	}
}

func TestSchema_ResolveRefs(t *testing.T) {
	basePath := "../"

	ref := "openapi/fixtures/test_schema.yaml"
	absoluteRef, _ := filepath.Abs(filepath.Join(basePath, ref))

	subRef := "openapi/fixtures/sub_schema.yml"
	absoluteSubRef, _ := filepath.Abs(filepath.Join(basePath, subRef))

	schema := Schema{
		Ref: ref,
	}

	t.Errorf("Fix me")
	t.FailNow() // Panic on next line

	newSchema, _ := Traverse(&schema, resolveRefs)

	if newSchema.(*Schema).Ref != absoluteRef {
		t.Log("Schema reference was not correctly resolved to an absolute path")
		t.Fail()
	}

	if newSchema.(*Schema).Items.Ref != absoluteSubRef {
		t.Log("Recursive Schema reference was not correctly resolved to an absolute path")
		t.Fail()
	}
}
