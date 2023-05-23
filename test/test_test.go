package test

import (
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/openapi/media"
	"github.com/sasswart/gin-in-a-can/openapi/operation"
	"github.com/sasswart/gin-in-a-can/openapi/path"
	"github.com/sasswart/gin-in-a-can/openapi/request"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"net/http"
	"testing"
)

func TestOpenAPI_Dig(t *testing.T) {
	apiSpec, _ := openapi.LoadAPISpec("../" + OpenapiFile)
	endpoint := Dig(apiSpec, Endpoint)
	// TODO check for identity, not just type
	if _, ok := endpoint.(*path.Item); !ok {
		t.Errorf("%#v should have been a %T", endpoint, &path.Item{})
	}

	method := Dig(endpoint, http.MethodPost)
	if _, ok := method.(*operation.Operation); !ok {
		t.Errorf("%#v should have been a %T", method, &operation.Operation{})
	}

	reqBody := Dig(method, request.BodyKey)
	if _, ok := reqBody.(*request.Body); !ok {
		t.Errorf("%#v should have been a %T", reqBody, &request.Body{})
	}

	mediaType := Dig(reqBody, media.JSONKey)
	if _, ok := mediaType.(*media.Type); !ok {
		t.Errorf("%#v should have been a %T", mediaType, &media.Type{})
	}

	s := Dig(mediaType, schema.PropertyKey)
	if _, ok := s.(*schema.Schema); !ok {
		t.Errorf("%#v should have been a %T", s, &schema.Schema{})
	}
}
