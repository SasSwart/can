package generator

import (
	"fmt"
	"strings"

	"github.gom/sasswart/gin-in-a-can/openapi"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func NewServerInterface(apiSpec openapi.OpenAPI) ServerInterface {
	serverInterface := ServerInterface{}
	for pathName, pathItem := range apiSpec.Paths {
		for method, operation := range pathItem.Operations() {
			serverInterface = append(serverInterface, NewRoute(
				pathName,
				method,
				append(pathItem.Parameters, operation.Parameters...),
				operation.RequestBody,
				operation.Responses,
			))
		}
	}
	return serverInterface
}

type ServerInterface []Route

type Route struct {
	Name        string
	Path        string
	Method      string
	Parameters  []openapi.Parameter
	Responses   map[string]Schema
	RequestBody Schema
}

type Schema struct {
	Properties map[string]Schema
	Type       string
}

func NewRequestBody(body openapi.RequestBody) Schema {
	return NewSchema(body.Content["application/json"].Schema)
}

func NewSchema(schema openapi.Schema) Schema {
	var schemaType string
	switch true {
	case schema.Type == "object":
		schemaType = "struct"

		schemas := make(map[string]Schema)
		for s, schema := range schema.Properties {
			schemas[s] = NewSchema(schema)
		}
		return Schema{
			Properties: schemas,
			Type:       schemaType,
		}
		break
	case schema.Type == "array":
		itemSchema := NewSchema(*schema.Items)
		return Schema{
			Type: "[]" + itemSchema.Type,
		}
		break
	case schema.Type == "boolean":
		return Schema{
			Type: "bool",
		}
	}
	return Schema{
		Type: schema.Type,
	}
}

func NewRoute(pathName, method string, parameters []openapi.Parameter, requestBody openapi.RequestBody, responses map[string]openapi.Response) Route {
	caser := cases.Title(language.English)
	transformedResponses := make(map[string]Schema)
	for r, response := range responses {
		transformedResponses[r] = NewSchema(response.Content["application/json"].Schema)
	}
	return Route{
		Method:      caser.String(method),
		Name:        funcName(pathName),
		Path:        ginPathName(pathName),
		Parameters:  parameters,
		Responses:   transformedResponses,
		RequestBody: NewRequestBody(requestBody),
	}
}

func funcName(pathName string) string {
	caser := cases.Title(language.English)

	// Replace - with _ (- is not allowed in go func names)
	pathSegments := strings.Split(pathName, "-")
	nameSegments := make([]string, len(pathSegments))
	for i, segment := range pathSegments {
		nameSegments[i] = caser.String(segment)
	}
	pathName = strings.Join(nameSegments, "_")

	// Convert from '/' delimited path to Camelcase func names
	pathSegments = strings.Split(pathName, "/")
	nameSegments = make([]string, len(pathSegments))
	for i, segment := range pathSegments {
		if len(segment) == 0 {
			continue
		}
		if segment[0] == '{' {
			continue
		}

		nameSegments[i] = caser.String(segment)
	}

	return strings.Join(nameSegments, "")
}

func ginPathName(pathName string) string {
	pathSegments := strings.Split(pathName, "/")
	for i, segment := range pathSegments {
		if len(segment) == 0 {
			continue
		}
		if segment[0] == '{' {
			pathSegments[i] = fmt.Sprintf(":%s", segment[1:len(segment)-1])
		}
	}
	return strings.Join(pathSegments, "/")
}
