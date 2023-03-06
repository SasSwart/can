package test

import (
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/render"
	golang "github.com/sasswart/gin-in-a-can/render/go"
	"github.com/sasswart/gin-in-a-can/tree"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"
)

func TestGolang_Renderer(t *testing.T) {
	tempFolder, _ := os.MkdirTemp(os.TempDir(), "CanTestArtifacts")
	defer os.RemoveAll(tempFolder)

	cfg := golang.NewTestRenderConfig()
	err := cfg.Load()
	if err != nil {
		t.Errorf(err.Error())
	}
	cfg.OutputPath = tempFolder
	e := render.Engine{}.With(&golang.Renderer{Base: &render.Base{}}, cfg)
	r := *e.GetRenderer()
	r.SetTemplateFuncMap(&template.FuncMap{
		"ToUpper": strings.ToUpper,
		"ToTitle": func(s string) string {
			caser := cases.Title(language.English)
			return caser.String(s)
		},
		"SanitiseName": r.SanitiseName,
		"SanitiseType": r.SanitiseType,
	})
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
