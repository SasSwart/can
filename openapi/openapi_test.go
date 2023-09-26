package openapi_test

import (
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/openapi/media"
	"github.com/sasswart/gin-in-a-can/openapi/path"
	"github.com/sasswart/gin-in-a-can/openapi/request"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/test"
	"github.com/sasswart/gin-in-a-can/tree"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestOpenAPI_LoadOpenAPI(t *testing.T) {
	specPath := "../" + test.OpenapiFile
	apiSpec, err := openapi.LoadFromYaml(specPath)
	if err != nil {
		t.Errorf(err.Error())
	}
	if apiSpec == nil {
		t.Errorf("could not load file %s:::%s", specPath, err.Error())
	}
}

func TestOpenAPI_GetAndSetBasePath(t *testing.T) {
	want := "test/Base/Path"
	o := openapi.OpenAPI{}
	o.SetBasePath(want)
	got := o.GetBasePath()
	if got != want {
		t.Fail()
	}
}

func TestOpenAPI_GetParent(t *testing.T) {
	o := openapi.OpenAPI{}
	p := o.GetParent()
	if p != nil {
		t.Fail()
	}
}

func TestGetOpenAPI_GetAndSetChildren(t *testing.T) {
	pathName := "/path"
	want := path.Item{Node: tree.Node{}}
	o := openapi.OpenAPI{Node: tree.Node{}}
	o.SetChild(pathName, &want)
	got := o.GetChildren()[pathName].(*path.Item)
	if !reflect.DeepEqual(&want, got) {
		t.Fail()
	}
}

func TestOpenAPI_GetAndSetName(t *testing.T) {
	want := "testName"
	o := openapi.OpenAPI{}
	o.SetName(want)
	got := strings.Join(o.GetName(), "")
	if got != want {
		t.Fail()
	}
}

func TestOpenAPI_GetAndSetMetadata(t *testing.T) {
	want := tree.Metadata{"key": "value"}
	o := openapi.OpenAPI{}
	o.SetMetadata(want)
	got := o.GetMetadata()
	if !reflect.DeepEqual(want, got) {
		t.Fail()
	}
}

func TestOpenAPI_MetadataSetPoint(t *testing.T) {
	apiSpec, _ := openapi.LoadFromYaml("../" + test.OpenapiFile)
	data := tree.Metadata{"this": "is", "some": "metadata"}
	traversable := test.Dig(apiSpec, test.Endpoint, http.MethodPost, request.BodyKey, media.JSONKey, schema.PropertyKey, "name")

	traversable.SetMetadata(data)
	if !reflect.DeepEqual(apiSpec.GetMetadata(), data) {
		t.Fatalf("new metadata did not populate to root of tree. wanted %v, got %v", data, apiSpec.GetMetadata())
	}
}
