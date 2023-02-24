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
		*cfg.Template.Name != "go-gin" ||
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
			got := cfg.validTemplateName()
			if got != test.expected {
				t.Fail()
			}
		})
	}
}
func TestConfig_absOpenAPIPaths(t *testing.T) {
	cfg := newTestConfig()
	absOpenApi, absOpenApiErr := filepath.Abs(cfg.OpenAPIFile)
	absConfigPath, absConfigPathErr := filepath.Abs(*cfg.ConfigPath)

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
			setuperr:    absOpenApiErr, // TODO check that this causes a skip
		},
		{
			name:           "config fallback",
			openapifile:    cfg.OpenAPIFile,
			configfilepath: absConfigPath,
			setuperr:       absConfigPathErr, // TODO check that this causes a skip
		},
		{
			name:           "working dir fallback",
			openapifile:    cfg.OpenAPIFile,
			configfilepath: *cfg.ConfigPath,
			workingdir:     cfg.workingDirectory,
		},
		//{
		//	name:        "should fail",
		//},
	}
	for _, test := range tests {
		if test.setuperr != nil {
			t.Skipf("Skipping %s due to %s", test.name, test.setuperr.Error())
		}
		t.Run(test.name, func(t *testing.T) {
			cfg := Data{}
			if test.openapifile != "" {
				cfg.OpenAPIFile = test.openapifile
			}
			if test.configfilepath != "" {
				cfg.ConfigPath = &test.configfilepath
			}
			if test.workingdir != "" {
				cfg.workingDirectory = test.workingdir
			}
			cfg.GetOpenAPIFilepath()
			if cfg.absOpenAPIPath == "" {
				t.Fail()
			}
		})
	}
}
func TestConfig_absTemplateDirs(t *testing.T) {
	t.Errorf("TODO")
}
func TestConfig_absOutputFilepaths(t *testing.T) {
	t.Errorf("TODO")
}

func newTestConfig() Data {
	os.Args = []string{"can", "-configFile=config_test.yml", "-template=go-gin"}
	return Data{
		Generator:   Generator{},
		Template:    Template{},
		OpenAPIFile: "../openapi/test/fixtures/validation_no_refs.yaml",
		OutputPath:  ".",
	}
}
