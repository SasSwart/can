package generator

import "github.com/sasswart/gin-in-a-can/openapi"

type TemplateConfig struct {
	OpenAPIFile          string
	ModuleName           string
	BasePackageName      string
	InvalidRequestStatus string
	ServerInterface      ServerInterface
}

func (tc *TemplateConfig) WithServer(api openapi.OpenAPI) TemplateConfig {
	tc.ServerInterface = NewServerInterface(*tc, api)
	return *tc
}
