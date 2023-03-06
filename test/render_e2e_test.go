package test

import (
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/render"
	golang "github.com/sasswart/gin-in-a-can/render/go"
	"github.com/sasswart/gin-in-a-can/tree"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGolang_GinServer_Renderer(t *testing.T) {
	tempFolder, _ := os.MkdirTemp(os.TempDir(), "CanTestArtifacts")
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Errorf(err.Error())
		}
	}(tempFolder)

	cfg := golang.NewGinServerTestConfig()
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

	// We have to pop the first element off the path constant
	apiTree, err := openapi.LoadAPISpec(filepath.Join(strings.Split(OpenapiFile, "/")[1:]...))
	if err != nil {
		t.Errorf(err.Error())
	}

	apiTree.SetMetadata(tree.Metadata{
		"package": cfg.Template.BasePackageName,
	})
	_, err = tree.Traverse(apiTree, e.BuildRenderNode())
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestGolang_GoClient_Renderer(t *testing.T) {
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

	// We have to pop the first element off the path constant
	apiTree, err := openapi.LoadAPISpec(filepath.Join(strings.Split(OpenapiFile, "/")[1:]...))
	if err != nil {
		t.Errorf(err.Error())
	}

	apiTree.SetMetadata(tree.Metadata{
		"package": cfg.Template.BasePackageName,
	})
	_, err = tree.Traverse(apiTree, e.BuildRenderNode())
	if err != nil {
		t.Errorf(err.Error())
	}
}
