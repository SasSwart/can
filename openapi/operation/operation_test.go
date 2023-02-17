package operation_test

import (
	"github.com/sasswart/gin-in-a-can/openapi/operation"
	"github.com/sasswart/gin-in-a-can/openapi/parameter"
	"github.com/sasswart/gin-in-a-can/openapi/request"
	"github.com/sasswart/gin-in-a-can/openapi/response"
	"github.com/sasswart/gin-in-a-can/openapi/root"
	"github.com/sasswart/gin-in-a-can/openapi/test"
	"github.com/sasswart/gin-in-a-can/tree"
	"testing"
)

func TestOpenAPI_Operation_GetBasePath(t *testing.T) {
	apiSpec, _ := root.LoadAPISpec(test.AbsOpenAPI)
	ops := test.Dig(apiSpec, test.Endpoint)

	var basePaths []string
	for method, op := range ops.GetChildren() {
		t.Logf("%v method found with base path: %v", method, op.GetBasePath())
		basePaths = append(basePaths, op.GetBasePath())
	}
	for _, path := range basePaths {
		if path != test.BasePath {
			t.Errorf("%v found, expected: %v", path, test.BasePath)
		}
	}
}

func TestOpenAPI_Operation_GetRef(t *testing.T) {
	apiSpec, _ := root.LoadAPISpec(test.AbsOpenAPI)
	ops := test.Dig(apiSpec, test.Endpoint)
	for _, op := range ops.GetChildren() {
		if o, ok := op.(*operation.Operation); ok {
			if o.GetRef() != "" {
				t.Errorf("%#v has an empty ref value", o)
			}
		}
	}
}

func TestOpenAPI_Operation_GetParent(t *testing.T) {
	apiSpec, _ := root.LoadAPISpec(test.AbsOpenAPI)
	ops := test.Dig(apiSpec, test.Endpoint)
	for _, o := range ops.GetChildren() {
		if o.GetParent() == nil {
			t.Errorf("operation %#v has a nil parent", o)
		}
	}
}

func TestOpenAPI_Operation_GetChildren(t *testing.T) {
	//apiSpec, _ := root.LoadAPISpec(test.AbsOpenAPI)
	//paths := test.Dig(apiSpec, test.Endpoint)
	//for _, traversable := range paths.GetChildren() {
	//	if operation, ok := traversable.(*Operation); ok {
	//		for key, child := range operation.getChildren() {
	//			if _, ok := child.(*request.Body); ok {
	//				t.Logf("%s contained an object of type *Body", key)
	//				t.Logf("heres your %s", test.ReqBody)
	//				continue
	//			}
	//			if p, ok := child.(*parameter.Parameter); ok {
	//				t.Logf("%s contained a *Parameter type", key)
	//				t.Logf("Parameter %s, in %s", p.name, p.In)
	//				continue
	//			}
	//			if _, ok := child.(*response.Response); ok {
	//				t.Logf("%s contained a *Response type", key)
	//				continue
	//			}
	//			t.Errorf("%s did not contain an object of type *Response, *Parameter, or *Body", key)
	//		}
	//	}
	//}
}

func TestOpenAPI_Operation_SetChild(t *testing.T) {
	apiSpec, _ := root.LoadAPISpec(test.AbsOpenAPI)
	traversable := test.Dig(apiSpec, test.Endpoint)
	operations := traversable.GetChildren()

	// Test Data
	p := parameter.Parameter{Node: tree.Node{Name: "newParameter"}}
	reqBody := request.Body{Node: tree.Node{Name: "newReqBody"}}
	r := response.Response{Node: tree.Node{Name: "newReqBody"}}
	httpResponseCode := "499"

	// Set children
	for method, op := range operations {
		t.Logf("Setting test children for %s method", method)
		op.SetChild("", &p)
		op.SetChild("", &reqBody)
		op.SetChild(httpResponseCode, &r)
	}

	// Verify set children
	//for method, op := range operations {
	//	t.Logf("Getting test children for %s method", method)
	//	children := op.GetChildren()
	//	if got, ok := children[httpResponseCode].(*r.Response); ok {
	//		if got.node.name != r.node.name {
	//			t.Errorf("child %s: %s was not set properly", method, httpResponseCode)
	//		}
	//	}
	//	if got, ok := children[test.ReqBody].(*request.Body); ok {
	//		if got.node.name != reqBody.node.name {
	//			t.Errorf("child %s: %s was not set properly", method, test.ReqBody)
	//		}
	//	}
	//	if got, ok := children[test.EmptyParamName].(*p.Parameter); ok {
	//		if got.node.name != p.node.name {
	//			t.Errorf("child %s: %s was not set properly", method, test.EmptyParamName)
	//		}
	//	}
	//}
}
