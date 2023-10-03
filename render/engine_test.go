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
	"text/template"
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
	_, err = tree.Traverse(test.OpenAPITree(), e.Render)
	if err != nil {
		t.Errorf(err.Error())
	}
	// TODO What dictates whether or not the schema type is rendered to file?
	// 	We're not currently rendering any 204 response schemas
}

func TestParseTemplate(t *testing.T) {
	tests := []struct {
		name              string
		templateFilename  string
		templateDirectory string
		funcMap           *template.FuncMap
		expectedErr       bool
	}{
		{
			name:              "Valid Input",
			templateFilename:  "openapi.tmpl",
			templateDirectory: "../templates/go-client",
			funcMap:           golang.DefaultFuncMap(),
			expectedErr:       false,
		},
		{
			name:              "Invalid Directory",
			templateFilename:  "test.tmpl",
			templateDirectory: "nonexistent_directory",
			funcMap:           golang.DefaultFuncMap(),
			expectedErr:       true,
		},
		{
			name:              "Empty FuncMap",
			templateFilename:  "test.tmpl",
			templateDirectory: "../templates/go-client",
			funcMap:           &template.FuncMap{},
			expectedErr:       true,
		},
		{
			name:              "Glob Error",
			templateFilename:  "test.tmpl",
			templateDirectory: "invalid_directory",
			funcMap:           golang.DefaultFuncMap(),
			expectedErr:       true,
		},
	}

	// Create a Renderer instance

	cfg := golang.NewGoClientTestConfig("../render/go/config_goclient_test.yml", "../openapi/test/fixtures/validation_no_refs.yaml")
	err := cfg.Load()
	if err != nil {
		t.Errorf(err.Error())
	}
	e := render.NewEngine(cfg)
	e.SetRenderer(&golang.Renderer{})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e.GetRenderer().SetTemplateFuncMap(test.funcMap)
			_, err := e.ParseTemplate(test.templateFilename, test.templateDirectory)
			if (err != nil) != test.expectedErr {
				t.Errorf("Expected error: %v, but got: %v", test.expectedErr, err)
			}
		})
	}
}
func TestEngineGetAndSetRenderer(t *testing.T) {
	mockRenderer := &render.MockRenderer{}
	engine := render.Engine{}
	engine.SetRenderer(mockRenderer)

	if engine.GetRenderer() != mockRenderer {
		t.Errorf("Expected renderer to be %v, but got %v", mockRenderer, engine.GetRenderer())
	}
}

func TestRender(t *testing.T) {
	mockRenderer := &render.MockRenderer{}
	mockConfig := config.Data{OutputPath: ".", TemplatesDir: "../../templates/go-client"}
	testEngine := render.NewEngine(mockConfig)
	testEngine.SetRenderer(mockRenderer)
	mockRenderer.SetTemplateFuncMap(golang.DefaultFuncMap())

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
