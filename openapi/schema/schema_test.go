package schema_test

import (
	"github.com/sasswart/gin-in-a-can/openapi/parameter"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/tree"
	"path/filepath"
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
	dir, _ := filepath.Split(ref)
	want := filepath.Join(basePath, dir)
	s := &schema.Schema{Ref: ref, Node: tree.Node{}}
	p := parameter.Parameter{Schema: s}
	s.SetParent(&p)

	p.GetChildren()[schema.Key].(*schema.Schema).SetBasePath(basePath)
	got := p.GetChildren()[schema.Key].(*schema.Schema).GetBasePath()
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
func TestOpenAPI_Schema_GetType(t *testing.T) {
	tests := []struct {
		name       string
		schemaType string
		expected   string
		properties *schema.Schema
	}{
		{
			name:       "boolean conversion",
			schemaType: "boolean",
			expected:   "bool",
		},
		{
			name:       "array conversion",
			schemaType: "array",
			expected:   "[]testname",
			properties: &schema.Schema{
				Node: tree.Node{
					Name: "testname",
				},
				Properties: map[string]*schema.Schema{
					"0": {},
				},
			},
		},
		{
			name:       "object conversion",
			schemaType: "object",
			expected:   "struct",
		},
		{
			name:       "other conversion",
			schemaType: "somethingelse",
			expected:   "somethingelse",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := &schema.Schema{}
			s.Type = test.schemaType
			if test.properties != nil {
				s.SetChild("0", test.properties)
			}
			want := test.expected
			got := s.GetType()
			if want != got {
				t.Errorf("Wanted %s but got %s\n", want, got)
			}
		})
	}
}
