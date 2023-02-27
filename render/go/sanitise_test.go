package golang_test

import (
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/render"
	golang "github.com/sasswart/gin-in-a-can/render/go"
	"github.com/sasswart/gin-in-a-can/tree"
	"testing"
)

func TestGolang_SetTemplateFuncMap(t *testing.T) {
	g := &golang.Renderer{Base: &render.Base{}}
	g.SetTemplateFuncMap(nil)
	if g.Base.TemplateFuncMapping == nil {
		t.Errorf("TemplateFuncMapping error")
	}
	if g.GetTemplateFuncMap() == nil {
		t.Errorf("GetTemplateFuncMap() error")
	}
}

func TestGolang_SanitiseType(t *testing.T) {
	arrayType := schema.Schema{
		Type: "array",
		Node: tree.Node{Name: "testname"}}
	arrayType.SetChild("0", &schema.Schema{
		Node: tree.Node{
			Name: "testname",
		},
		Properties: map[string]*schema.Schema{
			"0": {},
		},
	})

	tests := []struct {
		name     string
		expected string
		schema   *schema.Schema
	}{
		{
			name:     "boolean conversion",
			schema:   &schema.Schema{Type: "boolean"},
			expected: "bool",
		},
		{
			name:     "array conversion",
			schema:   &arrayType,
			expected: "[]testname",
		},
		{
			name:     "object conversion",
			schema:   &schema.Schema{Type: "object"},
			expected: "struct",
		},
		{
			name:     "integer",
			schema:   &schema.Schema{Type: "integer"},
			expected: "int",
		},
		{
			name:     "other conversion",
			schema:   &schema.Schema{Type: "list"},
			expected: "list",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			renderer := &golang.Renderer{}
			want := test.expected
			got := renderer.SanitiseType(test.schema)
			if want != got {
				t.Errorf("Wanted %s but got %s\n", want, got)
			}
		})
	}
}
func TestGolang_CleanFunctionString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "leading number",
			input:    "2functionname",
			expected: "functionname",
		},
		{
			name:     "spaces",
			input:    "function name",
			expected: "functionname",
		},
		{
			name:     "odd characters",
			input:    "function &^%$#@!|'\"name",
			expected: "functionname",
		},
		{
			name:     "check numbers",
			input:    "a1234567890",
			expected: "a1234567890",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := golang.CleanFunctionString(test.input)
			if got != test.expected {
				t.Fail()
			}
		})
	}
}
