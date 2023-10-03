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

	cfg := golang.NewGinServerTestConfig("../render/go/config_goginserver_test.yml", "../openapi/test/fixtures/validation_no_refs.yaml")
	err := cfg.Load()
	if err != nil {
		t.Errorf(err.Error())
	}
	cfg.OutputPath = tempFolder
	r := golang.Renderer{}
	r.SetTemplateFuncMap(golang.DefaultFuncMap())
	e := render.NewEngine(cfg)
	e.SetRenderer(&r)

	// We have to pop the first element off the path constant
	apiTree, err := openapi.LoadFromYaml(filepath.Join(strings.Split(OpenapiFile, "/")[1:]...))
	if err != nil {
		t.Errorf(err.Error())
	}

	apiTree.SetMetadata(tree.Metadata{
		"package": cfg.Template.BasePackageName,
	})
	_, err = tree.Traverse(apiTree, e.Render)
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

	cfg := golang.NewGoClientTestConfig("../render/go/config_goclient_test.yml", "../openapi/test/fixtures/validation_no_refs.yaml")
	err := cfg.Load()
	if err != nil {
		t.Errorf(err.Error())
	}
	cfg.OutputPath = tempFolder
	r := golang.Renderer{}
	r.SetTemplateFuncMap(golang.DefaultFuncMap())
	e := render.NewEngine(cfg)
	e.SetRenderer(&r)

	// We have to pop the first element off the path constant
	apiTree, err := openapi.LoadFromYaml(filepath.Join(strings.Split(OpenapiFile, "/")[1:]...))
	if err != nil {
		t.Errorf(err.Error())
	}

	apiTree.SetMetadata(tree.Metadata{
		"package": cfg.Template.BasePackageName,
	})
	_, err = tree.Traverse(apiTree, e.Render)
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
	cfg := golang.NewGoClientTestConfig("../render/go/config_goclient_test.yml", "../openapi/test/fixtures/validation_no_refs.yaml")
	if err := cfg.Load(); err != nil {
		t.Error(err)
	}
	cfg.OpenAPIFile = "test/fixtures/heavy_nesting.yaml"
	cfg.OutputPath = tempFolder
	r := golang.Renderer{}
	r.SetTemplateFuncMap(golang.DefaultFuncMap())
	e := render.NewEngine(cfg)
	e.SetRenderer(&r)

	// We have to pop the first element off the path constant
	apiTree, err := openapi.LoadFromYaml(filepath.Join(strings.Split(cfg.OpenAPIFile, "/")[1:]...))
	if err != nil {
		t.Error(err)
	}

	apiTree.SetMetadata(tree.Metadata{
		"package": cfg.Template.BasePackageName,
	})
	if _, err := tree.Traverse(apiTree, e.Render); err != nil {
		t.Error(err)
	}
	if err := filepath.Walk(tempFolder, assertFilesPresent(tempFolder, heavyNestingFilenames)); err != nil {
		t.Error(err)
	}
	file, err := os.Open(filepath.Join(tempFolder, "NestedApiExampleNestedendpointPostRequestBodyModelChildobjectGrandchildarrayItem.go"))
	if err != nil {
		t.Error(err)
	}
	want := "type NestedApiExampleNestedendpointPostRequestBodyModelChildobjectGrandchildarrayItem struct {\n\tGrandchildproperty string `json:\"grandchildProperty,omitempty\"`\n}"
	if err := fileShouldContain(file, want); err != nil {
		t.Error(err)
	}
	if err := file.Close(); err != nil {
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
	cfg := golang.NewGinServerTestConfig("../render/go/config_goginserver_test.yml", "../openapi/test/fixtures/validation_no_refs.yaml")
	if err := cfg.Load(); err != nil {
		t.Error(err)
	}
	cfg.OpenAPIFile = "test/fixtures/heavy_nesting.yaml"
	cfg.OutputPath = tempFolder
	r := golang.Renderer{}
	r.SetTemplateFuncMap(golang.DefaultFuncMap())
	e := render.NewEngine(cfg)
	e.SetRenderer(&r)

	// We have to pop the first element off the path constant
	apiTree, err := openapi.LoadFromYaml(filepath.Join(strings.Split(cfg.OpenAPIFile, "/")[1:]...))
	if err != nil {
		t.Error(err)
	}

	apiTree.SetMetadata(tree.Metadata{
		"package": cfg.Template.BasePackageName,
	})
	if _, err := tree.Traverse(apiTree, e.Render); err != nil {
		t.Error(err)
	}
	if err := filepath.Walk(tempFolder, assertFilesPresent(tempFolder, heavyNestingFilenames)); err != nil {
		t.Error(err)
	}
	file, err := os.Open(filepath.Join(tempFolder, "NestedApiExampleNestedendpointPostRequestBodyModelChildobjectGrandchildarrayItem.go"))
	if err != nil {
		t.Error(err)
	}
	want := "type NestedApiExampleNestedendpointPostRequestBodyModelChildobjectGrandchildarrayItem struct {\n\tGrandchildproperty string `json:\"grandchildProperty\"`\n}"
	if err := fileShouldContain(file, want); err != nil {
		t.Error(err)
	}
	if err := file.Close(); err != nil {
		t.Error(err)
	}
}

