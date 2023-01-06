package render

import (
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/test"
	"reflect"
	"testing"
)

func TestRender_Render(t *testing.T) {
	var (
		got      []byte
		expected []byte
		err      error

		schema           openapi.Traversable
		configData       = test.NewTestConfig()
		templateFileName = "schema.tmpl"
	)

	openAPI, _ := openapi.LoadOpenAPI(test.AbsOpenAPI)
	openapi.SetRenderer(openAPI, GinRenderer{})
	schema = openapi.Dig(openAPI, test.Endpoint, test.Method, test.ReqBody, test.MediaType, test.Schema)

	got, err = Render(configData, schema, templateFileName)
	if err != nil {
		t.Errorf("Test error: %s\n", err.Error())
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %s, expected %s\n", got, expected)
	}
}
