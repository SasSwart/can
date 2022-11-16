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
		t.Logf("could not load file %s:%s", openapiFile, err.Error())
		t.Fail()
	}
	if openapi == nil {
		t.Logf("could not load file %s:%s", openapiFile, err.Error())
		t.Fail()
	}
	paths := openapi.getChildren()
	if paths == nil {
		t.Fail()
	}
}
func TestGetOpenAPIBasePath(t *testing.T) {
	openapi, _ := LoadOpenAPI(openapiFile)
	before, _, _ := strings.Cut(openapiFile, "/")
	if openapi.basePath != before {
		t.Logf("could not get basePath %s, got %s", before, openapi.basePath)
		t.Fail()
	}
}

func TestGetOpenAPIParent(t *testing.T) {
	openapi, _ := LoadOpenAPI(openapiFile)
	p := openapi.getParent()
	if p != nil {
		t.Logf("the root openapi file found a parent: %v", p)
		t.Fail()
	}
}

func TestGetOpenAPIChildren(t *testing.T) {
	openapi, _ := LoadOpenAPI(openapiFile)
	paths := openapi.getChildren()
	if len(paths) == 0 {
		t.Logf("error occured or test yaml file has no paths to get")
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
	t.Logf("could not find expected child in openapi file")
	t.Fail()
}

func TestOpenAPISetChild(t *testing.T) {
	openapi, _ := LoadOpenAPI(openapiFile)
	pathKey := "new"
	description := "new path item"
	p := PathItem{
		Description: description,
		Get:         nil,
		Post:        nil,
		Patch:       nil,
		Delete:      nil,
		Parameters:  nil,
		Ref:         "",
	}
	openapi.setChild(pathKey, &p)

	paths := openapi.getChildren()
	for k, v := range paths {
		if k == pathKey {
			p, ok := v.(*PathItem) // test that it's a *PathItem
			if !ok {
				t.Logf("Non-valid pathItem found")
				t.Fail()
			}
			if p.Description != description {
				t.Logf("new key description is %v when it should be %v", p.Description, description)
				t.Fail()
			}
		}
	}
}
