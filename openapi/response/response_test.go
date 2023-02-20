package response_test

import (
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/openapi/test"
	"testing"
)

func TestOpenAPI_Response_GetName(t *testing.T) {
	apiSpec, _ := openapi.LoadAPISpec(test.AbsOpenAPI)
	//root.SetRenderer(apiSpec, render.GinRenderer{})
	post201 := "201"
	response := test.Dig(apiSpec, test.Endpoint, test.Method, post201)
	if response.GetName() != test.GinRenderedResponseName {
		t.Errorf("wanted %s, got %s", test.GinRenderedResponseName, response.GetName())
	}
}
