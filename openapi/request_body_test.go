package openapi

import "testing"

func TestOpenAPI_LoadsRequestBodyValidation(t *testing.T) {
	apiSpec, err := LoadOpenAPI(testAbsOpenAPI)
	if err != nil {
		t.Fail()
	}

	traversable := Dig(apiSpec, testEndpoint, testMethod, testReqBody, testMediaType, testSchema, "name")
	name := traversable.(*Schema)
	if name.MinLength != 1 {
		t.Errorf("got minLength %v, wanted %v", name.MinLength, 1)
	}
	if name.MaxLength != 64 {
		t.Errorf("got maxLength %v, wanted %v", name.MaxLength, 64)
	}
	if name.Pattern != testPattern {
		t.Errorf("got pattern %v, wanted %v", name.Pattern, testPattern)
	}
	if name.IsRequired("name") {
		t.Errorf("found required field %v at index [0], wanted %v", name.IsRequired("name"), "name")
	}
}

func TestOpenAPI_RequestBody_setChild(t *testing.T) {
	t.Errorf("TODO")
}
func TestOpenAPI_RequestBody_getChildren(t *testing.T) {
	t.Errorf("TODO")
}
func TestOpenAPI_RequestBody_getName(t *testing.T) {
	t.Errorf("TODO")
}
