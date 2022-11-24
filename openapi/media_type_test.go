package openapi

import (
	"reflect"
	"testing"
)

func TestMediaType_GetName(t *testing.T) {
	t.Errorf("Implement me")
}

func TestMediaType_GetChildren(t *testing.T) {
	openapi, _ := LoadOpenAPI(openapiFile)
	mt := Dig(openapi, testEndpoint, testMethod, testReqBody, testMediaType)
	children := mt.getChildren()
	for model, schema := range children {
		if model == testSchema {
			s, ok := schema.(*Schema)
			if !ok {
				t.Errorf("MediaType.getChildren() didn't return a *Schema")
			}
			if s == nil {
				t.Errorf("MediaType.getChildren() returned a nil schema")
			}
		}
	}
}

func TestMediaType_SetChild(t *testing.T) {
	openapi, _ := LoadOpenAPI(openapiFile)
	mt := Dig(openapi, testEndpoint, testMethod, testReqBody, testMediaType)
	s, _ := Dig(mt, testSchema).(*Schema)
	sOld := *s
	t.Logf("Original schema name: %v", s.name)
	newSchemaName := "NewSchema"
	mt.setChild(testSchema, &Schema{
		node: node{
			name: newSchemaName,
		},
	})
	sNew, _ := mt.getChildren()[testSchema].(*Schema)
	if reflect.DeepEqual(s, sNew) || (&sOld).name == sNew.name {
		t.Errorf("Child schema not set successfully")
	}
}
