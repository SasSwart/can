package main

import (
	"github.com/sasswart/gin-in-a-can/config"
	"github.com/sasswart/gin-in-a-can/openapi"
	"os"
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	wd, _ := os.Getwd()
	tests := []struct {
		configFile     string
		expectedConfig Config
		expectedErr    bool
	}{
		{configFile: "", expectedConfig: Config{}, expectedErr: true},
		{configFile: "test_fixtures/example.yaml", expectedConfig: Config{
			Generator: config.Config{
				ModuleName:           "github.com/sasswart/gin-in-a-can",
				BasePackageName:      "api",
				InvalidRequestStatus: "400",
				TemplateDirectory:    "",
			},
			OpenAPI: openapi.Config{
				OpenAPIFile: "./docs/openapi.yml",
			},
			OutputPath:       ".",
			WorkingDirectory: wd,
			ConfigFilePath:   "test_fixtures/example.yaml",
			Language:         "go-gin",
		}},
	}
	for i, test := range tests {
		os.Args = []string{"can"}
		if test.configFile != "" {
			os.Args = append(os.Args, "--configFile", test.configFile)
		}
		config, err := loadConfig()

		if !test.expectedErr && err != nil {
			t.Log("Test Case: ", i)
			t.Log("Unexpected error occurred while loading config file:", err)
			t.Fail()
		}
		if !reflect.DeepEqual(config, test.expectedConfig) {
			t.Log("Test Case: ", i)
			t.Log("Loaded Config did not match expected config")
			t.Fail()
		}
	}
}

func TestBuildRenderNode(t *testing.T) {
	testFunc := buildRenderNode(Config{
		Generator: config.Config{
			ModuleName:           "Test",
			BasePackageName:      "Test",
			InvalidRequestStatus: "500",
			TemplateDirectory:    "",
		},
		OpenAPI: openapi.Config{
			OpenAPIFile: "./openapi/fixtures/petstore_openapi.yml",
		},
		OutputPath:        "",
		TemplateDirectory: "",
		WorkingDirectory:  "",
		ConfigFilePath:    "",
		Language:          "go-gin",
	})

	schema := openapi.Schema{
		Description: "",
		Type:        "",
		Properties: map[string]*openapi.Schema{
			"required_field": {},
			"optional_field": {},
		},
		Items:                nil,
		Ref:                  "",
		AdditionalProperties: false,
		MinLength:            0,
		MaxLength:            0,
		Pattern:              "",
		Format:               "",
		Required: []string{
			"required_field",
		},
	}

	//  func(key string, parent Traversable, child Traversable) (Traversable, error)
	_, err := testFunc("", nil, &schema)
	if err != nil {
		t.Errorf("Traversal function error: %v", err)
		t.Fail()
	}
}
