package route

import (
	"github.com/sasswart/gin-in-a-can/generator"
	"github.com/sasswart/gin-in-a-can/model"
	"github.com/sasswart/gin-in-a-can/openapi"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"path/filepath"
	"strings"
)

type Route struct {
	Name        string `yaml:"name"`
	Path        string `yaml:"path"`
	Method      string `yaml:"method"`
	Parameters  []generator.Parameter
	Responses   map[string]generator.Response
	RequestBody model.Model
}

func NewRoute(openAPIFile string, pathName, method string, parameters []openapi.Parameter, operation openapi.Operation) Route {
	caser := cases.Title(language.English)
	transformedResponses := make(map[string]generator.Response)
	for r, response := range operation.Responses {

		transformedResponses[r] = generator.Response{
			Name:       generator.FuncName(pathName) + caser.String(method) + r + "Response",
			Model:      model.NewModel(openAPIFile, *response.Content["application/json"].Schema),
			StatusCode: r,
		}
	}

	route := Route{
		Method:    caser.String(method),
		Name:      generator.FuncName(pathName),
		Path:      generator.GinPathName(pathName),
		Responses: transformedResponses,
	}

	wrappedParameters := make([]generator.Parameter, len(parameters))
	for p, parameter := range parameters {
		wrappedParameters[p] = generator.NewParameterModel(openAPIFile, parameter)
	}
	route.Parameters = wrappedParameters

	if operation.RequestBody.Content != nil {
		route.RequestBody = model.NewModel(
			openAPIFile,
			*operation.RequestBody.Content["application/json"].Schema,
		)

		schema := operation.RequestBody.Content["application/json"].Schema
		name := strings.ReplaceAll(schema.Ref, filepath.Dir(openAPIFile), "")
		name = strings.ReplaceAll(name, filepath.Ext(schema.Ref), "")
		route.RequestBody.Name = generator.FuncName(name)
	}
	return route
}

func NewRoutes(openAPIFile string, apiSpec openapi.OpenAPI) (routes []Route) {
	//for pathName, pathItem := range apiSpec.Paths {
	//	for method, operation := range pathItem.Operations() {
	//		if operation == nil {
	//			continue
	//		}
	//		//routes = append(routes, NewRoute(
	//		//	openAPIFile,
	//		//	pathName,
	//		//	method,
	//		//	append(pathItem.Parameters, operation.Parameters...),
	//		//	*operation,
	//		//))
	//	}
	//}

	return routes
}
