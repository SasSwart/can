package render_test

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/sasswart/gin-in-a-can/config"
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/openapi/operation"
	"github.com/sasswart/gin-in-a-can/openapi/parameter"
	"github.com/sasswart/gin-in-a-can/openapi/path"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/render"
	golang "github.com/sasswart/gin-in-a-can/render/go"
	"github.com/sasswart/gin-in-a-can/test"
	"github.com/sasswart/gin-in-a-can/tree"
	"os"
	"path/filepath"
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
	r := &golang.Renderer{}
	r.SetTemplateFuncMap(golang.DefaultFuncMap())
	e := render.NewEngine(cfg)
	e.SetRenderer(r)
	_, err = tree.Traverse(test.OpenAPITree(), e.BuildRenderNode())
	if err != nil {
		t.Errorf(err.Error())
	}
	// TODO What dictates whether or not the schema type is rendered to file?
	// 	We're not currently rendering any 204 response schemas
}

func TestEngineGetAndSetRenderer(t *testing.T) {
	mockRenderer := render.MockRenderer{}
	engine := render.Engine{}
	engine.SetRenderer(mockRenderer)

	if engine.GetRenderer() != mockRenderer {
		t.Errorf("Expected renderer to be %v, but got %v", mockRenderer, engine.GetRenderer())
	}
}

func TestRender(t *testing.T) {
	mockRenderer := render.MockRenderer{}
	mockConfig := config.Data{OutputPath: "../templates"}
	testEngine := render.NewEngine(mockConfig)
	testEngine.SetRenderer(mockRenderer)

	mockNode := &schema.Schema{Type: "object"}
	mockParentNode := &parameter.Parameter{}
	outputChan := make(chan []byte)
	defer close(outputChan)
	// this function builder is being used to indirectly call render.render() and catch it's output
	renderFunc := testEngine.BuildTestRenderNode(outputChan)
	go func() {
		_, err := renderFunc("", mockParentNode, mockNode)
		if err != nil {
			t.Error(err)
		}
	}()
	output := <-outputChan

	expectedOutput := []byte("rendered_output")
	if !bytes.Equal(output, expectedOutput) {
		t.Errorf("Expected output to be %v, but got %v", expectedOutput, output)
	}
}

func TestWriteToDisk(t *testing.T) {
	tempDir := t.TempDir()
	contents := []byte("test_data")
	outPath := filepath.Join(tempDir, "test_file.txt")

	if err := render.WriteToDisk(contents, outPath); err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	_, err := os.Stat(outPath)
	if errors.Is(err, os.ErrNotExist) {
		t.Errorf("Expected file to be written, but it doesn't exist")
	}
}
func TestGetTemplateFilename(t *testing.T) {
	testCases := []struct {
		node     tree.NodeTraverser
		expected string
	}{
		{&openapi.OpenAPI{}, "openapi.tmpl"},
		{&path.Item{}, "path_item.tmpl"},
		{&operation.Operation{}, "operation.tmpl"},
		{&schema.Schema{}, "schema.tmpl"},
		{nil, ""},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Node type: %T", tc.node), func(t *testing.T) {
			result := render.GetTemplateFilename(tc.node)
			if result != tc.expected {
				t.Errorf("Expected template filename to be %s, but got %s", tc.expected, result)
			}
		})
	}
}
