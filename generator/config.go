package generator

import "github.com/sasswart/gin-in-a-can/openapi"

type Config struct {
	OpenAPIFile          string
	ModuleName           string
	BasePackageName      string
	InvalidRequestStatus string
	ServerInterface      ServerInterface
}

func (tc *Config) WithServer(api openapi.OpenAPI) Config {
	tc.ServerInterface = NewServerInterface(*tc, api)
	return *tc
}
