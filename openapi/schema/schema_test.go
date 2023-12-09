package schema_test

import (
	"github.com/sasswart/gin-in-a-can/openapi/parameter"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/tree"
	"path/filepath"
	"testing"
)

func TestOpenAPI_Schema_SetAndGetArrayChildren(t *testing.T) {
	s := new(schema.Schema)
	s.Items = new(schema.Schema)
	s.Properties = make(map[string]*schema.Schema, 4)
	s.Name = "Test Schema"
	s.Type = "array"
	this := &schema.Schema{
		Node: tree.Node{
			Name: "schema item",
		},
	}
	that := &schema.Schema{
		Node: tree.Node{
			Name: "that",
		},
	}
	theOther := &schema.Schema{
		Node: tree.Node{
			Name: "theOther",
		},
	}
	s.SetChild(schema.ItemsKey, this)
	s.SetChild("that", that)
	s.SetChild("theOther", theOther)
	children := s.GetChildren()

	// Check Item
	if children[schema.ItemsKey].(*schema.Schema) != this {
		t.Fail()
	}
	// We should not get any properties for an array
	if children["that"] != nil {
		t.Fail()
	}
	if children["theOther"] != nil {
		t.Fail()
	}
}

func TestOpenAPI_Schema_SetAndGetObjectChildren(t *testing.T) {
	s := new(schema.Schema)
	s.Items = new(schema.Schema)
	s.Properties = make(map[string]*schema.Schema, 4)
	s.Name = "Test Schema"
	s.Type = "object"
	this := &schema.Schema{
		Node: tree.Node{
			Name: "schema item",
		},
	}
	that := &schema.Schema{
		Node: tree.Node{
			Name: "that",
		},
	}
	theOther := &schema.Schema{
		Node: tree.Node{
			Name: "theOther",
		},
	}
	s.SetChild(schema.ItemsKey, this)
	s.SetChild("that", that)
	s.SetChild("theOther", theOther)
	children := s.GetChildren()

	// Check Item
	if children[schema.ItemsKey] != nil {
		t.Fail()
	}
	if children["that"].(*schema.Schema) != that {
		t.Fail()
	}
	if children["theOther"].(*schema.Schema) != theOther {
		t.Fail()
	}
}
func TestOpenAPI_Schema_GetAndSetBasePath(t *testing.T) {
	ref := "testRef/ref"
	basePath := "/base/path/"
	dir, _ := filepath.Split(ref)
	want := filepath.Join(basePath, dir)
	s := &schema.Schema{Ref: ref, Node: tree.Node{}}
	p := parameter.Parameter{Schema: s}
	s.SetParent(&p)

	p.GetChildren()[schema.PropertyKey].(*schema.Schema).SetBasePath(basePath)
	got := p.GetChildren()[schema.PropertyKey].(*schema.Schema).GetBasePath()
	if want != got {
		t.Fail()
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
