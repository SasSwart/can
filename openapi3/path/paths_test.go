package path_test

import (
	"github.com/sasswart/gin-in-a-can/openapi3"
	"github.com/sasswart/gin-in-a-can/openapi3/operation"
	"github.com/sasswart/gin-in-a-can/openapi3/path"
	"github.com/sasswart/gin-in-a-can/tree"
	"net/http"
	"path/filepath"
	"strings"
	"testing"
)

func TestOpenAPI_PathItem_GetAndSetName(t *testing.T) {
	want := "testName"
	i := path.Item{Node: tree.Node{Name: "this should be overwritten by a call to SetName()"}}
	i.SetName(want)
	got := strings.Join(i.GetName(), "")
	if got != want {
		t.Fail()
	}
}

func Test_Path_GetBasePath(t *testing.T) {
	piPath := filepath.Clean("./relative/reference/path.yaml")
	oaPath := filepath.Clean("/this/base/path")

	i := path.Item{
		Ref:  piPath,
		Node: tree.Node{},
	}
	o := openapi3.OpenAPI{
		Node: tree.Node{},
	}
	o.SetBasePath(oaPath)
	o.SetChild("/testEndpoint", &i)
	i.SetParent(&o)
	want := filepath.Dir(filepath.Join(oaPath, piPath))
	got := i.GetBasePath()
	if got != want {
		t.Fail()
	}

}

func TestOpenAPI_PathItem_GetAndSetChildren(t *testing.T) {
	i := path.Item{}
	i.SetChild(http.MethodPost, &operation.Operation{Node: tree.Node{
		Name: http.MethodPost,
	}})
	i.SetChild(http.MethodGet, &operation.Operation{Node: tree.Node{
		Name: http.MethodGet,
	}})
	i.SetChild(http.MethodPatch, &operation.Operation{Node: tree.Node{
		Name: http.MethodPatch,
	}})
	i.SetChild(http.MethodDelete, &operation.Operation{Node: tree.Node{
		Name: http.MethodDelete,
	}})
	for method, op := range i.GetChildren() {
		switch method {
		case http.MethodGet:
			fallthrough
		case http.MethodDelete:
			fallthrough
		case http.MethodPost:
			fallthrough
		case http.MethodPatch:
			o, ok := op.(*operation.Operation)
			if !ok {
				t.Errorf("PathItem.operations() is not successfully returning *operations")
			}
			t.Logf("%v\n", o.GetName())
		}
	}
}

func TestOpenAPI_PathItem_GetRef(t *testing.T) {
	want := "testRef"
	i := path.Item{Ref: want}
	got := i.GetRef()
	if got != want {
		t.Fail()
	}
}
