package openapi

import "testing"

func getOperation() map[string]Traversable {
	openapi, _ := LoadOpenAPI(openapiFile)
	return openapi.getChildren()["/endpoint"].(*PathItem).getChildren()
}

// Test Operation{}
func TestOperation_GetBasePath(t *testing.T) {
	ops := getOperation()
	var basePaths []string
	for method, operation := range ops {
		if op, ok := operation.(*Operation); ok {
			t.Logf("%v method found with base path: %v", method, op.getBasePath())
			basePaths = append(basePaths, op.getBasePath())
		}
	}
	for _, path := range basePaths {
		if path != expectedBasePath {
			t.Errorf("%v found, expected: %v", path, expectedBasePath)
			t.Fail()
		}
	}
}

func TestOperation_GetRef(t *testing.T) {
	ops := getOperation()
	for _, operation := range ops {
		if op, ok := operation.(*Operation); ok {
			if op.getRef() != "" {
				t.Fail()
			}
		}
	}
}

func TestOperation_GetParent(t *testing.T) {
	ops := getOperation()
	for _, operation := range ops {
		if op, ok := operation.(*Operation); ok {
			if op.getParent() != nil {
				t.Fail()
			}
		}
	}
}

func TestOperation_GetChildren(t *testing.T) {
	ops := getOperation()
	for _, operation := range ops {
		if op, ok := operation.(*Operation); ok {
			children := op.getChildren()
			if _, ok := children["RequestBody"].(*RequestBody); !ok {
				t.Errorf("Invalid requestBody")
				t.Fail()
			}
			for key, child := range children {
				if key == "RequestBody" {
					continue
				}
				if _, ok := child.(*Response); !ok {
					t.Errorf("invalid Response")
					t.Fail()
				}
			}

			// TODO implement testing for params
		}
	}

}

func TestOperation_SetChild(t *testing.T) {
	t.Errorf("Implement me")
	t.Fail()
}

//Test operationChildNode{}
