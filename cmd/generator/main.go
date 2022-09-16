package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"

	"github.gom/sasswart/gin-in-a-can/generator"
	"github.gom/sasswart/gin-in-a-can/openapi"
)

func main() {
	// Load config
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var config generator.TemplateConfig

	viper.Unmarshal(&config)

	openAPIEntryPoint := viper.GetString("openAPIFile")

	apiSpec, err := openapi.LoadOpenAPI(openAPIEntryPoint)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = apiSpec.ResolveRefs(path.Dir(openAPIEntryPoint))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	controllerFile := generator.GenerateController(config.WithSpec(*apiSpec))
	err = os.WriteFile("./api/controller/controller.go", controllerFile, 0777)
	if err != nil {
		fmt.Println(err)
	}

	modelsFile := generator.GenerateModels(config.WithSpec(*apiSpec))
	err = os.WriteFile("./api/models/models.go", modelsFile, 0777)
	if err != nil {
		fmt.Println(err)
	}

	unimplementedServerFile := generator.GenerateUnimplementedServer(config.WithSpec(*apiSpec))
	err = os.WriteFile("./api/controller/unimplemented.go", unimplementedServerFile, 0777)
	if err != nil {
		fmt.Println(err)
	}
}
