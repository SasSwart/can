package config

import (
	"github.com/sasswart/gin-in-a-can/openapi/root"
)

type Config struct {
	OpenAPI          root.Config
	OutputPath       string
	WorkingDirectory string
	ConfigFilePath   string
}
