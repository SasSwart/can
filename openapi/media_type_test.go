package openapi

import "testing"

func TestMediaType_GetName(t *testing.T) {
	openapi, _ := LoadOpenAPI(openapiFile)
	_ = openapi.getChildren()["/endpoint"].(*PathItem).getChildren()
}

func TestMediaType_GetChildren(t *testing.T) {

}

func TestMediaType_SetChild(t *testing.T) {

}
