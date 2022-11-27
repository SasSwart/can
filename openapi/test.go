package openapi

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"path"
	"path/filepath"
	"strings"
)

var caser = cases.Title(language.English)

// constants and other information used for unit testing.
// This file serves as a single source of truth for data drawn from in multiple places during testing

const testEndpoint = "/endpoint"
const testMethod = "post"
const testReqBody = "RequestBody"
const testEmptyParamName = "Param"
const testMediaType = "application/json"
const openapiFile = "fixtures/validation.yaml"
const testSchema = "Model" // the Dig() key used to access any schema held within a MediaType
const testPattern = "^([a-zA-Z0-9])+([-_ @\\.]([a-zA-Z0-9])+)*$"

var absOpenAPI, _ = filepath.Abs(openapiFile)
var ginRenderedPathItemName = caser.String(strings.TrimLeft(testEndpoint, "/"))
var testBasePath = path.Dir(absOpenAPI)
var testReqBodyJSON = fmt.Sprintf("%s[%s]", testReqBody, testMediaType)