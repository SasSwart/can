package path_test

import (
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/openapi/operation"
	"github.com/sasswart/gin-in-a-can/openapi/path"
	"github.com/sasswart/gin-in-a-can/openapi/test"
	"github.com/sasswart/gin-in-a-can/render"
	"testing"
)

func TestOpenAPI_PathItem_GetName(t *testing.T) {
	apiSpec, _ := openapi.LoadAPISpec(test.AbsOpenAPI)
	_ = render.Engine{}.New(render.GinRenderer{}, render.Config{})
	p := test.Dig(apiSpec, test.Endpoint)
	if p.GetName() != test.GinRenderedPathItemName {
		t.Errorf("got %v, expected %v", p.GetName(), test.GinRenderedPathItemName)
	}
}

func TestOpenAPI_PathItem_Operations(t *testing.T) {
	apiSpec, _ := openapi.LoadAPISpec(test.AbsOpenAPI)
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

func TestOpenAPI_PathItem_GetBasePath(t *testing.T) {
	apiSpec, _ := openapi.LoadAPISpec(test.AbsOpenAPI)
	for _, p := range apiSpec.GetChildren() {
		if p.GetBasePath() != test.BasePath {
			t.Errorf("got %v, expected %v", p.GetBasePath(), test.BasePath)
		}
	}
}

func TestOpenAPI_PathItem_GetParent(t *testing.T) {
	apiSpec, _ := openapi.LoadAPISpec(test.AbsOpenAPI)
	for _, p := range apiSpec.GetChildren() {
		parent := p.GetParent()
		_, ok := parent.(*openapi.OpenAPI)
		if !ok {
			t.Errorf("PathItem.GetParent() did not return an OpenAPI type")
		}
	}
}
