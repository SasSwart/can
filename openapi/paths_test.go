package openapi

import (
	"reflect"
	"testing"
)

func TestPathItem_GetName(t *testing.T) {
	openapi, _ := LoadOpenAPI(openapiFile)
	expected := "StratusEndpoint" // TODO make more robust
	for _, v := range openapi.getChildren() {
		p := v.(*PathItem)
		if p.getRenderer() == nil {
			t.Log("Renderer is nil, setting renderer manually")
			p.SetRenderer(GinRenderer{})
		}
		if p.GetName() != expected {
			t.Errorf("got %v, expected %v", p.GetName(), expected)
			t.Fail()
		}
	}
}

func TestPathItem_Operations(t *testing.T) {
	openapi, _ := LoadOpenAPI(openapiFile)
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
					t.Fail()
				}
				t.Logf(o.OperationId)
			}
		}
	}

}

func TestPathItem_SetRenderer(t *testing.T) {
	openapi, _ := LoadOpenAPI(openapiFile)
	for _, v := range openapi.getChildren() {
		p := v.(*PathItem)
		p.SetRenderer(GinRenderer{})
		GinRenderer := GinRenderer{}
		if !reflect.DeepEqual(p.getRenderer(), GinRenderer) {
			t.Errorf("SetRenderer(GinRenderer{}) was unsuccessful")
			t.Fail()
		}
	}
}

func TestPathItem_GetBasePath(t *testing.T) {
	openapi, _ := LoadOpenAPI(openapiFile)
	expected := "fixtures"
	for _, v := range openapi.getChildren() {
		p := v.(*PathItem)
		if p.getBasePath() != expected {
			t.Errorf("got %v, expected %v", p.getBasePath(), expected)
			t.Fail()
		}
	}
}

func TestPathItem_GetParent(t *testing.T) {
	openapi, _ := LoadOpenAPI(openapiFile)
	for _, v := range openapi.getChildren() {
		p := v.(*PathItem)
		parent := p.getParent()
		_, ok := parent.(*OpenAPI)
		if !ok {
			t.Errorf("PathItem.getParent() did not return an OpenAPI type")
			t.Fail()
		}
	}
}
