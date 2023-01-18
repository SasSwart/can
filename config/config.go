package config

import "github.com/sasswart/gin-in-a-can/openapi"

type Config struct {
	Generator        GeneratorConfig
	OpenAPI          openapi.Config
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
