package main

import (
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/render"
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
			Generator: render.Config{
				ModuleName:           "github.com/sasswart/gin-in-a-can",
				BasePackageName:      "api",
				InvalidRequestStatus: "400",
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
