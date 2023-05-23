package tree_test

import (
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/openapi/path"
	"github.com/sasswart/gin-in-a-can/tree"
	"reflect"
	"strings"
	"testing"
)

func TestOpenAPI_Node_GetAndSetName(t *testing.T) {
	want := "testname"
	n := tree.Node{}
	n.SetName(want)
	got := strings.Join(n.GetName(), "")
	if got != want {
		t.Fail()
	}
}

func TestOpenAPI_Node_GetAndSetBasePath(t *testing.T) {
	want := "testbasepath"
	n := tree.Node{}
	n.SetBasePath(want)
	got := n.GetBasePath()
	if got != want {
		t.Fail()
	}
}

func TestOpenAPI_Node_GetRef(t *testing.T) {
	// Should not be implemented for this type
}

func TestOpenAPI_Node_GetAndSetMetadata(t *testing.T) {
	p := path.Item{Node: tree.Node{}}
	o := openapi.OpenAPI{Node: tree.Node{}}
	p.SetParent(&o)
	o.SetChild("/test", &p)
	want := tree.Metadata{"key": "value"}
	p.Node.SetMetadata(want)
	got := o.Node.GetMetadata()
	if !reflect.DeepEqual(want, got) {
		t.Fail()
	}
}
