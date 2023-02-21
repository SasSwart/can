package test

import (
	"github.com/sasswart/gin-in-a-can/tree"
	"path"
	"path/filepath"
)

// constants and other information used for unit testing.
// This file serves as a single source of truth for data drawn from in multiple places during testing

const Endpoint = "/endpoint"
const Method = "post"
const ReqBody = "Body"
const MediaType = "application/json"
const OpenapiFile = "../test/fixtures/validation.yaml"
const Schema = "Model" // the Dig() key used to access any schema held within a MediaType
const Pattern = "^([a-zA-Z0-9])+([-_ @\\.]([a-zA-Z0-9])+)*$"

// These are expected validation strings before they hit the renderer
const PathName = "Validation Fixture/endpoint"
const ResponseName = "Validation Fixture /endpoint post201Response"
const MediaItemName = "Validation Fixture/endpointpostBodyapplication/json"
const OpenAPIName = "Validation Fixture"

// These are expected validation strings after they have been processed by the renderer
const GinRenderedPathItemName = "ValidationFixtureEndpoint"
const GinRenderedResponseName = "ValidationFixtureEndpointPost201Response"
const GinRenderedMediaItemName = "ValidationFixtureEndpointPostRequestBody"
const GinRenderedOpenAPIName = "ValidationFixture"

var AbsOpenAPI, _ = filepath.Abs(OpenapiFile)

var BasePath = path.Dir(AbsOpenAPI)

// These functions are used purely for testing purposes.
// If a function finds use outside of testing it should be moved out of this file.

func Dig(node tree.NodeTraverser, key ...string) tree.NodeTraverser {
	if len(key) == 0 {
		return node
	}
	return Dig(node.GetChildren()[key[0]], key[1:]...)
}
