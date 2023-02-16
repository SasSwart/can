package openapi

import (
	"reflect"
	"testing"
)

func TestOpenAPI_PathItem_GetName(t *testing.T) {
	openapi, _ := LoadOpenAPI(testOpenapiFile)
	_ = SetRenderer(openapi, GinRenderer{})
	path := Dig(openapi, testEndpoint)
	if path.GetName() != testGinRenderedPathItemName {
		t.Errorf("got %v, expected %v", path.GetName(), testGinRenderedPathItemName)
	}
}

func TestOpenAPI_PathItem_Operations(t *testing.T) {
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

func TestOpenAPI_PathItem_SetRenderer(t *testing.T) {
	openapi, _ := LoadOpenAPI(testAbsOpenAPI)
	for _, path := range openapi.getChildren() {
		want := GinRenderer{}
		path.setRenderer(want)
		got := path.getRenderer()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("SetRenderer(GinRenderer{}) was unsuccessful")
		}
	}
}

func TestOpenAPI_PathItem_GetBasePath(t *testing.T) {
	openapi, _ := LoadOpenAPI(testAbsOpenAPI)
	for _, path := range openapi.getChildren() {
		if path.getBasePath() != testBasePath {
			t.Errorf("got %v, expected %v", path.getBasePath(), testBasePath)
		}
	}
}

func TestOpenAPI_PathItem_GetParent(t *testing.T) {
	openapi, _ := LoadOpenAPI(testAbsOpenAPI)
	for _, path := range openapi.getChildren() {
		parent := path.GetParent()
		_, ok := parent.(*OpenAPI)
		if !ok {
			t.Errorf("PathItem.GetParent() did not return an OpenAPI type")
		}
	}
}
