package config

import (
	"github.com/sasswart/gin-in-a-can/test"
	"os"
	"reflect"
	"testing"
)

func TestConfig_LoadConfig(t *testing.T) {
	tests := []struct {
		configFile     string
		expectedConfig Config
		expectedErr    bool
	}{
		{configFile: "", expectedConfig: Config{}, expectedErr: true},
		{configFile: "test_fixtures/example.yaml", expectedConfig: test.NewTestConfig()},
	}
	for i, testCase := range tests {
		os.Args = []string{"can"}
		if testCase.configFile != "" {
			os.Args = append(os.Args, "--configFile", testCase.configFile)
		}
		configData, err := LoadConfig()

		if !testCase.expectedErr && err != nil {
			t.Log("Test Case: ", i)
			t.Log("Unexpected error occurred while loading config file:", err)
			t.Fail()
		}
		if !reflect.DeepEqual(configData, testCase.expectedConfig) {
			t.Log("Test Case: ", i)
			t.Log("Loaded Config did not match expected config")
			t.Fail()
		}
	}
}
