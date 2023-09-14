package test

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/render"
	golang "github.com/sasswart/gin-in-a-can/render/go"
	"github.com/sasswart/gin-in-a-can/tree"
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

var heavyNestingFilenames = [12]string{"NestedApiExample.go",
	"NestedApiExampleNestedendpoint.go",
	"NestedApiExampleNestedendpointPost.go",
	"NestedApiExampleNestedendpointPost201ResponseModel.go",
	"NestedApiExampleNestedendpointPost201ResponseModelNestedresource.go",
	"NestedApiExampleNestedendpointPost201ResponseModelNestedresourceChildobject.go",
	"NestedApiExampleNestedendpointPost201ResponseModelNestedresourceChildobjectGrandchildarray.go",
	"NestedApiExampleNestedendpointPost201ResponseModelNestedresourceChildobjectGrandchildarrayItem.go",
	"NestedApiExampleNestedendpointPostRequestBodyModel.go",
	"NestedApiExampleNestedendpointPostRequestBodyModelChildobject.go",
	"NestedApiExampleNestedendpointPostRequestBodyModelChildobjectGrandchildarray.go",
	"NestedApiExampleNestedendpointPostRequestBodyModelChildobjectGrandchildarrayItem.go",
}

func TestGolang_GoClient_Renderer_HeavyNesting(t *testing.T) {
	tempFolder, _ := os.MkdirTemp(os.TempDir(), "CanTestArtifacts")
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Errorf(err.Error())
		}
	}(tempFolder)
	cfg := golang.NewGoClientTestConfig()
	if err := cfg.Load(); err != nil {
		t.Error(err)
	}
	cfg.OpenAPIFile = "test/fixtures/heavy_nesting.yaml"
	cfg.OutputPath = tempFolder
	e := render.Engine{}.With(&golang.Renderer{Base: &render.Base{}}, cfg)
	r := *e.GetRenderer()
	r.SetTemplateFuncMap(nil)
	if r.GetTemplateFuncMap() == nil {
		t.Errorf("TemplateFuncMap should NOT be nil")
	}

	// We have to pop the first element off the path constant
	apiTree, err := openapi.LoadAPISpec(filepath.Join(strings.Split(cfg.OpenAPIFile, "/")[1:]...))
	if err != nil {
		t.Error(err)
	}

	apiTree.SetMetadata(tree.Metadata{
		"package": cfg.Template.BasePackageName,
	})
	if _, err := tree.Traverse(apiTree, e.BuildRenderNode()); err != nil {
		t.Error(err)
	}
	if err := filepath.Walk(tempFolder, assertFilesPresent(tempFolder, heavyNestingFilenames)); err != nil {
		t.Error(err)
	}
	file, err := os.Open(filepath.Join(tempFolder, "NestedApiExampleNestedendpointPostRequestBodyModelChildobjectGrandchildarrayItem.go"))
	if err != nil {
		t.Error(err)
	}
	if err := fileShouldContain(file, "type NestedApiExampleNestedendpointPostRequestBodyModelChildobjectGrandchildarrayItem struct {\n        Grandchildproperty  string `json:\"grandchildProperty,omitempty\"`\n}"); err != nil {
		t.Error(err)
	}
}

func TestGolang_GoGin_Renderer_HeavyNesting(t *testing.T) {
	tempFolder, _ := os.MkdirTemp(os.TempDir(), "CanTestArtifacts")
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Errorf(err.Error())
		}
	}(tempFolder)
	cfg := golang.NewGinServerTestConfig()
	if err := cfg.Load(); err != nil {
		t.Error(err)
	}
	cfg.OpenAPIFile = "test/fixtures/heavy_nesting.yaml"
	cfg.OutputPath = tempFolder
	e := render.Engine{}.With(&golang.Renderer{Base: &render.Base{}}, cfg)
	r := *e.GetRenderer()
	r.SetTemplateFuncMap(nil)
	if r.GetTemplateFuncMap() == nil {
		t.Errorf("TemplateFuncMap should NOT be nil")
	}

	// We have to pop the first element off the path constant
	apiTree, err := openapi.LoadAPISpec(filepath.Join(strings.Split(cfg.OpenAPIFile, "/")[1:]...))
	if err != nil {
		t.Error(err)
	}

	apiTree.SetMetadata(tree.Metadata{
		"package": cfg.Template.BasePackageName,
	})
	if _, err := tree.Traverse(apiTree, e.BuildRenderNode()); err != nil {
		t.Error(err)
	}
	if err := filepath.Walk(tempFolder, assertFilesPresent(tempFolder, heavyNestingFilenames)); err != nil {
		t.Error(err)
	}
	file, err := os.Open(filepath.Join(tempFolder, "NestedApiExampleNestedendpointPostRequestBodyModelChildobjectGrandchildarrayItem.go"))
	if err != nil {
		t.Error(err)
	}
	if err := fileShouldContain(file, "type NestedApiExampleNestedendpointPostRequestBodyModelChildobjectGrandchildarrayItem struct {\n        Grandchildproperty string `json:\"grandchildProperty\"`\n}"); err != nil {
		t.Error(err)
	}
}

func assertFilesPresent(parentDirectoryPath string, haystack [12]string) func(currentFilePath string, info os.FileInfo, err error) error {
	return func(currentFilePath string, info os.FileInfo, err error) error {
		// handle error, return if present
		if err != nil {
			return err
		}
		// Skip the directory itself
		if currentFilePath == parentDirectoryPath {
			return nil
		}
		return findFile(filepath.Base(currentFilePath), haystack)
	}
}

func findFile(needle string, haystack [12]string) error {
	found := false
	for _, file := range haystack {
		if needle == file {
			found = true
		}
	}
	if !found {
		return fmt.Errorf("could not find %s", needle)
	}
	return nil
}
func fileShouldContain(file *os.File, needle string) error {
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	buf := make([]byte, fileInfo.Size())
	if _, err := io.ReadFull(file, buf); err != nil {
		return err
	}
	if !strings.Contains(string(buf), needle) {
		return fmt.Errorf("rendered file does not contain %s", needle)
	}
	return nil
}
