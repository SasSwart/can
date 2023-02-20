package render_test

import (
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/render"
	"github.com/sasswart/gin-in-a-can/tree"
	"testing"
)

func TestOpenAPI_GinRenderer_sanitiseType(t *testing.T) {
	tests := []struct {
		schemaType string
		expected   string
		items      *schema.Schema
	}{
		{
			"boolean",
			"bool",
			nil,
		},
		{
			schemaType: "array",
			expected:   "[]testname",
			items: &schema.Schema{
				Node: tree.Node{
					Name: "testname",
				},
			},
		},
		{
			"object",
			"struct",
			nil,
		},
		{
			"somethingelse",
			"somethingelse",
			nil,
		},
	}
	for _, test := range tests {
		r := render.GinRenderer{}
		s := &schema.Schema{}
		s.Type = test.schemaType
		if test.items != nil {
			s.Items = test.items
		}
		want := test.expected
		got := r.SanitiseName(s.Name)
		if want != got {
			t.Errorf("Wanted %s but got %s\n", want, got)
		}
	}
}

func TestOpenAPI_GinRenderer_sanitiseName(t *testing.T) {
	t.Errorf("TODO")
}

func TestOpenAPI_GinRenderer_getOutputFile(t *testing.T) {
	t.Errorf("TODO")
}
