package request_test

import (
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/openapi/media"
	"github.com/sasswart/gin-in-a-can/openapi/path"
	"github.com/sasswart/gin-in-a-can/openapi/request"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/test"
	"github.com/sasswart/gin-in-a-can/tree"
	"testing"
)

func TestOpenAPI_LoadsRequestBodyValidation(t *testing.T) {
	apiSpec, _ := openapi.LoadAPISpec("../../" + test.OpenapiFile)

	traversable := test.Dig(apiSpec, test.Endpoint, path.Post, request.BodyKey, media.JSONKey, schema.Key, "name")
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

func TestOpenAPI_RequestBody_GetAndSetChild(t *testing.T) {
	rb := &request.Body{}
	want := &media.Type{Node: tree.Node{Name: "MediaTypeTest"}}
	rb.SetChild("1", want)
	got := rb.GetChildren()["1"]
	if got != want {
		t.Fail()
	}
}
func TestOpenAPI_RequestBody_SetAndGetName(t *testing.T) {
	rb := &request.Body{Node: tree.Node{
		Name: "this should be replaced by SetName()",
	}}
	want := "testName"
	rb.SetName(want)
	got := rb.GetName()
	if got != want {
		t.Fail()
	}
}
