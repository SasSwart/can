package generator

import "github.com/sasswart/gin-in-a-can/openapi"

type ServerInterface struct {
	Routes []Route
	Models []Model
}

func NewServerInterface(openAPIFile string, apiSpec openapi.OpenAPI) ServerInterface {
	serverInterface := ServerInterface{
		Routes: newRoutes(openAPIFile, apiSpec),
		Models: newModels(openAPIFile, apiSpec),
	}

	return serverInterface
}
