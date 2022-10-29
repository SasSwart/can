package generator

import (
	"github.com/sasswart/gin-in-a-can/openapi"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"path/filepath"
	"strings"
)

type Route struct {
	Name        string
	Path        string
	Method      string
	Parameters  []Parameter
	Responses   map[string]Response
	RequestBody Model
}

func NewRoute(templateConfig Config, pathName, method string, parameters []openapi.Parameter, operation openapi.Operation) Route {
	caser := cases.Title(language.English)
	transformedResponses := make(map[string]Response)
	for r, response := range operation.Responses {

		transformedResponses[r] = Response{
			Name:       funcName(pathName) + caser.String(method) + r + "Response",
			Model:      newModel(templateConfig, response.Content["application/json"].Schema),
			StatusCode: r,
		}
	}

	route := Route{
		Method:    caser.String(method),
		Name:      funcName(pathName),
		Path:      ginPathName(pathName),
		Responses: transformedResponses,
	}

	wrappedParameters := make([]Parameter, len(parameters))
	for p, parameter := range parameters {
		wrappedParameters[p] = newParameterModel(templateConfig, parameter)
	}
	route.Parameters = wrappedParameters

	if operation.RequestBody.Content != nil {
		route.RequestBody = newModel(
			templateConfig,
			operation.RequestBody.Content["application/json"].Schema,
		)

		schema := operation.RequestBody.Content["application/json"].Schema
		name := strings.ReplaceAll(schema.Ref, filepath.Dir(templateConfig.OpenAPIFile), "")
		name = strings.ReplaceAll(name, filepath.Ext(schema.Ref), "")
		route.RequestBody.Name = funcName(name)
	}
	return route
}

func newRoutes(templateConfig Config, apiSpec openapi.OpenAPI) (routes []Route) {
	for pathName, pathItem := range apiSpec.Paths {
		for method, operation := range pathItem.Operations() {
			if operation == nil {
				continue
			}
			routes = append(routes, NewRoute(
				templateConfig,
				pathName,
				method,
				append(pathItem.Parameters, operation.Parameters...),
				*operation,
			))
		}
	}

	return routes
}
