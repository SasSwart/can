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
		expectedConfig config.Config
		expectedErr    bool
	}{
		{configFile: "", expectedConfig: config.Config{}, expectedErr: true},
		{configFile: "test_fixtures/example.yaml", expectedConfig: config.Config{
			Generator: config.GeneratorConfig{
				ModuleName:      "github.com/sasswart/gin-in-a-can",
				BasePackageName: "api",
			},
			OpenAPI: openapi.Config{
				OpenAPIFile: "./docs/openapi.yml",
			},
			OutputPath:       ".",
			WorkingDirectory: wd,
			ConfigFilePath:   "test_fixtures/example.yaml",
		}},
	}
	for i, test := range tests {
		os.Args = []string{"can"}
		if test.configFile != "" {
			os.Args = append(os.Args, "--configFile", test.configFile)
		}
		configData, err := loadConfig()

		if !test.expectedErr && err != nil {
			t.Log("Test Case: ", i)
			t.Log("Unexpected error occurred while loading config file:", err)
			t.Fail()
		}
		if !reflect.DeepEqual(configData, test.expectedConfig) {
			t.Log("Test Case: ", i)
			t.Log("Loaded Config did not match expected config")
			t.Fail()
		}
	}
}
