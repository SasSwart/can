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
const ReqBody = "RequestBody"
const MediaType = "application/json"
const OpenapiFile = "fixtures/validation.yaml"
const Schema = "Model" // the Dig() key used to access any schema held within a MediaType
const Pattern = "^([a-zA-Z0-9])+([-_ @\\.]([a-zA-Z0-9])+)*$"
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
