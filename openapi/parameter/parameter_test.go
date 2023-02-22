package parameter_test

import (
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/openapi/parameter"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/openapi/test"
	"github.com/sasswart/gin-in-a-can/tree"
	"reflect"
	"testing"
)

func TestOpenAPILoadsParameterValidation(t *testing.T) {
	apiSpec, _ := openapi.LoadAPISpec(test.AbsOpenAPI)

	IDParam := apiSpec.Paths["/endpoint"].Post.Parameters[0]
	if IDParam.Required != true {
		t.Errorf("Expected id parameter->required to be boolean value true, not %v", IDParam.Required)
	}
	if IDParam.Schema.Type != "string" {
		t.Errorf("Expected id parameter->schema->type to contain `string`, not %v", IDParam.Schema.Type)
	}
	if IDParam.Schema.Format != "uuid" {
		t.Errorf("Expected id parameter->schema->format to contain `uuid`, not %v", IDParam.Schema.Format)
	}
}
func TestOpenAPI_Parameter_GetRef(t *testing.T) {
	want := "testRef"
	p := parameter.Parameter{Ref: want}
	got := p.GetRef()
	if want != got {
		t.Fail()
	}
}

func TestOpenAPI_Parameter_SetAndGetName(t *testing.T) {
	p := parameter.Parameter{
		Node: tree.Node{
			Name: "This should be overwritten by SetName()",
		},
	}
	want := "testName"
	p.SetName(want)
	want += "Parameter" // This is done in GetName()
	got := p.GetName()
	if want != got {
		t.Errorf("wanted %s but got %s", want, got)
	}
}
func TestOpenAPI_Parameter_SetAndGetChildren(t *testing.T) {
	p := parameter.Parameter{
		Node: tree.Node{
			Name: "This should be overwritten by SetName()",
		},
	}
	want := schema.Schema{Node: tree.Node{Name: "testName"}}
	p.SetChild("", &want)
	got := p.GetChildren()[schema.Key]
	if !reflect.DeepEqual(*got.(*schema.Schema), want) {
		t.Fail()
	}
}
