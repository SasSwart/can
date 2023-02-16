package openapi

import (
	"path"
	"path/filepath"
)

// constants and other information used for unit testing.
// This file serves as a single source of truth for data drawn from in multiple places during testing

const testEndpoint = "/endpoint"
const testMethod = "post"
const testReqBody = "RequestBody"
const testEmptyParamName = "Param"
const testMediaType = "application/json"
const testOpenapiFile = "fixtures/validation.yaml"
const testSchema = "Model" // the Dig() key used to access any schema held within a MediaType
const testPattern = "^([a-zA-Z0-9])+([-_ @\\.]([a-zA-Z0-9])+)*$"
const testGinRenderedPathItemName = "ValidationFixtureEndpoint"
const testGinRenderedResponseName = "ValidationFixtureEndpointPost201Response"
const testGinRenderedMediaItemName = "ValidationFixtureEndpointPostRequestBody"
const testGinRenderedOpenAPIName = "ValidationFixture"
const Debug = false

var testAbsOpenAPI, _ = filepath.Abs(testOpenapiFile)

var testBasePath = path.Dir(testAbsOpenAPI)

// These functions are used purely for testing purposes.
// If a function finds use outside of testing it should be moved out of this file.

func Dig(node Traversable, key ...string) Traversable {
	if len(key) == 0 {
		return node
	}
	return Dig(node.getChildren()[key[0]], key[1:]...)
}
