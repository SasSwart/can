package tree

import (
	"github.com/sasswart/gin-in-a-can/openapi/media_type"
	"github.com/sasswart/gin-in-a-can/openapi/operation"
	"github.com/sasswart/gin-in-a-can/openapi/path"
	"github.com/sasswart/gin-in-a-can/openapi/request_body"
	openapi2 "github.com/sasswart/gin-in-a-can/openapi/root"
	schema2 "github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/openapi/test"
	"testing"
)

func TestOpenAPI_Dig(t *testing.T) {
	openapi, _ := openapi2.LoadAPISpec(testAbsOpenAPI)
	endpoint := test.Dig(openapi, testEndpoint)
	// TODO check for identity, not just type
	if _, ok := endpoint.(*path.Item); !ok {
		t.Errorf("%#v should have been a %T", endpoint, &path.Item{})
	}

	method := test.Dig(endpoint, testMethod)
	if _, ok := method.(*operation.Operation); !ok {
		t.Errorf("%#v should have been a %T", method, &operation.Operation{})
	}

	reqBody := test.Dig(method, testReqBody)
	if _, ok := reqBody.(*request_body.RequestBody); !ok {
		t.Errorf("%#v should have been a %T", reqBody, &request_body.RequestBody{})
	}

	mediaType := test.Dig(reqBody, testMediaType)
	if _, ok := mediaType.(*media_type.MediaType); !ok {
		t.Errorf("%#v should have been a %T", mediaType, &media_type.MediaType{})
	}

	schema := test.Dig(mediaType, testSchema)
	if _, ok := schema.(*schema2.Schema); !ok {
		t.Errorf("%#v should have been a %T", schema, &schema2.Schema{})
	}
}

func TestOpenAPI_Node_GetMetadata(t *testing.T) {
	t.Errorf("TODO")
}

func TestOpenAPI_Node_GetOutputFile(t *testing.T) {
	t.Errorf("TODO")
}
