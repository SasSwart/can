package config

import (
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/server"
)

type Config struct {
	ModuleName           string
	BasePackageName      string
	InvalidRequestStatus string
	TemplateDirectory    string
	ServerInterface      server.Interface
}

func (c *Config) WithServer(openAPIFile string, api openapi.OpenAPI) Config {
	c.ServerInterface = server.NewServerInterface(openAPIFile, api)
	return *c
}
