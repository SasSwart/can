package route

import (
	"github.com/sasswart/gin-in-a-can/model"
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/sanitizer"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"path/filepath"
	"strings"
)

type Route struct {
	Name        string `yaml:"name"`
	Path        string `yaml:"path"`
	Method      string `yaml:"method"`
	Parameters  []model.Parameter
	Responses   map[string]model.Response
	RequestBody model.Model
}

func NewRoute(openAPIFile string, pathName, method string, parameters []openapi.Parameter, operation openapi.Operation) Route {
	caser := cases.Title(language.English)
	transformedResponses := make(map[string]model.Response)
	for r, response := range operation.Responses {

		transformedResponses[r] = model.Response{
			Name:       sanitizer.GoFuncName(pathName) + caser.String(method) + r + "Response",
			Model:      model.NewModel(openAPIFile, *response.Content["application/json"].Schema),
			StatusCode: r,
		}
	}

	route := Route{
		Method:    caser.String(method),
		Name:      sanitizer.GoFuncName(pathName),
		Path:      sanitizer.GinPathName(pathName),
		Responses: transformedResponses,
	}

	wrappedParameters := make([]model.Parameter, len(parameters))
	for p, parameter := range parameters {
		wrappedParameters[p] = model.NewParameterModel(openAPIFile, parameter)
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
		route.RequestBody.Name = sanitizer.GoFuncName(name)
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
