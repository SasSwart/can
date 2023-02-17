package operation

import (
	"github.com/sasswart/gin-in-a-can/openapi/parameter"
	"github.com/sasswart/gin-in-a-can/openapi/request_body"
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
	for method, operation := range ops.GetChildren() {
		t.Logf("%v method found with base path: %v", method, operation.GetBasePath())
		basePaths = append(basePaths, operation.GetBasePath())
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
	for _, operation := range ops.GetChildren() {
		if op, ok := operation.(*Operation); ok {
			if op.getRef() != "" {
				t.Errorf("%#v has an empty ref value", op)
			}
		}
	}
}

func TestOpenAPI_Operation_GetParent(t *testing.T) {
	apiSpec, _ := root.LoadAPISpec(test.AbsOpenAPI)
	ops := test.Dig(apiSpec, test.Endpoint)
	for _, operation := range ops.GetChildren() {
		if operation.GetParent() == nil {
			t.Errorf("operation %#v has a nil parent", operation)
		}
	}
}

func TestOpenAPI_Operation_GetChildren(t *testing.T) {
	apiSpec, _ := root.LoadAPISpec(test.AbsOpenAPI)
	paths := test.Dig(apiSpec, test.Endpoint)
	for _, traversable := range paths.GetChildren() {
		if operation, ok := traversable.(*Operation); ok {
			for key, child := range operation.getChildren() {
				if _, ok := child.(*request_body.RequestBody); ok {
					t.Logf("%s contained an object of type *RequestBody", key)
					t.Logf("heres your %s", test.ReqBody)
					continue
				}
				if p, ok := child.(*parameter.Parameter); ok {
					t.Logf("%s contained a *Parameter type", key)
					t.Logf("Parameter %s, in %s", p.name, p.In)
					continue
				}
				if _, ok := child.(*response.Response); ok {
					t.Logf("%s contained a *Response type", key)
					continue
				}
				t.Errorf("%s did not contain an object of type *Response, *Parameter, or *RequestBody", key)
			}
		}
	}
}

func TestOpenAPI_Operation_SetChild(t *testing.T) {
	apiSpec, _ := root.LoadAPISpec(test.AbsOpenAPI)
	traversable := test.Dig(apiSpec, test.Endpoint)
	operations := traversable.GetChildren()

	// Test Data
	parameter := parameter.Parameter{Node: tree.Node{name: "newParameter"}}
	reqBody := request_body.RequestBody{Node: tree.Node{name: "newReqBody"}}
	response := response.Response{Node: tree.Node{name: "newReqBody"}}
	httpResponseCode := "499"

	// Set children
	for method, op := range operations {
		t.Logf("Setting test children for %s method", method)
		op.SetChild("", &parameter)
		op.SetChild("", &reqBody)
		op.SetChild(httpResponseCode, &response)
	}

	// Verify set children
	for method, op := range operations {
		t.Logf("Getting test children for %s method", method)
		children := op.GetChildren()
		if got, ok := children[httpResponseCode].(*response.Response); ok {
			if got.node.name != response.node.name {
				t.Errorf("child %s: %s was not set properly", method, httpResponseCode)
			}
		}
		if got, ok := children[test.ReqBody].(*request_body.RequestBody); ok {
			if got.node.name != reqBody.node.name {
				t.Errorf("child %s: %s was not set properly", method, test.ReqBody)
			}
		}
		if got, ok := children[test.EmptyParamName].(*parameter.Parameter); ok {
			if got.node.name != parameter.node.name {
				t.Errorf("child %s: %s was not set properly", method, test.EmptyParamName)
			}
		}
	}
}
