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
var toRender = &openapi.OpenAPI{
	Node: tree.Node{Name: "openapi"},
	Paths: map[string]*path.Item{"/endpoint": {
		Node: tree.Node{Name: "pathitem"},
		Get: &operation.Operation{
			Node: tree.Node{Name: "operation"},
			RequestBody: request.Body{
				Node: tree.Node{Name: "requestbody"},
			},
		},
		Post: &operation.Operation{
			Node: tree.Node{Name: "pathitem2"},
			RequestBody: request.Body{
				Node: tree.Node{Name: "requestbody2"},
			},
		},
	}},
}

func resetTestRenderer(t *testing.T, cfg config.Data) {
	err := cfg.Load()
	if err != nil {
		t.Errorf(err.Error())
	}
	e = render.Engine{}.New(render.GinRenderer{}, cfg)
}

func Test_Render_Render(t *testing.T) {
	cfg := newTestConfig()
	resetTestRenderer(t, cfg)
	_, err := e.Render(toRender, cfg.GetTemplateDir())
	if err != nil {
		t.Errorf(err.Error())
	}
}
func newTestConfig() config.Data {
	os.Args = []string{"can", "-configFile=../config/config_test.yml", "-template=go-gin"}
	return config.Data{
		Generator:    config.Generator{},
		Template:     config.Template{},
		TemplatesDir: "../templates",
		OpenAPIFile:  "../openapi/test/fixtures/validation_no_refs.yaml",
		OutputPath:   ".",
	}
}
