package media_test

import (
	"github.com/sasswart/gin-in-a-can/openapi/media"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/tree"
	"reflect"
	"strings"
	"testing"
)

func TestOpenAPI_MediaType_SetAndGetName(t *testing.T) {
	want := "testName"
	mt := media.Type{}
	mt.SetName(want)
	got := strings.Join(mt.GetName(), "")
	if want != got {
		t.Fail()
	}
}

func TestOpenAPI_MediaType_GetAndSetChildren(t *testing.T) {
	want := schema.Schema{Node: tree.Node{}}
	mt := media.Type{Node: tree.Node{}}
	mt.SetChild("", &want)
	got := mt.GetChildren()[schema.Key].(*schema.Schema)
	if !reflect.DeepEqual(&want, got) {
		t.Fail()
	}
}
