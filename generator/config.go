package generator

import "github.com/sasswart/gin-in-a-can/openapi"

type Config struct {
	ModuleName           string
	BasePackageName      string
	InvalidRequestStatus string
	TemplateDirectory    string
	ServerInterface      ServerInterface
}

func (tc *Config) WithServer(openAPIFile string, api openapi.OpenAPI) Config {
	tc.ServerInterface = NewServerInterface(openAPIFile, api)
	return *tc
}
