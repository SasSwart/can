package openapi

import "testing"

func getOperation() map[string]Traversable {
	openapi, _ := LoadOpenAPI(openapiFile)
	return openapi.getChildren()["/endpoint"].(*PathItem).getChildren()
}

// Test Operation{}
func TestOperation_GetBasePath(t *testing.T) {
	openapi, _ := LoadOpenAPI(openapiFile)
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
	openapi, _ := LoadOpenAPI(openapiFile)
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
	openapi, _ := LoadOpenAPI(openapiFile)
	ops := Dig(openapi, testEndpoint)
	for _, operation := range ops.getChildren() {
		if operation.GetParent() != nil {
			t.Errorf("operation %#v has a nil parent", operation)
		}
	}
}

func TestOperation_GetChildren(t *testing.T) {
	openapi, _ := LoadOpenAPI(openapiFile)
	ops := Dig(openapi, testEndpoint)
	for _, traversable := range ops.getChildren() {
		if operation, ok := traversable.(*Operation); ok {
			for key, child := range operation.getChildren() {
				if key == testReqBody {
					continue
				}
				if _, ok := child.(*Response); !ok {
					t.Errorf("invalid Response")
				}
			}
			// TODO implement testing for params
		}
	}

}

func TestOperation_SetChild(t *testing.T) {
	t.Errorf("Implement me")
}

//Test operationChildNode{}
