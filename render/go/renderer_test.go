package golang_test

import (
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/openapi/media"
	"github.com/sasswart/gin-in-a-can/openapi/path"
	"github.com/sasswart/gin-in-a-can/openapi/request"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/render"
	golang "github.com/sasswart/gin-in-a-can/render/go"
	"github.com/sasswart/gin-in-a-can/test"
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

func TestGolang_SanitiseName(t *testing.T) {

	specPath := "../../" + test.OpenapiFile
	apiSpec, _ := openapi.LoadAPISpec(specPath)

	tests := []struct {
		name     string
		expected string
		node     tree.NodeTraverser
	}{
		{
			name:     "testopenapi root",
			node:     apiSpec,
			expected: "ValidationFixture",
		},
		{
			name:     "testopenapi path",
			node:     test.Dig(apiSpec, test.Endpoint),
			expected: "ValidationFixtureEndpoint",
		},
		{
			name:     "testopenapi pathitem",
			node:     test.Dig(apiSpec, test.Endpoint, path.Post),
			expected: "ValidationFixtureEndpointPost",
		},
		{
			// TODO check that this is creating param names properly
			name:     "testopenapi parameter",
			node:     test.Dig(apiSpec, test.Endpoint, path.Post, "0"),
			expected: "Parameter",
		},
		{
			name:     "testopenapi requestbody",
			node:     test.Dig(apiSpec, test.Endpoint, path.Post, request.BodyKey),
			expected: "ValidationFixtureEndpointPostRequestBody",
		},
		{
			name:     "testopenapi json mediaitem",
			node:     test.Dig(apiSpec, test.Endpoint, path.Post, request.BodyKey, media.JSONKey),
			expected: "ValidationFixtureEndpointPostRequestBody",
		},
		{
			name:     "testopenapi schema",
			node:     test.Dig(apiSpec, test.Endpoint, path.Post, request.BodyKey, media.JSONKey, schema.Key),
			expected: "ValidationFixtureEndpointPostRequestBodyModel",
		},
		{
			name:     "testopenapi property",
			node:     test.Dig(apiSpec, test.Endpoint, path.Post, request.BodyKey, media.JSONKey, schema.Key, "name"),
			expected: "ValidationFixtureEndpointPostRequestBodyModelName",
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			renderer := &golang.Renderer{}
			want := testCase.expected
			got := renderer.SanitiseName(testCase.node.GetName())
			if want != got {
				t.Errorf("Wanted %s but got %s\n", want, got)
			}
		})
	}
}

func TestGolang_SanitiseType(t *testing.T) {
	arrayType := schema.Schema{
		Type: "array",
		Node: tree.Node{Name: "testname"}}
	arrayType.SetChild(schema.SubSchemaKey, &schema.Schema{
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
			expected: "[]Testname",
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
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			renderer := &golang.Renderer{}
			want := testCase.expected
			got := renderer.SanitiseType(testCase.schema)
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

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got := golang.CreateGoFunctionString(testCase.input)
			if got != testCase.expected {
				t.Fail()
			}
		})
	}
}
