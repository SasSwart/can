package openapi

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestOpenAPI_LoadOpenAPI(t *testing.T) {
	openapi, err := LoadOpenAPI(testAbsOpenAPI)
	if err != nil {
		t.Errorf("could not load file %s:%s", testOpenapiFile, err.Error())
	}
	if openapi == nil {
		t.Errorf("could not load file %s:%s", testOpenapiFile, err.Error())
	}
}

func TestOpenAPI_SetRenderer(t *testing.T) {
	openapi, _ := LoadOpenAPI(testAbsOpenAPI)
	SetRenderer(openapi, GinRenderer{})

	//Check *PathItem
	paths := openapi.getChildren()
	for key, path := range paths {
		p := path
		if p.getRenderer() == nil {
			t.Errorf("Path %v is missing a renderer", key)
		}
	}

	//Check *Operation
	operations := paths[testEndpoint].getChildren()
	// TODO currently faiing
	for key, operation := range operations {
		if operation != nil {
			o := operation
			if o.getRenderer() == nil {
				t.Errorf("Operation %v is missing a renderer", key)
			}
		}
	}

	//Check *RequestBody
	requestBodyAndResponses := operations[testMethod].getChildren()
	// TODO currently faiing
	for key, mediaTypes := range requestBodyAndResponses {
		if key == "RequestBody" {
			r := mediaTypes
			if r.getRenderer() == nil {
				t.Errorf("%v is missing a renderer", key)
			}
			continue
		}
		//Check *Response
		r := mediaTypes
		if r.getRenderer() == nil {
			t.Errorf("Response %v is missing a renderer", key)
		}
	}

	//Check *Parameter
	// TODO *PathItem.getParameters() still to be done

	//Check *MediaType
	//Check *Schema and *Schema within *Schemae
	//	RequestBody
	requestBodyMediaTypes := requestBodyAndResponses["RequestBody"].getChildren()
	for mt, val := range requestBodyMediaTypes {
		if mt == testMediaType {
			container, _ := val.(*MediaType)
			schemae := container.getChildren()
			schema, ok := schemae[testSchema].(*Schema)
			if !ok {
				t.Errorf("Schema cast failed")
			}
			if schema.getRenderer() == nil {
				t.Errorf("RequestBody Schema of  %v is missing a renderer", mt)
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
					t.Errorf("%v:%v property schema is missing a renderer", mt, propKey)
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
		//	schema, ok := schemae[testSchema].(*Schema)
		//if !ok {
		//	// TODO find out why this cast is failing
		//	t.Errorf("Schema cast failed")
		//}
		//if schema.getRenderer() == nil {
		//	t.Errorf("%v Response Schema of type %v is missing a renderer", k, mt)
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
		//		t.Errorf("%v Response Schema of type %v is missing a renderer", k, mt)
		//	}
		//}
		//}
	}
	//Parameter
}

func TestOpenAPI_GetBasePath(t *testing.T) {
	openapi, _ := LoadOpenAPI(testAbsOpenAPI)
	wanted := filepath.Dir(testAbsOpenAPI)
	if openapi.getBasePath() != wanted {
		t.Errorf("could not get basePath %s, got %s", wanted, openapi.basePath)
	}
}

func TestOpenAPI_GetParent(t *testing.T) {
	openapi, _ := LoadOpenAPI(testAbsOpenAPI)
	p := openapi.GetParent()
	if p != nil {
		t.Errorf("the root openapi file found a parent: %v", p)
	}
}

func TestGetOpenAPI_GetChildren(t *testing.T) {
	openapi, _ := LoadOpenAPI(testAbsOpenAPI)
	paths := openapi.getChildren()
	if len(paths) == 0 {
		t.Errorf("error occured or test yaml file has no paths to get")
	}
	for k, v := range paths {
		if k == testEndpoint {
			if _, ok := v.(*PathItem); ok {
				return // test that it's a *PathItem
			}
		}
	}
	t.Errorf("could not find a valid child in openapi file")
}

func TestOpenAPI_SetChild(t *testing.T) {
	openapi, _ := LoadOpenAPI(testAbsOpenAPI)
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
	openapi, _ := LoadOpenAPI(testAbsOpenAPI)
	_ = SetRenderer(openapi, GinRenderer{})
	name := openapi.GetName()
	if name != testGinRenderedOpenAPIName {
		t.Errorf("wanted %s, got %s", testGinRenderedOpenAPIName, name)
	}
}

func TestOpenAPI_ResolveRefs(t *testing.T) {
	api := OpenAPI{
		node: node{
			basePath: filepath.Dir(testAbsOpenAPI),
		},
		Components: Components{},
		Paths:      map[string]*PathItem{},
	}
	content, _ := os.ReadFile(testAbsOpenAPI)
	_ = yaml.Unmarshal(content, &api)

	newApi, err := Traverse(&api, resolveRefs)

	if err != nil {
		t.Errorf(err.Error()) // just bubbling up is enough here
	}
	if newApi == nil {
		t.Errorf("%s resulted in a nil api tree", testOpenapiFile)
	}
}

func TestOpenAPI_SetMetadata(t *testing.T) {
	apiSpec, _ := LoadOpenAPI(testAbsOpenAPI)
	data := map[string]string{"this": "is", "some": "metadata"}
	traversable := Dig(apiSpec, testEndpoint, testMethod, testReqBody, testMediaType, testSchema, "name")

	traversable.SetMetadata(data)
	if !reflect.DeepEqual(apiSpec.metadata, data) {
		t.Fatalf("new metadata did not populate to root of tree. wanted %v, got %v", data, apiSpec.metadata)
	}
}
