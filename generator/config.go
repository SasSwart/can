package generator

import "github.gom/sasswart/gin-in-a-can/openapi"

type TemplateConfig struct {
	ModuleName           string
	BasePackageName      string
	InvalidRequestStatus string
	Spec                 ServerInterface
}

func (tc *TemplateConfig) WithSpec(api openapi.OpenAPI) TemplateConfig {
	tc.Spec = NewServerInterface(api)
	return *tc
}
