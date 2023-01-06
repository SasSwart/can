package openapi

import (
	"testing"
)

func TestOperation_GetBasePath(t *testing.T) {
	openapi, _ := LoadOpenAPI(testAbsOpenAPI)
	ops := Dig(openapi, testEndpoint)

	var basePaths []string
	for method, operation := range ops.getChildren() {
		t.Logf("%v method found with base path: %v", method, operation.getBasePath())
		basePaths = append(basePaths, operation.getBasePath())
	}
	for _, path := range basePaths {
		if path != testBasePath {
			t.Errorf("%v found, expected: %v", path, testBasePath)
		}
	}
}

func TestOperation_GetRef(t *testing.T) {
	openapi, _ := LoadOpenAPI(testAbsOpenAPI)
	ops := Dig(openapi, testEndpoint)
	for _, operation := range ops.getChildren() {
		if op, ok := operation.(*Operation); ok {
			if op.getRef() != "" {
				t.Errorf("%#v has an empty ref value", op)
			}
		}
	}
}

func TestOperation_GetParent(t *testing.T) {
	openapi, _ := LoadOpenAPI(testAbsOpenAPI)
	ops := Dig(openapi, testEndpoint)
	for _, operation := range ops.getChildren() {
		if operation.GetParent() == nil {
			t.Errorf("operation %#v has a nil parent", operation)
		}
	}
}

func TestOperation_GetChildren(t *testing.T) {
	openapi, _ := LoadOpenAPI(testAbsOpenAPI)
	paths := Dig(openapi, testEndpoint)
	for _, traversable := range paths.getChildren() {
		if operation, ok := traversable.(*Operation); ok {
			for key, child := range operation.getChildren() {
				if _, ok := child.(*RequestBody); ok {
					t.Logf("%s contained an object of type *RequestBody", key)
					t.Logf("heres your %s", testReqBody)
					continue
				}
				if p, ok := child.(*Parameter); ok {
					t.Logf("%s contained a *Parameter type", key)
					t.Logf("Parameter %s, in %s", p.name, p.In)
					continue
				}
				if _, ok := child.(*Response); ok {
					t.Logf("%s contained a *Response type", key)
					continue
				}
				t.Errorf("%s did not contain an object of type *Response, *Parameter, or *RequestBody", key)
			}
		}
	}
}

func TestOperation_SetChild(t *testing.T) {
	openapi, _ := LoadOpenAPI(testAbsOpenAPI)
	traversable := Dig(openapi, testEndpoint)
	operations := traversable.getChildren()

	// Test Data
	parameter := Parameter{node: node{name: "newParameter"}}
	reqBody := RequestBody{node: node{name: "newReqBody"}}
	response := Response{node: node{name: "newReqBody"}}
	httpResponseCode := "499"

	// Set children
	for method, op := range operations {
		t.Logf("Setting test children for %s method", method)
		op.setChild("", &parameter)
		op.setChild("", &reqBody)
		op.setChild(httpResponseCode, &response)
	}

	// Verify set children
	for method, op := range operations {
		t.Logf("Getting test children for %s method", method)
		children := op.getChildren()
		if got, ok := children[httpResponseCode].(*Response); ok {
			if got.node.name != response.node.name {
				t.Errorf("child %s: %s was not set properly", method, httpResponseCode)
			}
		}
		if got, ok := children[testReqBody].(*RequestBody); ok {
			if got.node.name != reqBody.node.name {
				t.Errorf("child %s: %s was not set properly", method, testReqBody)
			}
		}
		if got, ok := children[testEmptyParamName].(*Parameter); ok {
			if got.node.name != parameter.node.name {
				t.Errorf("child %s: %s was not set properly", method, testEmptyParamName)
			}
		}
	}
}
