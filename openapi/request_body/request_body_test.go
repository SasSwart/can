package request_body

import (
	openapi2 "github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/openapi/root"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/openapi/test"
	"testing"
)

func TestOpenAPI_LoadsRequestBodyValidation(t *testing.T) {
	apiSpec, err := root.LoadAPISpec(openapi2.testAbsOpenAPI)
	if err != nil {
		t.Fail()
	}

	traversable := test.Dig(apiSpec, openapi2.testEndpoint, openapi2.testMethod, openapi2.testReqBody, openapi2.testMediaType, openapi2.testSchema, "name")
	name := traversable.(*schema.Schema)
	if name.MinLength != 1 {
		t.Errorf("got minLength %v, wanted %v", name.MinLength, 1)
	}
	if name.MaxLength != 64 {
		t.Errorf("got maxLength %v, wanted %v", name.MaxLength, 64)
	}
	if name.Pattern != openapi2.testPattern {
		t.Errorf("got pattern %v, wanted %v", name.Pattern, openapi2.testPattern)
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
