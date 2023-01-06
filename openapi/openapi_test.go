package openapi

import (
	"github.com/sasswart/gin-in-a-can/render"
	"github.com/sasswart/gin-in-a-can/test"
	"path/filepath"
	"reflect"
	"testing"
)

func TestOpenAPI_LoadOpenAPI(t *testing.T) {
	openapi, err := LoadOpenAPI(test.AbsOpenAPI)
	if err != nil {
		t.Errorf("could not load file %s:%s", test.OpenAPIFile, err.Error())
	}
	if openapi == nil {
		t.Errorf("could not load file %s:%s", test.OpenAPIFile, err.Error())
	}
}

func TestOpenAPI_SetRenderer(t *testing.T) {
	openapi, _ := LoadOpenAPI(test.AbsOpenAPI)
	render.SetRenderer(openapi, render.GinRenderer{})

	//Check *PathItem
	paths := openapi.getChildren()
	for key, path := range paths {
		p := path
		if p.getRenderer() == nil {
			t.Errorf("Path %v is missing a render", key)
		}
	}

	//Check *Operation
	operations := paths[test.Endpoint].getChildren()
	// TODO currently faiing
	for key, operation := range operations {
		if operation != nil {
			o := operation
			if o.getRenderer() == nil {
				t.Errorf("Operation %v is missing a render", key)
			}
		}
	}

	//Check *RequestBody
	requestBodyAndResponses := operations[test.Method].getChildren()
	// TODO currently faiing
	for key, mediaTypes := range requestBodyAndResponses {
		if key == "RequestBody" {
			r := mediaTypes
			if r.getRenderer() == nil {
				t.Errorf("%v is missing a render", key)
			}
			continue
		}
		//Check *Response
		r := mediaTypes
		if r.getRenderer() == nil {
			t.Errorf("Response %v is missing a render", key)
		}
	}

	//Check *Parameter
	// TODO *PathItem.getParameters() still to be done

	//Check *MediaType
	//Check *Schema and *Schema within *Schemae
	//	RequestBody
	requestBodyMediaTypes := requestBodyAndResponses["RequestBody"].getChildren()
	for mt, val := range requestBodyMediaTypes {
		if mt == test.MediaType {
			container, _ := val.(*MediaType)
			schemae := container.getChildren()
			schema, ok := schemae[test.Schema].(*Schema)
			if !ok {
				t.Errorf("Schema cast failed")
			}
			if schema.getRenderer() == nil {
				t.Errorf("RequestBody Schema of  %v is missing a render", mt)
			}
			c := schema.getChildren()
			for propKey, s := range c {
				if propKey == "item" { // check everything but "item"
					continue
				}
				s, ok := s.(*Schema)
				if !ok {
					t.Errorf("Invalid schema nesting")
				}
				if s.getRenderer() == nil {
					t.Errorf("%v:%v property schema is missing a render", mt, propKey)
				}
			}
		}
	}
	//	Responses
	for k, val := range requestBodyAndResponses {
		if k == "RequestBody" {
			continue
		}
		content := val.getChildren()
		if content == nil {
			continue
		}
		//for mt, val := range content {
		//	schemae := val.getChildren()
		//	schema, ok := schemae[Schema].(*Schema)
		//if !ok {
		//	// TODO find out why this cast is failing
		//	t.Errorf("Schema cast failed")
		//}
		//if schema.getRenderer() == nil {
		//	t.Errorf("%v Response Schema of type %v is missing a render", k, mt)
		//}
		//c := schema.getChildren()
		//for propKey, s := range c {
		//	if propKey == "item" { // check everything but "item"
		//		continue
		//	}
		//	s, ok := s.(*Schema)
		//	if !ok {
		//		t.Errorf("Invalid schema nesting")
		//	}
		//	if s.getRenderer() == nil {
		//		t.Errorf("%v Response Schema of type %v is missing a render", k, mt)
		//	}
		//}
		//}
	}
	//Parameter
}

func TestOpenAPI_GetBasePath(t *testing.T) {
	openapi, _ := LoadOpenAPI(test.AbsOpenAPI)
	wanted := filepath.Dir(test.AbsOpenAPI)
	if openapi.getBasePath() != wanted {
		t.Errorf("could not get basePath %s, got %s", wanted, openapi.basePath)
	}
}

func TestOpenAPI_GetParent(t *testing.T) {
	openapi, _ := LoadOpenAPI(test.AbsOpenAPI)
	p := openapi.GetParent()
	if p != nil {
		t.Errorf("the root openapi file found a parent: %v", p)
	}
}

func TestGetOpenAPI_GetChildren(t *testing.T) {
	openapi, _ := LoadOpenAPI(test.AbsOpenAPI)
	paths := openapi.getChildren()
	if len(paths) == 0 {
		t.Errorf("error occured or test yaml file has no paths to get")
	}
	for k, v := range paths {
		if k == test.Endpoint {
			if _, ok := v.(*PathItem); ok {
				return // test that it's a *PathItem
			}
		}
	}
	t.Errorf("could not find a valid child in openapi file")
}

func TestOpenAPI_SetChild(t *testing.T) {
	openapi, _ := LoadOpenAPI(test.AbsOpenAPI)
	pathKey := "new"
	p := PathItem{
		Description: "new path item",
	}
	openapi.setChild(pathKey, &p)

	paths := openapi.getChildren()
	for k, v := range paths {
		if k == pathKey {
			path, ok := v.(*PathItem) // test that it's a *PathItem
			if !ok {
				t.Errorf("Non-valid pathItem found")
			}
			if !reflect.DeepEqual(*path, p) {
				t.Errorf("path item set is not equivalent to path item retrieved")
			}
		}
	}
}

func TestOpenAPI_GetName(t *testing.T) {
	openapi, _ := LoadOpenAPI(test.AbsOpenAPI)
	render.SetRenderer(openapi, render.GinRenderer{})
	name := openapi.GetName()
	if name != test.GinRenderedOpenAPIName {
		t.Errorf("wanted %s, got %s", test.GinRenderedOpenAPIName, name)
	}
}

func TestOpenAPI_ResolveRefs(t *testing.T) {
	// TODO
	//(key string, parent, child Traversable) (Traversable, error) {
}
