package path_test

import (
	"github.com/sasswart/gin-in-a-can/openapi/operation"
	"github.com/sasswart/gin-in-a-can/openapi/path"
	"github.com/sasswart/gin-in-a-can/openapi/root"
	"github.com/sasswart/gin-in-a-can/openapi/test"
	"testing"
)

func TestOpenAPI_PathItem_GetName(t *testing.T) {
	apiSpec, _ := root.LoadAPISpec(test.AbsOpenAPI)
	//_ = root.SetRenderer(apiSpec, render.GinRenderer{})
	path := test.Dig(apiSpec, test.Endpoint)
	if path.GetName() != test.GinRenderedPathItemName {
		t.Errorf("got %v, expected %v", path.GetName(), test.GinRenderedPathItemName)
	}
}

func TestOpenAPI_PathItem_Operations(t *testing.T) {
	apiSpec, _ := root.LoadAPISpec(test.AbsOpenAPI)
	for _, v := range apiSpec.GetChildren() {
		p := v.(*path.Item)
		for method, op := range p.Operations() {
			switch method {
			case "post":
			case "get":
			case "patch":
			case "delete":
				o, ok := op.(*operation.Operation)
				if !ok {
					t.Errorf("PathItem.Operations() is not successfully returning *Operations")
				}
				t.Logf(o.OperationId)
			}
		}
	}

}

func TestOpenAPI_PathItem_SetRenderer(t *testing.T) {
	//apiSpec, _ := root.LoadAPISpec(test.AbsOpenAPI)
	//for _, path := range apiSpec.GetChildren() {
	//	want := render.GinRenderer{}
	//	path.SetRenderer(want)
	//	got := path.GetRenderer()
	//	if !reflect.DeepEqual(got, want) {
	//		t.Errorf("SetRenderer(GinRenderer{}) was unsuccessful")
	//	}
	//}
}

func TestOpenAPI_PathItem_GetBasePath(t *testing.T) {
	apiSpec, _ := root.LoadAPISpec(test.AbsOpenAPI)
	for _, path := range apiSpec.GetChildren() {
		if path.GetBasePath() != test.BasePath {
			t.Errorf("got %v, expected %v", path.GetBasePath(), test.BasePath)
		}
	}
}

func TestOpenAPI_PathItem_GetParent(t *testing.T) {
	apiSpec, _ := root.LoadAPISpec(test.AbsOpenAPI)
	for _, path := range apiSpec.GetChildren() {
		parent := path.GetParent()
		_, ok := parent.(*root.Root)
		if !ok {
			t.Errorf("PathItem.GetParent() did not return an Root type")
		}
	}
}
