package openapi

import (
	"github.com/sasswart/gin-in-a-can/render"
	"github.com/sasswart/gin-in-a-can/test"
	"reflect"
	"testing"
)

func TestMediaType_GetName(t *testing.T) {
	openAPI, _ := LoadOpenAPI(test.AbsOpenAPI)
	SetRenderer(openAPI, render.GinRenderer{})
	mt := Dig(openAPI, test.Endpoint, test.Method, test.ReqBody, test.MediaType)
	name := mt.GetName()
	if name != test.GinRenderedMediaItemName {
		t.Errorf("expected %v, got %v", test.GinRenderedMediaItemName, name)
	}
}

func TestMediaType_GetChildren(t *testing.T) {
	openapi, _ := LoadOpenAPI(test.AbsOpenAPI)
	mt := Dig(openapi, test.Endpoint, test.Method, test.ReqBody, test.MediaType)
	children := mt.getChildren()
	for model, schema := range children {
		if model == test.Schema {
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

func TestMediaType_SetChild(t *testing.T) {
	openapi, _ := LoadOpenAPI(test.AbsOpenAPI)
	mt := Dig(openapi, test.Endpoint, test.Method, test.ReqBody, test.MediaType)
	s, _ := Dig(mt, test.Schema).(*Schema)
	sOld := *s
	t.Logf("Original schema name: %v", s.name)
	newSchemaName := "NewSchema"
	mt.setChild(test.Schema, &Schema{
		node: node{
			name: newSchemaName,
		},
	})
	sNew, _ := mt.getChildren()[test.Schema].(*Schema)
	if reflect.DeepEqual(s, sNew) || (&sOld).name == sNew.name {
		t.Errorf("Child schema not set successfully")
	}
}
