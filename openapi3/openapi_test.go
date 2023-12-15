package openapi3_test

import (
	"github.com/sasswart/gin-in-a-can/openapi3"
	"github.com/sasswart/gin-in-a-can/openapi3/media"
	"github.com/sasswart/gin-in-a-can/openapi3/path"
	"github.com/sasswart/gin-in-a-can/openapi3/request"
	"github.com/sasswart/gin-in-a-can/openapi3/schema"
	"github.com/sasswart/gin-in-a-can/test"
	"github.com/sasswart/gin-in-a-can/tree"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestOpenAPI_LoadOpenAPI(t *testing.T) {
	specPath := "../" + test.OpenapiFile
	apiSpec, err := openapi3.LoadFromYaml(specPath)
	if err != nil {
		t.Errorf(err.Error())
	}
	if apiSpec == nil {
		t.Errorf("could not load file %s:::%s", specPath, err.Error())
	}
}

func TestOpenAPI_GetAndSetBasePath(t *testing.T) {
	want := "test/Base/Path"
	o := openapi3.OpenAPI{}
	o.SetBasePath(want)
	got := o.GetBasePath()
	if got != want {
		t.Fail()
	}
}

func TestOpenAPI_GetParent(t *testing.T) {
	o := openapi3.OpenAPI{}
	p := o.GetParent()
	if p != nil {
		t.Fail()
	}
}

func TestGetOpenAPI_GetAndSetChildren(t *testing.T) {
	pathName := "/path"
	want := path.Item{Node: tree.Node{}}
	o := openapi3.OpenAPI{Node: tree.Node{}}
	o.SetChild(pathName, &want)
	got := o.GetChildren()[pathName].(*path.Item)
	if !reflect.DeepEqual(&want, got) {
		t.Fail()
	}
}

func TestOpenAPI_GetAndSetName(t *testing.T) {
	want := "testName"
	o := openapi3.OpenAPI{}
	o.SetName(want)
	got := strings.Join(o.GetName(), "")
	if got != want {
		t.Fail()
	}
}

func TestOpenAPI_GetAndSetMetadata(t *testing.T) {
	want := tree.Metadata{"key": "value"}
	o := openapi3.OpenAPI{}
	o.SetMetadata(want)
	got := o.GetMetadata()
	if !reflect.DeepEqual(want, got) {
		t.Fail()
	}
}

func TestOpenAPI_MetadataSetPoint(t *testing.T) {
	apiSpec, _ := openapi3.LoadFromYaml("../" + test.OpenapiFile)
	data := tree.Metadata{"this": "is", "some": "metadata"}
	traversable := test.Dig(apiSpec, test.Endpoint, http.MethodPost, request.BodyKey, media.JSONKey, schema.PropertyKey, "name")

	traversable.SetMetadata(data)
	if !reflect.DeepEqual(apiSpec.GetMetadata(), data) {
		t.Fatalf("new metadata did not populate to root of tree. wanted %v, got %v", data, apiSpec.GetMetadata())
	}
}
