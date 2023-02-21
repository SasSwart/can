package render_test

import (
	"github.com/sasswart/gin-in-a-can/config"
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/openapi/operation"
	"github.com/sasswart/gin-in-a-can/openapi/path"
	"github.com/sasswart/gin-in-a-can/openapi/request"
	"github.com/sasswart/gin-in-a-can/render"
	"github.com/sasswart/gin-in-a-can/tree"
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

func resetTestRenderer() {
	cfg := config.Data{}
	err := cfg.Load()
	if err != nil {
		panic(err.Error())
	}
	e = render.Engine{}.New(render.GinRenderer{}, cfg)
}

func Test_Render_Render(t *testing.T) {
	resetTestRenderer()
	templatePath := ""
	_, err := e.Render(toRender, templatePath)
	if err != nil {
		t.Errorf(err.Error())
	}
}
