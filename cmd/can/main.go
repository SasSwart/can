package main

import (
	"flag"
	"fmt"
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/openapi/operation"
	"github.com/sasswart/gin-in-a-can/openapi/path"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/render"
	"github.com/sasswart/gin-in-a-can/tree"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var Renderer *render.Engine

func main() {
	configData, err := loadConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("loadConfig error: %w", err))
		os.Exit(1)
	}

	fmt.Printf("Reading API specification from \"%s\"\n", absoluteOpenAPIFile(configData))
	apiSpec, err := openapi.LoadAPISpec(
		absoluteOpenAPIFile(configData),
	)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = setupRenderer(configData)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	apiSpec.SetMetadata(map[string]string{
		"package": configData.BasePackageName,
	})

	renderNode := buildRenderNode()
	_, err = tree.Traverse(apiSpec, renderNode)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func buildRenderNode() tree.TraversalFunc {
	return func(key string, parent, node tree.NodeTraverser) (tree.NodeTraverser, error) {
		var templateFile string
		switch node.(type) {
		case *openapi.OpenAPI:
			templateFile = "openapi.tmpl"
		case *path.Item:
			templateFile = "path_item.tmpl"
		case *schema.Schema:
			schemaType := node.(*schema.Schema).Type
			if schemaType != "object" && schemaType != "array" {
				return node, nil
			}
			templateFile = "schema.tmpl"
		case *operation.Operation:
			templateFile = "operation.tmpl"
		}

		if templateFile == "" {
			return node, nil
		}
		_, err := Renderer.Render(node, templateFile)
		if err != nil {
			return node, err
		}

		return node, nil
	}
}

func loadConfig() (openapi.Config, error) {
	wd, err := os.Getwd()
	if err != nil {
		return openapi.Config{}, fmt.Errorf("loadConfig:: could not determine working directory: %w\n", err)
	}

	args := flag.NewFlagSet("can", flag.ExitOnError)

	var configFilePath = args.String("configFile", "", "Specify which config file to use")
	args.BoolVar(&tree.DEBUG, "DEBUG", false, "Enable debug logging")
	_ = args.Parse(os.Args[1:])

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
		return openapi.Config{}, fmt.Errorf("loadConfig:: could not read config file: %w\n", err)
	}

	configData := openapi.Config{
		WorkingDirectory: wd,
		ConfigFilePath:   viper.ConfigFileUsed(),
	}

	err = viper.Unmarshal(&configData)
	if err != nil {
		return openapi.Config{}, fmt.Errorf("loadConfig:: could not parse config file: %w\n", err)
	}

	return configData, nil
}

// absoluteOpenAPIFile uses the current working directory, resolved config file and the openAPI file that was specified
// in the config file to determine the absolute path to and OpenAPI file. It takes into account that any of these,
// except the working directory could be relative.
func absoluteOpenAPIFile(config openapi.Config) string {
	var absoluteOpenAPIFile string
	if filepath.IsAbs(config.OpenAPIFile) {
		absoluteOpenAPIFile = config.OpenAPIFile
	} else {
		if filepath.IsAbs(config.ConfigFilePath) {
			absoluteOpenAPIFile = filepath.Join(
				filepath.Dir(config.ConfigFilePath),
				config.OpenAPIFile,
			)
		} else {
			absoluteOpenAPIFile = filepath.Join(
				config.WorkingDirectory,
				filepath.Dir(config.ConfigFilePath),
				config.OpenAPIFile,
			)
		}
	}

	return absoluteOpenAPIFile
}

func setupRenderer(c openapi.Config) error {
	exe, err := os.Readlink("/proc/self/exe")
	if err != nil {
		return fmt.Errorf("setupRenderer:: could not read /proc/self/exe: %w", err)
	}

	Renderer = render.Engine{}.New(render.GinRenderer{}, render.Config{
		TemplateDirectory: filepath.Join(filepath.Dir(exe), "templates"),

		// this allows us to keep the config.Config type out of the render package
		ModuleName:       c.ModuleName,
		BasePackageName:  c.BasePackageName,
		TemplateName:     c.TemplateName,
		OutputPath:       c.OutputPath,
		WorkingDirectory: c.WorkingDirectory,
		ConfigFilePath:   c.ConfigFilePath,
	})
	return nil
}
