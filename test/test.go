package test

import (
	"fmt"
	"github.com/sasswart/gin-in-a-can/config"
	"github.com/sasswart/gin-in-a-can/openapi"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// constants and other information used for unit testing.
// This file serves as a single source of truth for data drawn from in multiple places during testing
const (
	RepositoryName           = "gin-in-a-can"
	Endpoint                 = "/endpoint"
	Method                   = "post"
	ReqBody                  = "RequestBody"
	EmptyParamName           = "Param"
	MediaType                = "application/json"
	OpenAPIFile              = "../openapi/fixtures/validation.yaml"
	Schema                   = "Model" // the Dig() key used to access any schema held within a MediaType
	Pattern                  = "^([a-zA-Z0-9])+([-_ @\\.]([a-zA-Z0-9])+)*$"
	GinRenderedPathItemName  = "EndpointValidationFixture"
	GinRenderedResponseName  = "PostEndpointValidationFixture201Response"
	GinRenderedMediaItemName = "PostEndpointValidationFixtureRequestbody"
	GinRenderedOpenAPIName   = "ValidationFixture"
)

var AbsOpenAPI, _ = filepath.Abs(OpenAPIFile)

var BasePath = path.Dir(AbsOpenAPI)

// These functions are used purely for testing purposes.
// If a function finds use outside of testing it should be moved out of this file.

func NewTestConfig() config.Config {
	wd, err := os.Getwd()
	if err != nil {
		panic("NewTestConfig() error: " + err.Error())
	}
	rootFolder := findRootFolder(wd)

	return config.Config{
		Generator: config.GeneratorConfig{
			ModuleName:        "github.com/sasswart/gin-in-a-can",
			BasePackageName:   "api",
			TemplateDirectory: rootFolder + "/templates/go-gin",
		},
		OpenAPI: openapi.Config{
			OpenAPIFile: fmt.Sprint(rootFolder + "/openapi/fixtures/testRefs/validation.yaml"),
		},
		OutputPath:       rootFolder + "/test/output",
		WorkingDirectory: wd,
		ConfigFilePath:   rootFolder + "/templates/go-gin",
	}
}

func findRootFolder(s string) string {
	if strings.HasSuffix(s, "/") {
		strings.TrimSuffix(s, "/")
	}
	words := strings.Split(s, "/")

	retArr := make([]string, 0, len(words))
	for !strings.HasSuffix(s, RepositoryName) {
		for _, w := range words {
			if w == RepositoryName {
				retArr = append(retArr, w)
				return strings.Join(retArr, "/")
			}
			retArr = append(retArr, w)
		}
	}
	return s
}