func TestRegression_GoClient_EmptyRequestAndResponseBodiesShouldRender(t *testing.T) {
	tempFolder, _ := os.MkdirTemp(os.TempDir(), "CanTestArtifacts")
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Errorf(err.Error())
		}
	}(tempFolder)

	cfg := golang.NewGoClientTestConfig("fixtures/regressions/empty_bodies/config_goclient_empty_bodies.yml", "fixtures/regressions/empty_bodies/empty_bodies.yml")
	err := cfg.Load()
	if err != nil {
		t.Errorf(err.Error())
	}
	cfg.OutputPath = tempFolder
	r := golang.Renderer{}
	r.SetTemplateFuncMap(golang.DefaultFuncMap())
	e := render.NewEngine(cfg)
	e.SetRenderer(&r)
	api, err := openapi.LoadFromYaml(cfg.OpenAPIFile)
	if err != nil {
		t.Error(err)
	}

	api.SetMetadata(tree.Metadata{
		"package": cfg.Template.BasePackageName,
	})
	_, err = tree.Traverse(api, e.Render)
	if err != nil {
		t.Errorf(err.Error())
	}

	// Patch file assertions
	file, err := os.Open(filepath.Join(tempFolder, "EmptyRequestAndResponseApiResourcePatch.go"))
	if err != nil {
		t.Error(err)
	}
	want1 := "type EmptyRequestAndResponseApiResourcePatchRequestBody struct{}"
	want2 := "type EmptyRequestAndResponseApiResourcePatch204Response struct{}"
	if err := fileShouldContain(file, want1, want2); err != nil {
		t.Error(err)
	}
	if err := file.Close(); err != nil {
		t.Error(err)
	}

	// Delete file assertions
	file, err = os.Open(filepath.Join(tempFolder, "EmptyRequestAndResponseApiResourceDelete.go"))
	if err != nil {
		t.Error(err)
	}
	want1 = "type EmptyRequestAndResponseApiResourceDeleteRequestBody struct{}"
	want2 = "type EmptyRequestAndResponseApiResourceDelete204Response struct{}"
	if err := fileShouldContain(file, want1, want2); err != nil {
		t.Error(err)
	}
	if err := file.Close(); err != nil {
		t.Error(err)
	}

	// Get file assertions
	file, err = os.Open(filepath.Join(tempFolder, "EmptyRequestAndResponseApiResourceGet.go"))
	if err != nil {
		t.Error(err)
	}
	want1 = "type EmptyRequestAndResponseApiResourceGetRequestBody struct{}"
	want2 = "type EmptyRequestAndResponseApiResourceGet200Response struct{}"
	if err := fileShouldContain(file, want1, want2); err != nil {
		t.Error(err)
	}
	if err := file.Close(); err != nil {
		t.Error(err)
	}

	// Post file assertions
	file, err = os.Open(filepath.Join(tempFolder, "EmptyRequestAndResponseApiResourcePost.go"))
	if err != nil {
		t.Error(err)
	}
	want1 = "type EmptyRequestAndResponseApiResourcePostRequestBody struct{}"
	want2 = "type EmptyRequestAndResponseApiResourcePost201Response struct{}"
	if err := fileShouldContain(file, want1, want2); err != nil {
		t.Error(err)
	}
	if err := file.Close(); err != nil {
		t.Error(err)
	}
}

