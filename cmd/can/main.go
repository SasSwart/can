package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/sasswart/gin-in-a-can/config"
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/render"
	golang "github.com/sasswart/gin-in-a-can/render/go"
	"github.com/sasswart/gin-in-a-can/tree"
	"io"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	if !flag.Parsed() { // Sorts out buggy tests
		flag.BoolVar(&config.VersionFlagSet, "version", false, "Print Can version and exit")
		flag.BoolVar(&config.Debug, "debug", false, "Enable debug logging")
		flag.StringVar(&config.ConfigFilePath, "configFile", ".", "Specify which config file to use")
		flag.BoolVar(&config.Dryrun, "dryrun", false,
			"Toggles whether to perform a render without writing to disk."+
				"This works particularly well in combination with -debug")
		flag.StringVar(&config.TemplateNameFlag, "template", "", "Specify which template set to use")
		// TODO should we support stdout as an option here?
		flag.StringVar(&config.OutputPathFlag, "outputPath", "", "Specify where to write output to")
		flag.Parse()
	}

	absCfgPath, err := filepath.Abs(config.ConfigFilePath)
	if err != nil {
		fmt.Printf("could not resolve relative config path: %v\n", err)
		return
	}
	config.ConfigFilePath = absCfgPath

	if config.Debug {
		fmt.Printf("Can: v%s\n", config.SemVer)
	}
	if config.VersionFlagSet {
		os.Exit(0)
	}

	if config.Debug {
		fmt.Printf("[v%s]::Using config file \"%s\".\n", config.SemVer, config.ConfigFilePath)
	}

	configFileReader, err := os.Open(absCfgPath)
	if err != nil {
		fmt.Printf("failed to open config file: %v", err)
		os.Exit(1)
	}

	configs := config.ReadConfigs(configFileReader)

	wg := sync.WaitGroup{}
	for configBytes := range configs {
		go func(configBytes []byte) {
			wg.Add(1)
			defer wg.Done()

			if len(configBytes) == 0 {
				if config.Debug {
					fmt.Println("Skipping empty config")
				}
				return
			}

			configReader := bytes.NewReader(configBytes)
			cfg := mustLoadConfig(configReader)

			err = executeJobForConfig(cfg)
			if err != nil {
				fmt.Println(err)
			}
		}(configBytes)
	}
	wg.Wait()
}

func executeJobForConfig(cfg config.Data) error {
	fmt.Printf("Executing job: %s\n", cfg.Name)

	if config.Debug {
		fmt.Printf("Reading API specification from \"%s\"\n", cfg.GetOpenAPIFilepath())
	}
	apiSpec := mustLoadOpenApiFile(cfg.GetOpenAPIFilepath())

	engine := render.NewEngine(cfg)

	// Setup appropriate renderer via the `strategy` design pattern
	err := setStrategy(&engine, cfg.Template.Strategy)
	if err != nil {
		return err
	}

	apiSpec.SetMetadata(tree.Metadata{
		"package": cfg.Template.BasePackageName,
	})

	if _, err := tree.Traverse(apiSpec, engine.Render); err != nil {
		return err
	}
	return nil
}

func mustLoadOpenApiFile(path string) *openapi.OpenAPI {
	// TODO: What if JSON?
	apiSpec, err := openapi.LoadFromYaml(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return apiSpec
}

// setStrategy creates and applies the renderer for the given strategy. An error is returned if the strategy is invalid
func setStrategy(e *render.Engine, strategy string) error {
	var r render.Renderer
	switch strategy {
	case "go":
		r = &golang.Renderer{}
		r.SetTemplateFuncMap(golang.DefaultFuncMap())
		e.SetRenderer(r)
		return nil
	}
	return fmt.Errorf("%s render strategy not implemented yet", strategy)
}

func mustLoadConfig(reader io.Reader) config.Data {
	cfg := config.Data{}
	err := cfg.Load(reader)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return cfg
}
