package openapi

import "testing"

func TestOpenAPI_Dig(t *testing.T) {
	openapi, _ := LoadOpenAPI(openapiFile)
	endpoint := Dig(openapi, testEndpoint)
	if _, ok := endpoint.(*PathItem); !ok {
		t.Errorf("%#v shoud have been a %T", endpoint, &PathItem{})
	}

	method := Dig(endpoint, testMethod)
	if _, ok := method.(*Operation); !ok {
		t.Errorf("%#v shoud have been a %T", method, &Operation{})
	}

	reqBody := Dig(method, testReqBody)
	if _, ok := reqBody.(*RequestBody); !ok {
		t.Errorf("%#v shoud have been a %T", reqBody, &RequestBody{})
	}

	mediaType := Dig(reqBody, testMediaType)
	if _, ok := mediaType.(*MediaType); !ok {
		t.Errorf("%#v shoud have been a %T", mediaType, &MediaType{})
	}

	schema := Dig(mediaType, testSchema)
	if _, ok := schema.(*Schema); !ok {
		t.Errorf("%#v shoud have been a %T", schema, &Schema{})
	}
}
