package main

import (
	"flag"
	"fmt"
	"github.com/sasswart/gin-in-a-can/config"
	"github.com/sasswart/gin-in-a-can/render"
	"os"
	"path/filepath"

	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/spf13/viper"
)

type Config struct {
	Generator         config.Config
	OpenAPI           openapi.Config
	Language          string
	OutputPath        string
	TemplateDirectory string
	WorkingDirectory  string
	ConfigFilePath    string
}

func main() {
	config, err := loadConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("loadConfig error: %w", err))
		os.Exit(1)
	}

	fmt.Printf("Reading API specification from \"%s\"\n", absoluteOpenAPIFile(config))
	apiSpec, err := openapi.LoadOpenAPI(
		absoluteOpenAPIFile(config),
	)
	if err != nil {
		fmt.Println(fmt.Errorf("openapi.LoadOpenAPI error: %w", err))
		os.Exit(1)
	}

	openapi.SetRenderer(apiSpec, openapi.GinRenderer{})

	renderNode := buildRenderNode(config)
	_, err = openapi.Traverse(apiSpec, renderNode)
	if err != nil {
		fmt.Println(fmt.Errorf("openapi.Traverse(apiSpec, renderNode) error: %w", err))
		os.Exit(1)
	}
}

func buildRenderNode(config Config) openapi.TraversalFunc {
	return func(key string, parent, child openapi.Traversable) (openapi.Traversable, error) {
		var templateFile string
		switch child.(type) {
		case *openapi.OpenAPI:
			templateFile = "openapi.tmpl"
		case *openapi.PathItem:
			templateFile = "path_item.tmpl"
		case *openapi.Schema:
			templateFile = "schema.tmpl"
		}
		if templateFile == "" {
			return child, nil
		}
		bytes, err := render.Render(config.Generator, child, templateFile)
		if err != nil {
			return child, err
		}
		fmt.Println(string(bytes))
		fmt.Println(string(bytes))

		return child, nil
	}
}

func loadConfig() (Config, error) {
	wd, err := os.Getwd()
	if err != nil {
		return Config{}, fmt.Errorf("could not determine working directory: %w\n", err)
	}

	args := flag.NewFlagSet("can", flag.ExitOnError)

	var configFilePath = args.String("configFile", "", "Specify which config file to use")
	args.Parse(os.Args[1:])

	if configFilePath == nil {
		fmt.Println("No config file specified.")
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
	} else {
		fmt.Printf("Using config file \"%s\" as specified.\n", *configFilePath)
		viper.SetConfigFile(*configFilePath)
	}

	err = viper.ReadInConfig()
	if err != nil {
		return Config{}, fmt.Errorf("could not read config file: %w\n", err)
	}

	config := Config{
		WorkingDirectory: wd,
		ConfigFilePath:   viper.ConfigFileUsed(),
	}

	viper.Unmarshal(&config)

	return config, nil
}

// absoluteOpenAPIFile uses the current working directory, resolved config file and the openAPI file that was specified
// in the config file to determine the absolute path to and OpenAPI file. It takes into account that any of these,
// except the working directory could be relative.
func absoluteOpenAPIFile(config Config) string {
	var absoluteOpenAPIFile string
	if filepath.IsAbs(config.OpenAPI.OpenAPIFile) {
		absoluteOpenAPIFile = config.OpenAPI.OpenAPIFile
	} else {
		if filepath.IsAbs(config.ConfigFilePath) {
			absoluteOpenAPIFile = filepath.Join(
				filepath.Dir(config.ConfigFilePath),
				config.OpenAPI.OpenAPIFile,
			)
		} else {
			absoluteOpenAPIFile = filepath.Join(
				config.WorkingDirectory,
				filepath.Dir(config.ConfigFilePath),
				config.OpenAPI.OpenAPIFile,
			)
		}
	}

	return absoluteOpenAPIFile
}
