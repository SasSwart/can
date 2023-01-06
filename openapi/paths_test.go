package openapi

import (
	"reflect"
	"testing"
)

func TestPathItem_GetName(t *testing.T) {
	openapi, _ := LoadOpenAPI(testOpenapiFile)
	SetRenderer(openapi, GinRenderer{})
	path := Dig(openapi, testEndpoint)
	if path.getRenderer() == nil {
		t.Log("Renderer is nil, setting renderer manually")
		path.setRenderer(GinRenderer{})
	}
	if path.GetName() != testGinRenderedPathItemName {
		t.Errorf("got %v, expected %v", path.GetName(), testGinRenderedPathItemName)
	}
}

func TestPathItem_Operations(t *testing.T) {
	openapi, _ := LoadOpenAPI(testAbsOpenAPI)
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
	openapi, _ := LoadOpenAPI(testAbsOpenAPI)
	for _, path := range openapi.getChildren() {
		path.setRenderer(GinRenderer{})
		GinRenderer := GinRenderer{}
		if !reflect.DeepEqual(path.getRenderer(), GinRenderer) {
			t.Errorf("SetRenderer(GinRenderer{}) was unsuccessful")
		}
	}
}

func TestPathItem_GetBasePath(t *testing.T) {
	openapi, _ := LoadOpenAPI(testAbsOpenAPI)
	for _, path := range openapi.getChildren() {
		if path.getBasePath() != testBasePath {
			t.Errorf("got %v, expected %v", path.getBasePath(), testBasePath)
		}
	}
}

func TestPathItem_GetParent(t *testing.T) {
	openapi, _ := LoadOpenAPI(testAbsOpenAPI)
	for _, path := range openapi.getChildren() {
		parent := path.GetParent()
		_, ok := parent.(*OpenAPI)
		if !ok {
			t.Errorf("PathItem.GetParent() did not return an OpenAPI type")
		}
	}
}
