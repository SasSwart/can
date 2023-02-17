package config

import (
	"github.com/sasswart/gin-in-a-can/openapi/root"
)

type Config struct {
	Generator        GeneratorConfig
	OpenAPI          root.Config
	OutputPath       string
	WorkingDirectory string
	ConfigFilePath   string
}

type GeneratorConfig struct {
	ModuleName        string
	BasePackageName   string
	TemplateDirectory string
	TemplateName      string
}
