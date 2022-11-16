package openapi

import (
	"strings"
	"testing"
)

var openapiFile = "fixtures/validation.yaml"
var expectedPathInFile = "/endpoint"

func TestLoadOpenAPI(t *testing.T) {
	openapi, err := LoadOpenAPI(openapiFile)
	if err != nil {
		t.Errorf("could not load file %s:%s", openapiFile, err.Error())
		t.Fail()
	}
	if openapi == nil {
		t.Errorf("could not load file %s:%s", openapiFile, err.Error())
		t.Fail()
	}
	paths := openapi.getChildren()
	if paths == nil {
		t.Fail()
	}
}

func TestSetRenderer(t *testing.T) {
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
	//for key, operation := range operations {
	//	if operation != nil {
	//		o := operation.(*Operation)
	//		if o.getRenderer() == nil {
	//			t.Errorf("Operation %v is missing a renderer", key)
	//			t.Fail()
	//		}
	//	}
	//}

	//Check *RequestBody
	requestBody := operations["post"].(*Operation).getChildren()
	for key, mediaTypes := range requestBody {
		r := mediaTypes.(*RequestBody)
		if r.getRenderer() == nil {
			t.Errorf("RequestBody %v is missing a renderer", key)
			t.Fail()
		}
	}

	//Check *Response
	//Check *Parameter
	//Check *MediaType
	//Check *Schema
	//Check *Schema within *Schema

}

func TestGetOpenAPIBasePath(t *testing.T) {
	openapi, _ := LoadOpenAPI(openapiFile)
	before, _, _ := strings.Cut(openapiFile, "/")
	if openapi.basePath != before {
		t.Errorf("could not get basePath %s, got %s", before, openapi.basePath)
		t.Fail()
	}
}

func TestGetOpenAPIParent(t *testing.T) {
	openapi, _ := LoadOpenAPI(openapiFile)
	p := openapi.getParent()
	if p != nil {
		t.Errorf("the root openapi file found a parent: %v", p)
		t.Fail()
	}
}

func TestGetOpenAPIChildren(t *testing.T) {
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

func TestOpenAPISetChild(t *testing.T) {
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
