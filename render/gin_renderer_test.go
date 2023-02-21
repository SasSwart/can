package render_test

import (
	"testing"
)

func TestOpenAPI_GinRenderer_SanitiseType(t *testing.T) {
	// TODO uncomment this when sanitise type is properly implemented
	//tests := []struct {
	//	schemaType string
	//	expected   string
	//	items      *schema.Schema
	//}{
	//	{
	//		"boolean",
	//		"bool",
	//		nil,
	//	},
	//	{
	//		schemaType: "array",
	//		expected:   "[]testname",
	//		items: &schema.Schema{
	//			Node: tree.Node{
	//				Name: "testname",
	//			},
	//		},
	//	},
	//	{
	//		"object",
	//		"struct",
	//		nil,
	//	},
	//	{
	//		"somethingelse",
	//		"somethingelse",
	//		nil,
	//	},
	//}
	//for _, test := range tests {
	//	r := render.GinRenderer{}
	//	s := &schema.Schema{}
	//	s.Type = test.schemaType
	//	if test.items != nil {
	//		s.Items = test.items
	//	}
	//	want := test.expected
	//	got := r.SanitiseName(s.Name)
	//	if want != got {
	//		t.Errorf("Wanted %s but got %s\n", want, got)
	//	}
	//}
	t.Errorf("TODO")
}

func TestOpenAPI_GinRenderer_SanitiseName(t *testing.T) {
	t.Errorf("TODO")
}

func TestOpenAPI_GinRenderer_GetOutputFile(t *testing.T) {
	t.Errorf("TODO")
}
