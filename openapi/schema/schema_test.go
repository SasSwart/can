package schema_test

import (
	"github.com/sasswart/gin-in-a-can/openapi/parameter"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/tree"
	"strings"
	"testing"
)

func TestOpenAPI_Schema_SetAndGetChildren(t *testing.T) {
	s := new(schema.Schema)
	s.Items = new(schema.Schema)
	s.Properties = make(map[string]*schema.Schema, 4)
	s.Name = "Test Schema"
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
	s.SetChild("item", this)
	s.SetChild("that", that)
	s.SetChild("theOther", theOther)
	children := s.GetChildren()

	// Check Item
	if children["item"].(*schema.Schema) != this {
		t.Fail()
	}
	// Check Properties
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
	want := basePath + strings.Split(ref, "/")[0]
	s := &schema.Schema{Ref: ref, Node: tree.Node{}}
	p := parameter.Parameter{Schema: s}
	s.SetParent(&p)

	p.GetChildren()["Model"].(*schema.Schema).SetBasePath(basePath)
	got := p.GetChildren()["Model"].(*schema.Schema).GetBasePath()
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
