package parameter_test

import (
	"github.com/sasswart/gin-in-a-can/openapi/root"
	"github.com/sasswart/gin-in-a-can/openapi/test"
	"testing"
)

func TestOpenAPILoadsParameterValidation(t *testing.T) {
	apiSpec, _ := root.LoadAPISpec(test.AbsOpenAPI)

	IDParam := apiSpec.Paths["/endpoint"].Post.Parameters[0]
	if IDParam.Required != true {
		t.Errorf("Expected id parameter->required to be boolean value true, not %v", IDParam.Required)
	}
	if IDParam.Schema.Type != "string" {
		t.Errorf("Expected id parameter->schema->type to contain `string`, not %v", IDParam.Schema.Type)
	}
	if IDParam.Schema.Format != "uuid" {
		t.Errorf("Expected id parameter->schema->format to contain `uuid`, not %v", IDParam.Schema.Format)
	}
}

func TestOpenAPI_Parameter_getChildren(t *testing.T) {
	t.Errorf("TODO")
}