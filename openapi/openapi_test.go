package openapi

import (
	"strings"
	"testing"
)

var openapiFile = "fixtures/validation.yaml"
var expectedPathInFile = "/endpoint"

func TestOpenAPI_LoadOpenAPI(t *testing.T) {
	openapi, err := LoadOpenAPI(openapiFile)
	if err != nil {
		t.Errorf("could not load file %s:%s", openapiFile, err.Error())
		t.Fail()
	}
	if openapi == nil {
		t.Errorf("could not load file %s:%s", openapiFile, err.Error())
		t.Fail()
	}
}

func TestOpenAPI_SetRenderer(t *testing.T) {
	openapi, _ := LoadOpenAPI(openapiFile)
	SetRenderer(openapi, GinRenderer{})

	//Check *PathItem
	paths := openapi.getChildren()
	for key, path := range paths {
		p := path.(*PathItem)
		if p.getRenderer() == nil {
			t.Errorf("Path %v is missing a renderer", key)
			t.Fail()
		}
	}

	//Check *Operation
	operations := paths["/endpoint"].(*PathItem).getChildren()
	// TODO currently faiing
	for key, operation := range operations {
		if operation != nil {
			o := operation.(*Operation)
			if o.getRenderer() == nil {
				t.Errorf("Operation %v is missing a renderer", key)
				t.Fail()
			}
		}
	}

	//Check *RequestBody
	requestBodyAndResponses := operations["post"].(*Operation).getChildren()
	// TODO currently faiing
	for key, mediaTypes := range requestBodyAndResponses {
		if key == "RequestBody" {
			r := mediaTypes.(*RequestBody)
			if r.getRenderer() == nil {
				t.Errorf("%v is missing a renderer", key)
				t.Fail()
			}
			continue
		}
		//Check *Response
		r := mediaTypes.(*Response)
		if r.getRenderer() == nil {
			t.Errorf("Response %v is missing a renderer", key)
			t.Fail()
		}
	}

	//Check *Parameter
	// TODO *PathItem.getParameters() still to be done

	//Check *MediaType
	//Check *Schema and *Schema within *Schemae
	//	RequestBody
	requestBodyMediaTypes := requestBodyAndResponses["RequestBody"].getChildren()
	for mt, val := range requestBodyMediaTypes {
		if mt == "application/json" {
			container, _ := val.(*MediaType)
			schemae := container.getChildren()
			schema, ok := schemae["Model"].(*Schema)
			if !ok {
				t.Errorf("Schema cast failed")
				t.Fail()
			}
			if schema.getRenderer() == nil {
				t.Errorf("RequestBody Schema of  %v is missing a renderer", mt)
				t.Fail()
			}
			c := schema.getChildren()
			for propKey, s := range c {
				if propKey == "item" { // check everything but "item"
					continue
				}
				s, ok := s.(*Schema)
				if !ok {
					t.Errorf("Invalid schema nesting")
					t.Fail()
				}
				if s.getRenderer() == nil {
					t.Errorf("%v:%v property schema is missing a renderer", mt, propKey)
					t.Fail()
				}
			}
		}
	}
	//	Responses
	for k, val := range requestBodyAndResponses {
		if k == "RequestBody" {
			continue
		}
		content := val.(*Response).getChildren()
		if content == nil {
			continue
		}
		for mt, val := range content {
			container, _ := val.(*MediaType)
			schemae := container.getChildren()
			schema, ok := schemae["Model"].(*Schema)
			if !ok {
				t.Errorf("Schema cast failed")
				t.Fail()
			}
			if schema.getRenderer() == nil {
				t.Errorf("%v Response Schema of type %v is missing a renderer", k, mt)
				t.Fail()
			}
			c := schema.getChildren()
			for propKey, s := range c {
				if propKey == "item" { // check everything but "item"
					continue
				}
				s, ok := s.(*Schema)
				if !ok {
					t.Errorf("Invalid schema nesting")
					t.Fail()
				}
				if s.getRenderer() == nil {
					t.Errorf("%v Response Schema of type %v is missing a renderer", k, mt)
					t.Fail()
				}
			}
		}
	}
	//Parameter
}

func TestGetOpenAPIBasePath(t *testing.T) {
	openapi, _ := LoadOpenAPI(openapiFile)
	before, _, _ := strings.Cut(openapiFile, "/")
	if openapi.basePath != before {
		t.Errorf("could not get basePath %s, got %s", before, openapi.basePath)
		t.Fail()
	}
}

func TestGetOpenAPI_Parent(t *testing.T) {
	openapi, _ := LoadOpenAPI(openapiFile)
	p := openapi.getParent()
	if p != nil {
		t.Errorf("the root openapi file found a parent: %v", p)
		t.Fail()
	}
}

func TestGetOpenAPI_Children(t *testing.T) {
	openapi, _ := LoadOpenAPI(openapiFile)
	paths := openapi.getChildren()
	if len(paths) == 0 {
		t.Errorf("error occured or test yaml file has no paths to get")
		t.Fail()
	}
	for k, v := range paths {
		if k == expectedPathInFile {
			_, ok := v.(*PathItem) // test that it's a *PathItem
			if ok {
				return
			}
		}
	}
	t.Errorf("could not find expected child in openapi file")
	t.Fail()
}

func TestOpenAPI_SetChild(t *testing.T) {
	openapi, _ := LoadOpenAPI(openapiFile)
	pathKey := "new"
	description := "new path item"
	p := PathItem{
		Description: description,
	}
	openapi.setChild(pathKey, &p)

	paths := openapi.getChildren()
	for k, v := range paths {
		if k == pathKey {
			p, ok := v.(*PathItem) // test that it's a *PathItem
			if !ok {
				t.Errorf("Non-valid pathItem found")
				t.Fail()
			}
			if p.Description != description {
				t.Errorf("new key description is %v when it should be %v", p.Description, description)
				t.Fail()
			}
		}
	}
}

func TestOpenAPI_ResolveRefs(t *testing.T) {
	//(key string, parent, child Traversable) (Traversable, error) {
}
