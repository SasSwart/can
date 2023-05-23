// config_test should not test any of the config packages functions in a way that relies on other packages within can.
// This statement is being transgressed if any other local import is seen in the import block below.
package config

import (
	"os"
	"path/filepath"
	"testing"
)

// FYI - if your tests are utilizing flags from os.Args, flags are likely to be redefined through the
// config.Data{}.Load() method. This will cause test suites to panic!
func TestConfig_Load(t *testing.T) {
	absTestTemplateDir, err := filepath.Abs(filepath.Join("../", testTemplateDir))
	if err != nil {
		t.Errorf(err.Error())
	}
	cfg := newTestConfig()
	err = cfg.Load()
	if err != nil {
		t.Errorf(err.Error())
	}

	if cfg.Template.ModuleName != testModuleName ||
		cfg.Template.BasePackageName != testBasePackageName ||
		// will fail if the defined template name does not exist in the templates directory
		cfg.Template.Name != testTemplateName ||
		cfg.TemplatesDir != absTestTemplateDir ||
		cfg.OpenAPIFile != filepath.Join("../", testOpenAPIDefinition) ||
		cfg.OutputPath != testOutputDir {
		t.Fail()
	}
}

func TestConfig_validTemplateName(t *testing.T) {
	cfg := newTestConfig()
	t.Logf("using template directory " + cfg.TemplatesDir + " for testing")
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
		openapiFile    string
		configFilepath string
		workingDir     string
		expected       bool
		setupErr       error
	}{
		{
			name:        "absolute path",
			openapiFile: absOpenApi,
			setupErr:    absOpenApiErr,
		},
		{
			name:           "config fallback",
			openapiFile:    openAPIFilePath,
			configFilepath: absConfigPath,
			setupErr:       absConfigPathErr,
		},
		{
			name:           "working dir fallback",
			openapiFile:    openAPIFilePath,
			configFilepath: configFilePath,
			workingDir:     "../",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.setupErr != nil {
				t.Skipf("Skipping %s due to %s", test.name, test.setupErr.Error())
			}
			cfg := Data{}
			if test.openapiFile != "" {
				cfg.OpenAPIFile = test.openapiFile
			}
			if test.configFilepath != "" {
				ConfigFilePath = test.configFilepath
			}
			if test.workingDir != "" {
				ProcWorkingDir = test.workingDir
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

// Constants needed for sane testing. These are consciously intended to be agnostic of language and framework.
const (
	testTemplateDir       = "templates"
	testTemplateName      = "openapi-3"
	testModuleName        = "testModuleName"
	testBasePackageName   = "testBasePackageName"
	testConfigPath        = "test_config.yaml"
	testOpenAPIDefinition = "openapi/test/fixtures/validation_no_refs.yaml"
	testOutputDir         = "."
)

func newTestConfig() Data {
	ConfigFilePath = testConfigPath
	return Data{
		Template: Template{
			Name:            testTemplateName,
			BasePackageName: testBasePackageName,
			ModuleName:      testModuleName,
		},
		TemplatesDir: filepath.Join("../", testTemplateDir),
		OpenAPIFile:  filepath.Join("../", testOpenAPIDefinition),
		OutputPath:   testOutputDir,
	}

}