func TestRegression_GoGin_EmptyRequestAndResponseBodiesShouldRender(t *testing.T) {
	tempFolder, _ := os.MkdirTemp(os.TempDir(), "CanTestArtifacts")
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Errorf(err.Error())
		}
	}(tempFolder)

	cfg := golang.NewGinServerTestConfig("fixtures/regressions/empty_bodies/config_goclient_empty_bodies.yml", "fixtures/regressions/empty_bodies/empty_bodies.yml")
	err := cfg.Load()
	if err != nil {
		t.Errorf(err.Error())
	}
	cfg.OutputPath = tempFolder
	r := golang.Renderer{}
	r.SetTemplateFuncMap(golang.DefaultFuncMap())
	e := render.NewEngine(cfg)
	e.SetRenderer(&r)
	api, err := openapi.LoadFromYaml(cfg.OpenAPIFile)
	if err != nil {
		t.Error(err)
	}

	api.SetMetadata(tree.Metadata{
		"package": cfg.Template.BasePackageName,
	})
	_, err = tree.Traverse(api, e.Render)
	if err != nil {
		t.Errorf(err.Error())
	}

	// Patch file assertions
	file, err := os.Open(filepath.Join(tempFolder, "EmptyRequestAndResponseApiResourcePatch.go"))
	if err != nil {
		t.Error(err)
	}
	want1 := "type EmptyRequestAndResponseApiResourcePatchRequestBody struct{}"
	want2 := "type EmptyRequestAndResponseApiResourcePatch204Response struct{}"
	if err := fileShouldContain(file, want1, want2); err != nil {
		t.Error(err)
	}
	if err := file.Close(); err != nil {
		t.Error(err)
	}

	// Delete file assertions
	file, err = os.Open(filepath.Join(tempFolder, "EmptyRequestAndResponseApiResourceDelete.go"))
	if err != nil {
		t.Error(err)
	}
	want1 = "type EmptyRequestAndResponseApiResourceDeleteRequestBody struct{}"
	want2 = "type EmptyRequestAndResponseApiResourceDelete204Response struct{}"
	if err := fileShouldContain(file, want1, want2); err != nil {
		t.Error(err)
	}
	if err := file.Close(); err != nil {
		t.Error(err)
	}

	// Get file assertions
	file, err = os.Open(filepath.Join(tempFolder, "EmptyRequestAndResponseApiResourceGet.go"))
	if err != nil {
		t.Error(err)
	}
	want1 = "type EmptyRequestAndResponseApiResourceGetRequestBody struct{}"
	want2 = "type EmptyRequestAndResponseApiResourceGet200Response struct{}"
	if err := fileShouldContain(file, want1, want2); err != nil {
		t.Error(err)
	}
	if err := file.Close(); err != nil {
		t.Error(err)
	}

	// Post file assertions
	file, err = os.Open(filepath.Join(tempFolder, "EmptyRequestAndResponseApiResourcePost.go"))
	if err != nil {
		t.Error(err)
	}
	want1 = "type EmptyRequestAndResponseApiResourcePostRequestBody struct{}"
	want2 = "type EmptyRequestAndResponseApiResourcePost201Response struct{}"
	if err := fileShouldContain(file, want1, want2); err != nil {
		t.Error(err)
	}
	if err := file.Close(); err != nil {
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
func fileShouldContain(file *os.File, needles ...string) error {
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	buf := make([]byte, fileInfo.Size())
	if _, err := io.ReadFull(file, buf); err != nil {
		return err
	}
	for _, needle := range needles {
		if !strings.Contains(string(buf), needle) {
			return fmt.Errorf("rendered file does not contain %s", needle)
		}
	}
	return nil
}
