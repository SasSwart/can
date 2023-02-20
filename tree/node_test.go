package tree_test

import (
	"github.com/sasswart/gin-in-a-can/tree"
	"testing"
)

func TestOpenAPI_Node_GetAndSetName(t *testing.T) {
	want := "testname"
	n := tree.Node{}
	n.SetName(want)
	got := n.GetName()
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

func TestOpenAPI_Node_GetMetadata(t *testing.T) {

	t.Errorf("TODO")
}

func TestOpenAPI_Node_GetOutputFile(t *testing.T) {
	t.Errorf("TODO")
}
