package openapi

import (
	"reflect"
	"testing"
)

func TestOpenAPI_MediaType_GetName(t *testing.T) {
	openapi, _ := LoadOpenAPI(testAbsOpenAPI)
	_ = SetRenderer(openapi, GinRenderer{})
	mt := Dig(openapi, testEndpoint, testMethod, testReqBody, testMediaType)
	name := mt.GetName()
	if name != testGinRenderedMediaItemName {
		t.Errorf("expected %v, got %v", testGinRenderedMediaItemName, name)
	}
}

func TestOpenAPI_MediaType_GetChildren(t *testing.T) {
	openapi, _ := LoadOpenAPI(testAbsOpenAPI)
	mt := Dig(openapi, testEndpoint, testMethod, testReqBody, testMediaType)
	children := mt.getChildren()
	for model, schema := range children {
		if model == testSchema {
			s, ok := schema.(*Schema)
			if !ok {
				// TODO These tests would be more valuable if we asserted against the content of the schema object we expect to make sure we get the right one.
				t.Errorf("MediaType.getChildren() didn't return a *Schema")
			}
			if s == nil {
				t.Errorf("MediaType.getChildren() returned a nil schema")
			}
		}
	}
}

func TestOpenAPI_MediaType_SetChild(t *testing.T) {
	openapi, _ := LoadOpenAPI(testAbsOpenAPI)
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
