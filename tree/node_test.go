package tree_test

import (
	"github.com/sasswart/gin-in-a-can/openapi/media"
	"github.com/sasswart/gin-in-a-can/openapi/operation"
	"github.com/sasswart/gin-in-a-can/openapi/path"
	"github.com/sasswart/gin-in-a-can/openapi/request"
	"github.com/sasswart/gin-in-a-can/openapi/root"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/openapi/test"
	"testing"
)

func TestOpenAPI_Dig(t *testing.T) {
	apiSpec, _ := root.LoadAPISpec(test.AbsOpenAPI)
	endpoint := test.Dig(apiSpec, test.Endpoint)
	// TODO check for identity, not just type
	if _, ok := endpoint.(*path.Item); !ok {
		t.Errorf("%#v should have been a %T", endpoint, &path.Item{})
	}

	method := test.Dig(endpoint, test.Method)
	if _, ok := method.(*operation.Operation); !ok {
		t.Errorf("%#v should have been a %T", method, &operation.Operation{})
	}

	reqBody := test.Dig(method, test.ReqBody)
	if _, ok := reqBody.(*request.Body); !ok {
		t.Errorf("%#v should have been a %T", reqBody, &request.Body{})
	}

	mediaType := test.Dig(reqBody, test.MediaType)
	if _, ok := mediaType.(*media.Type); !ok {
		t.Errorf("%#v should have been a %T", mediaType, &media.Type{})
	}

	s := test.Dig(mediaType, test.Schema)
	if _, ok := s.(*schema.Schema); !ok {
		t.Errorf("%#v should have been a %T", s, &schema.Schema{})
	}
}

func TestOpenAPI_Node_GetMetadata(t *testing.T) {
	t.Errorf("TODO")
}

func TestOpenAPI_Node_GetOutputFile(t *testing.T) {
	t.Errorf("TODO")
}
