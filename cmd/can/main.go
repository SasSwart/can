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
	generator.Config
}

func main() {
	config := loadConfig()

	apiSpec, err := openapi.LoadOpenAPI(config.OpenAPIFile)
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

func loadConfig() Config {
	var config Config

	configFilePath := flag.String("configFile", "", "Specify which config file to use")
	flag.Parse()
	if *configFilePath == "" {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
	} else {
		viper.SetConfigFile(*configFilePath)
	}

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.Unmarshal(config)

	return config
}
