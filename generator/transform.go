package generator

import (
	"fmt"
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

	return Route{
		Method: caser.String(method),
		Name:   funcName(pathName, method),
		Path:   ginPathName(pathName),
	}
}

func funcName(pathName, method string) string {
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
