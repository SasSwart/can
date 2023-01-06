package openapi

import "testing"

func TestResponse_GetName(t *testing.T) {
	openapi, _ := LoadOpenAPI(testAbsOpenAPI)
	SetRenderer(openapi, GinRenderer{})
	post201 := "201"
	response := Dig(openapi, testEndpoint, testMethod, post201)
	if response.GetName() != testGinRenderedResponseName {
		t.Errorf("wanted %s, got %s", testGinRenderedResponseName, response.GetName())
	}
}
