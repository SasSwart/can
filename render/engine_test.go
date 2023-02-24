package render_test

import (
	"github.com/sasswart/gin-in-a-can/config"
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/openapi/operation"
	"github.com/sasswart/gin-in-a-can/openapi/path"
	"github.com/sasswart/gin-in-a-can/openapi/request"
	"github.com/sasswart/gin-in-a-can/render"
	"github.com/sasswart/gin-in-a-can/tree"
	"os"
	"testing"
)

var e *render.Engine
var toRender = buildTestSpec()
func buildTestSpec() *openapi.OpenAPI {
	root := openapi.OpenAPI{
		Node: tree.Node{Name: "openapi"},
	}

	requestBody := request.Body{
		Node: tree.Node{Name: "requestbody"},
	}

	requestBody2 :=  request.Body{
		Node: tree.Node{Name: "requestbody2"},
	}

	path := map[string]*path.Item{"/endpoint": {
		Node: tree.Node{Name: "pathitem"},
	}

	get := operation.Operation{
		Node: tree.Node{Name: "operation"},
	}
	post := &operation.Operation{
		Node: tree.Node{Name: "pathitem2"},
	}
	return &openapi
}

func resetTestRenderer(cfg config.Data) {
	e = render.Engine{}.New(render.GinRenderer{}, cfg)
}

func Test_Render_Render(t *testing.T) {
	tempFolder, err := os.MkdirTemp(os.TempDir(), "CanTestArtifacts")
	defer os.RemoveAll(tempFolder)

	cfg := newTestConfig()
	err = cfg.Load()
	if err != nil {
		t.Errorf(err.Error())
	}
	cfg.OutputPath = tempFolder
	resetTestRenderer(cfg)
	_, err = tree.Traverse(toRender, e.BuildRenderNode())
	if err != nil {
		t.Errorf(err.Error())
	}
}
func newTestConfig() config.Data {
	os.Args = []string{"can", "-configFile=../config/config_test.yml", "-template=go-gin", "-debug=true"}
	return config.Data{
		Generator:    config.Generator{},
		Template:     config.Template{},
		TemplatesDir: "../templates",
		OpenAPIFile:  "../openapi/test/fixtures/validation_no_refs.yaml",
		OutputPath:   ".",
	}
}
