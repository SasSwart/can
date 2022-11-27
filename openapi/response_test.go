package openapi

import "testing"

func TestResponse_GetName(t *testing.T) {
	openapi, _ := LoadOpenAPI(absOpenAPI)
	post201 := "201"
	response := Dig(openapi, testEndpoint, testMethod, post201)
	if response.GetName() != post201 {
		t.Errorf("wanted %s, got %s", post201, response.GetName())
	}
}
