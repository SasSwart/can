package response_test

import (
	"github.com/sasswart/gin-in-a-can/openapi/media"
	"github.com/sasswart/gin-in-a-can/openapi/response"
	"github.com/sasswart/gin-in-a-can/tree"
	"reflect"
	"testing"
)

func TestOpenAPI_Response_GetAndSetName(t *testing.T) {
	want := "testName"
	r := response.Response{}
	r.SetName(want)
	got := r.GetName()
	if got != want {
		t.Fail()
	}
}
func TestOpenAPI_Response_GetRef(t *testing.T) {
	want := "testRef"
	r := response.Response{Ref: want}
	got := r.GetRef()
	if got != want {
		t.Fail()
	}
}
func TestOpenAPI_Response_GetAndSetChildren(t *testing.T) {
	mtName := "testMediaType"
	mediaTypeString := "application/json"
	want := media.Type{Node: tree.Node{Name: mtName}}
	r := response.Response{}
	r.SetChild(mediaTypeString, &want)
	got := r.GetChildren()[mediaTypeString].(*media.Type)
	if !reflect.DeepEqual(got, &want) {
		t.Fail()
	}
}
