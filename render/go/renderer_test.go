package golang_test

import (
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/openapi/media"
	"github.com/sasswart/gin-in-a-can/openapi/request"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/render"
	golang "github.com/sasswart/gin-in-a-can/render/go"
	"github.com/sasswart/gin-in-a-can/test"
	"github.com/sasswart/gin-in-a-can/tree"
	"net/http"
	"testing"
)

func TestGolang_SetTemplateFuncMap(t *testing.T) {
	g := &golang.Renderer{Base: &render.Base{}}
	g.SetTemplateFuncMap(nil)
	if g.GetTemplateFuncMap() == nil {
		t.Errorf("GetTemplateFuncMap() error")
	}
}

func TestGolang_SanitiseName(t *testing.T) {

	specPath := "../../" + test.OpenapiFile
	apiSpec, _ := openapi.LoadFromYaml(specPath)
	goPropertiesWithDashes := "go-properties-with-dashes"

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
			node:     test.Dig(apiSpec, test.Endpoint, http.MethodPost),
			expected: "ValidationFixtureEndpointPost",
		},
		{
			name:     "testopenapi parameter",
			node:     test.Dig(apiSpec, test.Endpoint, http.MethodPost, "0"),
			expected: "Parameter",
		},
		{
			name:     "testopenapi requestbody",
			node:     test.Dig(apiSpec, test.Endpoint, http.MethodPost, request.BodyKey),
			expected: "ValidationFixtureEndpointPostRequestBody",
		},
		{
			name:     "testopenapi json mediaitem",
			node:     test.Dig(apiSpec, test.Endpoint, http.MethodPost, request.BodyKey, media.JSONKey),
			expected: "ValidationFixtureEndpointPostRequestBody",
		},
		{
			name:     "testopenapi schema",
			node:     test.Dig(apiSpec, test.Endpoint, http.MethodPost, request.BodyKey, media.JSONKey, schema.PropertyKey),
			expected: "ValidationFixtureEndpointPostRequestBodyModel",
		},
		{
			name:     "testopenapi property",
			node:     test.Dig(apiSpec, test.Endpoint, http.MethodPost, request.BodyKey, media.JSONKey, schema.PropertyKey, "name"),
			expected: "ValidationFixtureEndpointPostRequestBodyModelName",
		},
		{
			name:     "testopenapi property with dashes in name",
			node:     test.Dig(apiSpec, test.Endpoint, http.MethodPost, request.BodyKey, media.JSONKey, schema.PropertyKey, goPropertiesWithDashes),
			expected: "ValidationFixtureEndpointPostRequestBodyModelGoPropertiesWithDashes",
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
	arrayType.SetChild(schema.ItemsKey, &schema.Schema{
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
			got := golang.CreateFunctionString(testCase.input)
			if got != testCase.expected {
				t.Fail()
			}
		})
	}
}

func TestGolang_ToTitle(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "handle snake case",
			input:    "domain_id",
			expected: "DomainId",
		},
		{
			name:     "handle adjacent symbols",
			input:    "this/:domain_id",
			expected: "ThisDomainId",
		},
		{
			name:     "handle chaotic naming",
			input:    "#$this/:domain^&_id*(Is INCREDIBLY()chaotic",
			expected: "ThisDomainIdIsIncrediblyChaotic",
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			want := testCase.expected
			got := golang.ToTitle(testCase.input)
			if want != got {
				t.Errorf("Wanted %s but got %s\n", want, got)
			}
		})
	}
}
