package render_test

import (
	"github.com/sasswart/gin-in-a-can/render"
	golang "github.com/sasswart/gin-in-a-can/render/go"
	"github.com/sasswart/gin-in-a-can/test"
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
	cfg := golang.NewGinServerTestConfig("../render/go/config_goginserver_test.yml", "../openapi/test/fixtures/validation_no_refs.yaml")
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
	_, err = tree.Traverse(test.OpenAPITree(), e.BuildRenderNode())
	if err != nil {
		t.Errorf(err.Error())
	}
	// TODO What dictates whether or not the schema type is rendered to file?
	// 	We're not currently rendering any 204 response schemas
}
