package config

import (
	"github.com/sasswart/gin-in-a-can/generator"
	"github.com/sasswart/gin-in-a-can/openapi"
)

type Config struct {
	ModuleName           string
	BasePackageName      string
	InvalidRequestStatus string
	TemplateDirectory    string
	ServerInterface      generator.ServerInterface
}

func (tc *Config) WithServer(openAPIFile string, api openapi.OpenAPI) Config {
	tc.ServerInterface = generator.NewServerInterface(openAPIFile, api)
	return *tc
}
