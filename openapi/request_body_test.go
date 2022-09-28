package openapi

import "testing"

func TestOpenAPILoadsRequestBodyValidation(t *testing.T) {
	apiSpec, err := LoadOpenAPI("fixtures/validation.yaml")
	if err != nil {
		t.Fail()
	}

	bodySchema := apiSpec.Paths["/endpoint"].Post.RequestBody.Content["application/json"].Schema
	if bodySchema.Properties["name"].MinLength != 1 {
		t.Fail()
	}
	if bodySchema.Properties["name"].MaxLength != 64 {
		t.Fail()
	}
	if bodySchema.Properties["name"].Pattern != "^([a-zA-Z0-9])+([-_ @\\.]([a-zA-Z0-9])+)*$" {
		t.Fail()
	}
	if bodySchema.Required[0] != "name" {
		t.Fail()
	}
}
