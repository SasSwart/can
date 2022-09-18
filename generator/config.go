package generator

import "github.gom/sasswart/gin-in-a-can/openapi"

type TemplateConfig struct {
	ModuleName           string
	BasePackageName      string
	InvalidRequestStatus string
	ServerInterface      ServerInterface
}

func (tc *TemplateConfig) WithServer(api openapi.OpenAPI) TemplateConfig {
	tc.ServerInterface = NewServerInterface(api)
	return *tc
}
