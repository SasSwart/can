package operation_test

import (
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/openapi/operation"
	"github.com/sasswart/gin-in-a-can/openapi/parameter"
	"github.com/sasswart/gin-in-a-can/openapi/request"
	"github.com/sasswart/gin-in-a-can/openapi/response"
	"github.com/sasswart/gin-in-a-can/test"
	"reflect"
	"testing"
)

func TestOpenAPI_Operation_SetAndGetBasePath(t *testing.T) {
	want := "test/Base/Path"
	o := operation.Operation{}
	o.SetBasePath(want)
	got := o.GetBasePath()
	if want != got {
		t.Fail()
	}
}

func TestOpenAPI_Operation_GetRef(t *testing.T) {
	apiSpec, _ := openapi.LoadFromYaml("../../" + test.OpenapiFile)
	ops := test.Dig(apiSpec, test.Endpoint)
	for _, op := range ops.GetChildren() {
		if o, ok := op.(*operation.Operation); ok {
			if o.GetRef() != "" {
				t.Errorf("%#v has an empty ref value", o)
			}
		}
	}
}

func TestOpenAPI_Operation_SetAndGetParent(t *testing.T) {
	apiSpec, _ := openapi.LoadFromYaml("../../" + test.OpenapiFile)
	ops := test.Dig(apiSpec, test.Endpoint)
	for _, o := range ops.GetChildren() {
		if o.GetParent() == nil {
			t.Errorf("operation %#v has a nil parent", o)
		}
	}
}

func TestOpenAPI_Operation_SetAndGetChildren(t *testing.T) {
	paramKey := "0"
	paramWant := &parameter.Parameter{}
	responseKey := "testKey"
	responseWant := &response.Response{}
	requestBodyWant := &request.Body{}

	o := operation.Operation{}
	o.SetChild(paramKey, paramWant)
	o.SetChild("", requestBodyWant)
	o.SetChild(responseKey, responseWant)
	childCount := 3

	paramGot := o.GetChildren()[paramKey].(*parameter.Parameter)
	if !reflect.DeepEqual(*paramGot, *paramWant) {
		t.Errorf("param fetched differs from param prepared")
	}
	requestBodyGot := o.GetChildren()[request.BodyKey].(*request.Body)
	if !reflect.DeepEqual(*requestBodyGot, *requestBodyWant) {
		t.Errorf("request body fetched differs from request body prepared")
	}
	responseGot := o.GetChildren()[responseKey].(*response.Response)
	if !reflect.DeepEqual(*responseGot, *responseWant) {
		t.Errorf("response fetched differs from response prepared")
	}
	if len(o.GetChildren()) != childCount {
		t.Errorf("we allocated %d children but received %d children when checked\n", childCount, len(o.GetChildren()))
	}
}
