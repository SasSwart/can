package openapi

import (
	"github.com/sasswart/gin-in-a-can/render"
	"github.com/sasswart/gin-in-a-can/test"
	"reflect"
	"testing"
)

func TestPathItem_GetName(t *testing.T) {
	openapi, _ := LoadOpenAPI(test.OpenAPIFile)
	render.SetRenderer(openapi, render.GinRenderer{})
	path := Dig(openapi, test.Endpoint)
	if path.getRenderer() == nil {
		t.Log("Renderer is nil, setting render manually")
		path.setRenderer(render.GinRenderer{})
	}
	if path.GetName() != test.GinRenderedPathItemName {
		t.Errorf("got %v, expected %v", path.GetName(), test.GinRenderedPathItemName)
	}
}

func TestPathItem_Operations(t *testing.T) {
	openapi, _ := LoadOpenAPI(test.AbsOpenAPI)
	for _, v := range openapi.getChildren() {
		p := v.(*PathItem)
		for method, op := range p.Operations() {
			switch method {
			case "post":
			case "get":
			case "patch":
			case "delete":
				o, ok := op.(*Operation)
				if !ok {
					t.Errorf("PathItem.Operations() is not successfully returning *Operations")
				}
				t.Logf(o.OperationId)
			}
		}
	}

}

func TestPathItem_SetRenderer(t *testing.T) {
	openapi, _ := LoadOpenAPI(test.AbsOpenAPI)
	for _, path := range openapi.getChildren() {
		path.setRenderer(render.GinRenderer{})
		GinRenderer := render.GinRenderer{}
		if !reflect.DeepEqual(path.getRenderer(), GinRenderer) {
			t.Errorf("SetRenderer(GinRenderer{}) was unsuccessful")
		}
	}
}

func TestPathItem_GetBasePath(t *testing.T) {
	openapi, _ := LoadOpenAPI(test.AbsOpenAPI)
	for _, path := range openapi.getChildren() {
		if path.getBasePath() != test.BasePath {
			t.Errorf("got %v, expected %v", path.getBasePath(), test.BasePath)
		}
	}
}

func TestPathItem_GetParent(t *testing.T) {
	openapi, _ := LoadOpenAPI(test.AbsOpenAPI)
	for _, path := range openapi.getChildren() {
		parent := path.GetParent()
		_, ok := parent.(*OpenAPI)
		if !ok {
			t.Errorf("PathItem.GetParent() did not return an OpenAPI type")
		}
	}
}
