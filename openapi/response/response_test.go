package response

import (
	"github.com/sasswart/gin-in-a-can/openapi/root"
	"github.com/sasswart/gin-in-a-can/openapi/test"
	"github.com/sasswart/gin-in-a-can/render"
	"testing"
)

func TestOpenAPI_Response_GetName(t *testing.T) {
	apiSpec, _ := root.LoadAPISpec(test.AbsOpenAPI)
	root.SetRenderer(apiSpec, render.GinRenderer{})
	post201 := "201"
	response := test.Dig(apiSpec, test.Endpoint, test.Method, post201)
	if response.GetName() != test.GinRenderedResponseName {
		t.Errorf("wanted %s, got %s", test.GinRenderedResponseName, response.GetName())
	}
}
