package openapi

import "testing"

func TestOpenAPILoadsParameterValidation(t *testing.T) {
	apiSpec, err := LoadOpenAPI("fixtures/validation.yaml")
	if err != nil {
		t.Fail()
	}

	IDParam := apiSpec.Paths["/endpoint"].Post.Parameters[0]
	if IDParam.Required != true {
		t.Fail()
	}
	if IDParam.Schema.Type != "string" {
		t.Fail()
	}
	if IDParam.Schema.Format != "uuid" {
		t.Fail()
	}
}
