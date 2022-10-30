package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/sasswart/gin-in-a-can/generator"
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/spf13/viper"
)

type Config struct {
	Generator         generator.Config
	OpenAPI           openapi.Config
	OutputPath        string
	TemplateDirectory string
}

func main() {
	config, err := loadConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	apiSpec, err := openapi.LoadOpenAPI(config.OpenAPI.OpenAPIFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	basePackageName := viper.GetString("basePackageName")
	configWithSpec := config.Generator.WithServer(config.OpenAPI.OpenAPIFile, *apiSpec)
	for _, target := range []struct {
		pkg      string
		file     string
		template string
	}{
		{"controller", "controller.go", "controller.tmpl"},
		{"controller", "unimplemented.go", "unimplemented.tmpl"},
		{"models", "models.go", "models.tmpl"},
	} {
		file, err := generator.Generate(configWithSpec, config.TemplateDirectory, target.template)
		if err != nil {
			fmt.Println(err)
		}

		err = os.WriteFile(path.Join(config.OutputPath, basePackageName, target.pkg, target.file), file, 0777)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func loadConfig() (Config, error) {
	var config Config

	args := flag.NewFlagSet("can", flag.ExitOnError)

	configFilePath := args.String("configFile", "", "Specify which config file to use")
	args.Parse(os.Args)

	if *configFilePath == "" {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
	} else {
		viper.SetConfigFile(*configFilePath)
	}

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, fmt.Errorf("could not read config file: %w", err)
	}

	viper.Unmarshal(&config)

	return config, nil
}
