package render_test

import (
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/render"
	golang "github.com/sasswart/gin-in-a-can/render/go"
	"github.com/sasswart/gin-in-a-can/tree"
	"os"
	"testing"
)

func Test_Render_Render(t *testing.T) {
	tempFolder, _ := os.MkdirTemp(os.TempDir(), "CanTestArtifacts")
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Errorf(err.Error())
		}
	}(tempFolder)

	// TODO test this in a language agnostic way or move to E2E testing suite
	cfg := golang.NewGinServerTestConfigStratus()
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
	spec, err := openapi.LoadAPISpec(cfg.OpenAPIFile)
	if err != nil {
		t.Errorf(err.Error())
	}
	_, err = tree.Traverse(spec, e.BuildRenderNode())
	if err != nil {
		t.Errorf(err.Error())
	}
	// TODO What dictates whether or not the schema type is rendered to file?
	// 	We're not currently rendering any 204 response schemas
}
