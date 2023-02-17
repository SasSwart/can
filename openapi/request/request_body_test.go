package request_test

import (
	"github.com/sasswart/gin-in-a-can/openapi/root"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/openapi/test"
	"testing"
)

func TestOpenAPI_LoadsRequestBodyValidation(t *testing.T) {
	apiSpec, err := root.LoadAPISpec(test.AbsOpenAPI)
	if err != nil {
		t.Fail()
	}

	traversable := test.Dig(apiSpec, test.Endpoint, test.Method, test.ReqBody, test.MediaType, test.Schema, "name")
	name := traversable.(*schema.Schema)
	if name.MinLength != 1 {
		t.Errorf("got minLength %v, wanted %v", name.MinLength, 1)
	}
	if name.MaxLength != 64 {
		t.Errorf("got maxLength %v, wanted %v", name.MaxLength, 64)
	}
	if name.Pattern != test.Pattern {
		t.Errorf("got pattern %v, wanted %v", name.Pattern, test.Pattern)
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
