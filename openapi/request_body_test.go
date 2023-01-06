package openapi

import (
	"github.com/sasswart/gin-in-a-can/test"
	"testing"
)

func TestOpenAPI_LoadsRequestBodyValidation(t *testing.T) {
	apiSpec, err := LoadOpenAPI(test.AbsOpenAPI)
	if err != nil {
		t.Fail()
	}

	traversable := Dig(apiSpec, test.Endpoint, test.Method, test.ReqBody, test.MediaType, test.Schema)
	name := traversable.(*Schema)
	if name.MinLength != 1 {
		t.Errorf("got minLength %v, wanted %v", name.MinLength, 1)
	}
	if name.MaxLength != 64 {
		t.Errorf("got maxLength %v, wanted %v", name.MaxLength, 64)
	}
	if name.Pattern != test.Pattern {
		t.Errorf("got pattern %v, wanted %v", name.Pattern, test.Pattern)
	}
	if name.Required[0] != "name" {
		t.Errorf("found required field %v at index [0], wanted %v", name.Required[0], "name")
	}
}
