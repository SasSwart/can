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

	outputPath := viper.GetString("outputPath")
	basePackageName := viper.GetString("basePackageName")
	configWithSpec := config.WithSpec(*apiSpec)
	for _, target := range []struct {
		pkg      string
		file     string
		template string
	}{
		{"controller", "controller.go", "controller.tmpl"},
		{"controller", "unimplemented.go", "unimplemented.tmpl"},
		{"models", "models.go", "models.tmpl"},
	} {
		file, err := generator.Generate(configWithSpec, target.template)
		if err != nil {
			fmt.Println(err)
		}

		err = os.WriteFile(path.Join(outputPath, basePackageName, target.pkg, target.file), file, 0777)
		if err != nil {
			fmt.Println(err)
		}
	}
}
