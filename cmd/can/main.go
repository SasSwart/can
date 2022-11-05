package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/sasswart/gin-in-a-can/generator"
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/spf13/viper"
)

type Config struct {
	Generator         generator.Config
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

	generatorConfig := config.Generator.WithServer(*apiSpec)
	for _, target := range []struct {
		pkg      string
		file     string
		template string
	}{
		{"controller", "controller.go", "controller.tmpl"},
		{"controller", "unimplemented.go", "unimplemented.tmpl"},
		{"models", "models.go", "models.tmpl"},
	} {
		file, err := generator.Generate(generatorConfig, target.template)
		if err != nil {
			fmt.Println(err)
		}

		outputPath := path.Join(
			filepath.Dir(absoluteOpenAPIFile(config)),
			config.OutputPath,
			config.Generator.BasePackageName,
			target.pkg,
		)
		_, err = os.Stat(outputPath)
		switch true {
		case os.IsNotExist(err):
			fmt.Printf("Output path \"%s\" does not exist. Creating it now.\n", outputPath)
			err = os.MkdirAll(outputPath, 0755)
			if err != nil {
				fmt.Println(fmt.Errorf("could not create outputPath: %w\n", err))
			}
		case err != nil:
			fmt.Println(fmt.Errorf("could not determine whether outputPath exists: %w\n", err))
			os.Exit(1)
		}

		fmt.Printf("Rendering %s\n", path.Join(outputPath, target.file))
		err = os.WriteFile(path.Join(outputPath, target.file), file, 0777)
		if err != nil {
			fmt.Println(err)
		}
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
	if filepath.IsAbs(config.Generator.OpenAPIFile) {
		absoluteOpenAPIFile = config.Generator.OpenAPIFile
	} else {
		if filepath.IsAbs(config.ConfigFilePath) {
			absoluteOpenAPIFile = filepath.Join(
				filepath.Dir(config.ConfigFilePath),
				config.Generator.OpenAPIFile,
			)
		} else {
			absoluteOpenAPIFile = filepath.Join(
				config.WorkingDirectory,
				filepath.Dir(config.ConfigFilePath),
				config.Generator.OpenAPIFile,
			)
		}
	}

	return absoluteOpenAPIFile
}
