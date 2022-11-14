package generator

import (
	"github.com/sasswart/gin-in-a-can/model"
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/route"
)

type ServerInterface struct {
	Routes []route.Route
	Models []model.Model
}

func NewServerInterface(openAPIFile string, apiSpec openapi.OpenAPI) ServerInterface {
	serverInterface := ServerInterface{
		Routes: route.NewRoutes(openAPIFile, apiSpec),
		Models: model.NewModels(openAPIFile, apiSpec),
	}

	return serverInterface
}
