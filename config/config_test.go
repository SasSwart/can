package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestConfig_Load(t *testing.T) {
	cfg := newTestConfig()
	err := cfg.Load()
	if err != nil {
		t.Errorf(err.Error())
	}

	if cfg.Generator.ModuleName != "github.com/test/api" ||
		cfg.Generator.BasePackageName != "test" ||
		cfg.Template.Name != "go-gin" ||
		cfg.Template.Directory != "./templates/go-gin" ||
		cfg.TemplatesDir != "../templates" ||
		cfg.OpenAPIFile != "openapi/test/fixtures/validation_no_refs.yaml" ||
		cfg.OutputPath != "." {
		t.Fail()
	}
}
func TestConfig_validTemplateName(t *testing.T) {
	cfg := newTestConfig()
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "go-client",
			input:    "go-client",
			expected: true,
		},
		{
			name:     "go-gin",
			input:    "go-gin",
			expected: true,
		},
		{
			name:     "oas-3-0-0",
			input:    "oas-3-0-0",
			expected: true,
		},
		{
			name:     "should fail",
			input:    "DoesNotExistInTemplates",
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cfg.Template.Name = test.input
			got := cfg.validTemplateName()
			if got != test.expected {
				t.Fail()
			}
		})
	}
}
func TestConfig_GetOpenAPIFilepath(t *testing.T) {
	cfg := newTestConfig()
	_ = cfg.Load()
	absOpenApi, absOpenApiErr := filepath.Abs(cfg.OpenAPIFile)
	absConfigPath, absConfigPathErr := filepath.Abs(cfg.ConfigPath)

	tests := []struct {
		name           string
		openapifile    string
		configfilepath string
		workingdir     string
		expected       bool
		setuperr       error
	}{
		{
			name:        "absolute path",
			openapifile: absOpenApi,
			setuperr:    absOpenApiErr,
		},
		{
			name:           "config fallback",
			openapifile:    cfg.OpenAPIFile,
			configfilepath: absConfigPath,
			setuperr:       absConfigPathErr,
		},
		{
			name:           "working dir fallback",
			openapifile:    cfg.OpenAPIFile,
			configfilepath: cfg.ConfigPath,
			workingdir:     cfg.workingDirectory,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.setuperr != nil {
				t.Skipf("Skipping %s due to %s", test.name, test.setuperr.Error())
			}
			cfg := Data{}
			if test.openapifile != "" {
				cfg.OpenAPIFile = test.openapifile
			}
			if test.configfilepath != "" {
				cfg.ConfigPath = test.configfilepath
			}
			if test.workingdir != "" {
				cfg.workingDirectory = test.workingdir
			}
			// TODO figure out how to isolate test context from host filesystem for accurate testing
			if cfg.GetOpenAPIFilepath() == "" {
				t.Fail()
			}
		})
	}
}
func TestConfig_GetTemplateDir(t *testing.T) {
	// TODO see TestConfig_GetOpenAPIFilepath
}
func TestConfig_GetOutputFilepath(t *testing.T) {
	// TODO see TestConfig_GetOpenAPIFilepath
}

func newTestConfig() Data {
	os.Args = []string{"can", "-configFile=config_test.yml", "-template=go-gin"}
	return Data{
		Generator:    Generator{},
		Template:     Template{},
		TemplatesDir: "../templates",
		OpenAPIFile:  "../openapi/test/fixtures/validation_no_refs.yaml",
		OutputPath:   ".",
	}
}
