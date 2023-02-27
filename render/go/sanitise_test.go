package golang

import (
	"github.com/sasswart/gin-in-a-can/config"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/render"
	"github.com/sasswart/gin-in-a-can/tree"
	"testing"
)

func TestGolang_SanitiseType(t *testing.T) {
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
			e := render.Engine{}.New(Renderer{}, newTestConfig())
			s := &schema.Schema{}
			s.Type = test.schemaType
			if test.properties != nil {
				s.SetChild("0", test.properties)
			}
			want := test.expected
			got := e.GetRenderer().SanitiseType(s)
			if want != got {
				t.Errorf("Wanted %s but got %s\n", want, got)
			}
		})
	}
}
func TestGolang_clean(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "cleanFunctionString leading number",
			input:    "2functionname",
			expected: "functionname",
		},
		{
			name:     "cleanFunctionString spaces",
			input:    "function name",
			expected: "functionname",
		},
		{
			name:     "cleanFunctionString odd characters",
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
			got := cleanFunctionString(test.input)
			if got != test.expected {
				t.Fail()
			}
		})
	}
}

func newTestConfig() config.Data {
	config.ConfigPath = "../config/config_test.yml"
	config.Debug = true
	return config.Data{
		Generator: config.Generator{},
		Template: config.Template{
			Name: "go-gin",
		},
		TemplatesDir: "../templates",
		OpenAPIFile:  "../openapi/test/fixtures/validation_no_refs.yaml",
		OutputPath:   ".",
	}
}
