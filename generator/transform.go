package generator

import (
	"strings"

	"github.gom/sasswart/gin-in-a-can/openapi"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func NewServerInterface(apiSpec *openapi.OpenAPI) ServerInterface {
	serverInterface := ServerInterface{}
	for pathName, pathItem := range apiSpec.Paths {
		for method := range pathItem.Operations() {
			serverInterface = append(serverInterface, NewRoute(
				pathName,
				method,
			))
		}
	}
	return serverInterface
}

type ServerInterface []Route

type Route struct {
	Name   string
	Path   string
	Method string
}

func NewRoute(pathName, method string) Route {
	caser := cases.Title(language.English)

	pathSegments := strings.Split(pathName, "/")
	nameSegments := make([]string, len(pathSegments))
	for _, segment := range pathSegments {
		if len(segment) == 0 {
			continue
		}
		if segment[0] == '{' {
			continue
		}

		nameSegments = append(nameSegments, caser.String(segment))
	}

	return Route{
		Method: caser.String(method),
		Name:   strings.Join(nameSegments, ""),
		Path:   pathNameToGinPathName(pathName),
	}
}
