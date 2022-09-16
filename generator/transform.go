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
				pathItem.Parameters,
				operation.Responses,
			))
		}
	}
	return serverInterface
}

type ServerInterface []Route

type Route struct {
	Name       string
	Path       string
	Method     string
	Parameters []openapi.Parameter
	Responses  map[string]openapi.Response
}

func NewRoute(pathName, method string, parameters []openapi.Parameter, responses map[string]openapi.Response) Route {
	caser := cases.Title(language.English)
	return Route{
		Method:     caser.String(method),
		Name:       funcName(pathName),
		Path:       ginPathName(pathName),
		Parameters: parameters,
		Responses:  responses,
	}
}

type Response struct{}

func NewResponseSlice(map[string]openapi.Response) {

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
