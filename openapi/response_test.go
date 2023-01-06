package openapi

import (
	"github.com/sasswart/gin-in-a-can/render"
	"github.com/sasswart/gin-in-a-can/test"
	"testing"
)

func TestResponse_GetName(t *testing.T) {
	openapi, _ := LoadOpenAPI(test.AbsOpenAPI)
	SetRenderer(openapi, render.GinRenderer{})
	post201 := "201"
	response := Dig(openapi, test.Endpoint, test.Method, post201)
	if response.GetName() != test.GinRenderedResponseName {
		t.Errorf("wanted %s, got %s", test.GinRenderedResponseName, response.GetName())
	}
}
