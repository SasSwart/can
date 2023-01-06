package main

import (
	"fmt"
	"github.com/sasswart/gin-in-a-can/config"
	"github.com/sasswart/gin-in-a-can/render"
	"os"
	"path/filepath"

	"github.com/sasswart/gin-in-a-can/openapi"
)

func main() {
	configData, err := config.LoadConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("loadConfig error: %w", err))
		os.Exit(1)
	}

	fmt.Printf("Reading API specification from \"%s\"\n", absoluteOpenAPIFile(configData))
	apiSpec, err := openapi.LoadOpenAPI(
		absoluteOpenAPIFile(configData),
	)
	if err != nil {
		fmt.Println(fmt.Errorf("openapi.LoadOpenAPI error: %w", err))
		os.Exit(1)
	}

	err = openapi.SetRenderer(apiSpec, render.GinRenderer{})
	if err != nil {
		fmt.Println(fmt.Errorf("openapi.SetRenderer error: %w", err))
		os.Exit(1)
	}

	apiSpec.SetMetadata(map[string]string{
		"package": configData.Generator.BasePackageName,
	})

	renderNode := buildRenderNode(configData)
	_, err = openapi.Traverse(apiSpec, renderNode)
	if err != nil {
		fmt.Println(fmt.Errorf("openapi.Traverse(apiSpec, renderNode) error: %w", err))
		os.Exit(1)
	}
}

func buildRenderNode(config config.Config) openapi.TraversalFunc {
	return func(key string, parent, child openapi.Traversable) (openapi.Traversable, error) {
		var templateFileName string
		switch child.(type) {
		case *openapi.OpenAPI:
			templateFileName = "openapi.tmpl"
		case *openapi.PathItem:
			templateFileName = "path_item.tmpl"
		case *openapi.Schema:
			if child.(*openapi.Schema).Type != "object" {
				return child, nil
			}
			templateFileName = "schema.tmpl"
		case *openapi.Operation:
			templateFileName = "operation.tmpl"
		}

		if templateFileName == "" {
			return child, nil
		}
		_, err := render.Render(config, child, templateFileName)
		if err != nil {
			return child, err
		}

		return child, nil
	}
}

// absoluteOpenAPIFile uses the current working directory, resolved config file and the openAPI file that was specified
// in the config file to determine the absolute path to and OpenAPI file. It takes into account that any of these,
// except the working directory could be relative.
func absoluteOpenAPIFile(config config.Config) string {
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
