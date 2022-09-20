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
	configWithSpec := config.WithServer(*apiSpec)
	for _, target := range []struct {
		templateDir string
		pkg         string
		file        string
		template    string
	}{
		{"go-gin", "controller", "controller.go", "controller.tmpl"},
		{"go-gin", "controller", "unimplemented.go", "unimplemented.tmpl"},
		{"go-gin", "models", "models.go", "models.tmpl"},
		{"oas-3-0-0", "api", "openapi.yaml", "openapi.tmpl"},
	} {
		file, err := generator.Generate(configWithSpec, target.templateDir, target.template)
		if err != nil {
			fmt.Println(err)
		}

		err = os.WriteFile(path.Join(outputPath, basePackageName, target.pkg, target.file), file, 0777)
		if err != nil {
			fmt.Println(err)
		}
	}
}
