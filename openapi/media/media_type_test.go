package media_test

import (
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/openapi/test"
	"github.com/sasswart/gin-in-a-can/tree"
	"reflect"
	"testing"
)

func TestOpenAPI_MediaType_GetName(t *testing.T) {
	apiSpec, _ := openapi.LoadAPISpec(test.AbsOpenAPI)
	//_ = root.SetRenderer(apiSpec, render.GinRenderer{})
	mt := test.Dig(apiSpec, test.Endpoint, test.Method, test.ReqBody, test.MediaType)
	name := mt.GetName()
	if name != test.GinRenderedMediaItemName {
		t.Errorf("expected %v, got %v", test.GinRenderedMediaItemName, name)
	}
}

func TestOpenAPI_MediaType_GetChildren(t *testing.T) {
	apiSpec, _ := openapi.LoadAPISpec(test.AbsOpenAPI)
	mt := test.Dig(apiSpec, test.Endpoint, test.Method, test.ReqBody, test.MediaType)
	children := mt.GetChildren()
	for model, s := range children {
		if model == test.Schema {
			s, ok := s.(*schema.Schema)
			if !ok {
				// TODO These test.s would be more valuable if we asserted against the content of the s object we expect to make sure we get the right one.
				t.Errorf("Type.getChildren() didn't return a *Schema")
			}
			if s == nil {
				t.Errorf("Type.getChildren() returned a nil s")
			}
		}
	}
}

func TestOpenAPI_MediaType_SetChild(t *testing.T) {
	apiSpec, _ := openapi.LoadAPISpec(test.AbsOpenAPI)
	mt := test.Dig(apiSpec, test.Endpoint, test.Method, test.ReqBody, test.MediaType)
	s, _ := test.Dig(mt, test.Schema).(*schema.Schema)
	sOld := *s
	t.Logf("Original schema name: %v", s.Name)
	newSchemaName := "NewSchema"
	mt.SetChild(test.Schema, &schema.Schema{
		Node: tree.Node{
			Name: newSchemaName,
		},
	})
	sNew, _ := mt.GetChildren()[test.Schema].(*schema.Schema)
	if reflect.DeepEqual(s, sNew) || (&sOld).Name == sNew.Name {
		t.Errorf("Child schema not set successfully")
	}
}
