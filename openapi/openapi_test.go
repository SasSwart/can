package openapi

import "testing"

var _ node = &OpenAPI{}
var _ node = &Paths{}
var _ node = &PathItem{}
var _ node = &Operation{}
var _ node = &Parameter{}
var _ node = &RequestBody{}
var _ node = &Schema{}
var _ node = &MediaType{}

func TestOpenAPILoadsValidation(t *testing.T) {
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

	IDParam := apiSpec.Paths["/endpoint"].Post.Parameters[0]
	if IDParam.Required != true {
		t.Fail()
	}
	if IDParam.Schema.Format != "uuid" {
		t.Fail()
	}
}
