package generator

import "github.com/sasswart/gin-in-a-can/openapi"

type ServerInterface struct {
	Routes []Route
	Models []Model
}

func NewServerInterface(tc TemplateConfig, apiSpec openapi.OpenAPI) ServerInterface {
	serverInterface := ServerInterface{
		Routes: newRoutes(tc, apiSpec),
		Models: newModels(tc, apiSpec),
	}

	return serverInterface
}
