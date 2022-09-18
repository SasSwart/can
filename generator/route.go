package generator

import (
	"github.gom/sasswart/gin-in-a-can/openapi"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type ServerInterface []Route

func NewServerInterface(apiSpec openapi.OpenAPI) ServerInterface {
	serverInterface := ServerInterface{}
	for pathName, pathItem := range apiSpec.Paths {
		for method, operation := range pathItem.Operations() {
			serverInterface = append(serverInterface, NewRoute(
				pathName,
				method,
				append(pathItem.Parameters, operation.Parameters...),
				*operation,
			))
		}
	}
	return serverInterface
}

type Route struct {
	Name        string
	Path        string
	Method      string
	Parameters  []openapi.Parameter
	Responses   Responses
	RequestBody RequestBody
}

func NewRoute(pathName, method string, parameters []openapi.Parameter, operation openapi.Operation) Route {
	caser := cases.Title(language.English)
	transformedResponses := Responses{
		Interface: funcName(pathName) + caser.String(method) + "Response",
		Responses: map[string]Response{},
	}
	for r, response := range operation.Responses {
		transformedResponses.Responses[r] = Response{
			Name:   funcName(pathName) + caser.String(method) + r + "Response",
			Schema: response.Content["application/json"].Schema,
		}
	}
	return Route{
		Method:     caser.String(method),
		Name:       funcName(pathName),
		Path:       ginPathName(pathName),
		Parameters: parameters,
		Responses:  transformedResponses,
		RequestBody: RequestBody{
			Name: funcName(pathName) + caser.String(method) + "RequestBody",
			Schema: Schema{
				Schema: operation.RequestBody.Content["application/json"].Schema,
			},
		},
	}
}
