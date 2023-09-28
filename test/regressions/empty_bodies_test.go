package regressions

import (
	"os"
	"testing"

	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/render"
	golang "github.com/sasswart/gin-in-a-can/render/go"
	"github.com/sasswart/gin-in-a-can/tree"
)

func TestEmptyRequestAndResponseBodiesShouldRender(t *testing.T) {
	const testYaml = `openapi: 3.0.0
info:
  title: Empty Request and Response API
paths:
  /resource:
    post:
      requestBody:
        required: true
        content:
          application/json: {}
      responses:
        '201':
          content:
            application/json: {}
    patch:
      requestBody:
        required: true
        content:
          application/json: {}
      responses:
        '204':
          content:
            application/json: {}
    get:
      responses:
        '200':
          content:
            application/json: {}
    delete:
      responses:
        '204':
          content:
            application/json: {}
`
	tempFolder, _ := os.MkdirTemp(os.TempDir(), "CanTestArtifacts")
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Errorf(err.Error())
		}
	}(tempFolder)

	cfg := golang.NewGoClientTestConfig()
	err := cfg.Load()
	if err != nil {
		t.Errorf(err.Error())
	}
	cfg.OutputPath = tempFolder
	e := render.Engine{}.With(&golang.Renderer{Base: &render.Base{}}, cfg)
	r := *e.GetRenderer()
	r.SetTemplateFuncMap(nil)
	if r.GetTemplateFuncMap() == nil {
		t.Errorf("TemplateFuncMap should NOT be nil")
	}
	api := openapi.NewBaseOpenApi()
	if err := api.Load([]byte(testYaml)); err != nil {
		t.Error(err)
	}

	api.SetMetadata(tree.Metadata{
		"package": cfg.Template.BasePackageName,
	})
	_, err = tree.Traverse(&api, e.BuildRenderNode())
	if err != nil {
		t.Errorf(err.Error())
	}
}
