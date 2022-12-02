package openapi

import "testing"

func TestOpenAPI_LoadsRequestBodyValidation(t *testing.T) {
	apiSpec, err := LoadOpenAPI(absOpenAPI)
	if err != nil {
		t.Fail()
	}

	transversable := Dig(apiSpec, testEndpoint, testMethod, testReqBody, testMediaType, testSchema)
	name := transversable.(*Schema)
	if name.MinLength != 1 {
		t.Errorf("got minLength %v, wanted %v", name.MinLength, 1)
	}
	if name.MaxLength != 64 {
		t.Errorf("got maxLength %v, wanted %v", name.MaxLength, 64)
	}
	if name.Pattern != testPattern {
		t.Errorf("got pattern %v, wanted %v", name.Pattern, testPattern)
	}
	if name.Required[0] != "name" {
		t.Errorf("found required field %v at index [0], wanted %v", name.Required[0], "name")
	}
}
