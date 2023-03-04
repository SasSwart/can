package config

import (
	"os"
	"path/filepath"
	"testing"
)

// FYI - if your tests are utilizing flags from os.Args, flags are likely to be redefined through the
// config.Data{}.Load() method. This will cause test suites to panic!
func TestConfig_Load(t *testing.T) {
	cfg := newTestConfig()
	err := cfg.Load()
	if err != nil {
		t.Errorf(err.Error())
	}

	if cfg.Generator.ModuleName != "github.com/test/api" ||
		cfg.Generator.BasePackageName != "test" ||
		cfg.Template.Name != "go-gin" ||
		cfg.Template.Directory != filepath.Clean("../templates/go-gin") ||
		cfg.TemplatesDir != filepath.Clean("../templates") ||
		cfg.OpenAPIFile != filepath.Clean("openapi/test/fixtures/validation_no_refs.yaml") ||
		cfg.OutputPath != "." {
		t.Fail()
	}
}

func TestConfig_validTemplateName(t *testing.T) {
	cfg := newTestConfig()
	var err error
	ProcWorkingDir, err = os.Getwd()
	if err != nil {
		t.Fail()
	}
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
			name:     "openapi-3",
			input:    "openapi-3",
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
	err := cfg.Load()
	if err != nil {
		t.Errorf(err.Error())
	}
	openAPIFilePath := "openapi/test/fixtures/validation_no_refs.yaml"
	configFilePath := "config_test.yml"
	absOpenApi, absOpenApiErr := filepath.Abs(openAPIFilePath)
	absConfigPath, absConfigPathErr := filepath.Abs(configFilePath)

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
			openapifile:    openAPIFilePath,
			configfilepath: absConfigPath,
			setuperr:       absConfigPathErr,
		},
		{
			name:           "working dir fallback",
			openapifile:    openAPIFilePath,
			configfilepath: configFilePath,
			workingdir:     "../",
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
				ConfigFilePath = test.configfilepath
			}
			if test.workingdir != "" {
				ProcWorkingDir = test.workingdir
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
	ConfigFilePath = "config_test.yml"
	return Data{
		Generator: Generator{},
		Template: Template{
			Name: "go-gin",
		},
		TemplatesDir: "../templates",
		OpenAPIFile:  "../openapi/test/fixtures/validation_no_refs.yaml",
		OutputPath:   ".",
	}

}
